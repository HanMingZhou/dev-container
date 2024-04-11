package docker

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	common_models "go-zero-container/common/global/models"
	"go-zero-container/common/utils"
	"go-zero-container/common/utils/aes"
	"go-zero-container/common/utils/container"
	"go-zero-container/common/utils/mapset"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"go-zero-container/app/docker/cmd/api/internal/svc"
	"go-zero-container/app/docker/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateContainerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateContainerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateContainerLogic {
	return &CreateContainerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateContainerLogic) CreateContainer(req *common_models.CreateContainerReq, r *http.Request) (resp *types.CreateContainerResp, err error) {
	// todo: add your logic here and delete this line
	//从 jwt 中解析用户信息
	userName := fmt.Sprintf("%s", l.ctx.Value("Username"))
	userUuid := fmt.Sprintf("%s", l.ctx.Value("UUID"))
	//userName := "jakc"
	//userUuid := "4a04d55b-adef-4a7d-8dd0-b1724799202a"
	//uid := fmt.Sprintf("%s", l.ctx.Value("ID"))
	//conID := r.URL.Query().Get("conId")
	nodeInt, err := strconv.Atoi(r.URL.Query().Get("nodeId"))
	gpuNum, err := strconv.Atoi(r.URL.Query().Get("gpuNum"))
	logx.Error("nodeInt:", nodeInt, " gpuNum:", gpuNum)
	// 判断是否是管理员
	//if userName == "admin" {
	//	return nil, errors.New("管理员不允许创建容器")
	//}
	// 1 *判断容器名称是否重复
	var dataCon common_models.Container
	if !errors.Is(l.svcCtx.DB.Where("containe_name = ? and user_uuid = ?", req.Name, userUuid).First(&dataCon).Error, gorm.ErrRecordNotFound) {
		logx.Error("创建容器-容器名称已创建", zap.Error(err))
		logx.Error("容器名称已创建：", " containe_name:", req.Name, " user_uuid:", userUuid)
		return nil, errors.New("容器名称已创建")
	}
	// 7 *判断ENV——“USER_PASSWD”是否为空，否则 创建8位的随机密码
	env := req.Env
	var passwd string
	var userPasswd = l.svcCtx.Config.DockerAccount.UserPasswd
	statu, word := mapset.InSlice(env, userPasswd)

	if statu {
		//密码校验
		if !utils.RegexpPlas("^(?=.*[a-zA-Z])(?=.*\\d)[a-zA-Z\\d]{6,18}$", word) {
			logx.Error("创建容器-容器密码校验失败")
			return nil, errors.New("密码包含数字和字母，长度6-18位")
		}
		passwd = word

	} else {
		passwd, err = mapset.GenerateRandomString(8)
		if err != nil {
			logx.Error("创建容器-生成随机密码失败", zap.Error(err))
			return nil, errors.New("生成随机密码失败")
		}
		rootPasswd := fmt.Sprintf("%s=%s", userPasswd, passwd)
		req.Env = append(req.Env, rootPasswd)
	}

	args := make(map[string]string)

	// 7.5创建容器-创建nfs卷
	// 通过节点id获取节点ip
	//var nodeInfo common_models.Node
	//if err := l.svcCtx.DB.Where("portainer_id = ?", nodeInt).First(&nodeInfo).Error; err != nil {
	//	logx.Error("创建容器-获取节点信息失败", zap.Error(err))
	//	return nil, err
	//}
	nodeInt = 2
	// 7.7 根据GPU个数绑定GPU
	gpus := make([]string, 0)
	if gpuNum > 0 {
		gpus, err = l.GetAvailableGPU(int32(nodeInt), gpuNum)
		if err != nil {
			logx.Error("创建容器-获取可用GPU失败", zap.Error(err))
			return nil, err
		}
	}

	// 修改portainer容器名
	ContainerName := req.Name
	req.Name = fmt.Sprintf("%s-%s-%s", l.svcCtx.Config.DockerAccount.ConPrefix, userName, ContainerName)
	args["name"] = req.Name
	// 创建容器
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(req)
	if err != nil {
		return nil, err
	}
	// 链接远程服务器
	client, err := container.NewContainer()
	if err != nil {
		logx.Error("Portainer认证失败", zap.Error(err))
		return nil, err
	}
	// 创建
	_, err, rsp := client.CreateContainer(int32(nodeInt), &buf, args)
	if err != nil {
		logx.Error("创建容器失败1", zap.Error(err))
		return nil, err
	}

	gpuStr, _ := json.Marshal(gpus)

	// 容器信息写入数据库
	dataCon = common_models.Container{
		ContainerId:   rsp.ID,
		ContaineName:  ContainerName,
		PortainerName: req.Name,
		Image:         req.Image,
		NodeId:        fmt.Sprint(nodeInt),
		NickName:      userName,
		UserUuid:      userUuid,
		Password:      aes.AesEncrypt(passwd, l.svcCtx.Config.AES.Key),
		Gpus:          string(gpuStr),
		SharedStorage: true,
		OfficialImage: true,
	}
	err = l.svcCtx.DB.Create(&dataCon).Error
	if err != nil {
		logx.Error("创建容器失败2", zap.Error(err))

		return nil, errors.New(err.Error())
	}

	// GPU使用情况更新
	for _, gpu := range gpus {
		// 更新gpu_monitor表
		if err := l.svcCtx.DB.Where("uuid = ?", gpu).Model(&common_models.GpuMonitor{}).Updates(map[string]interface{}{"used": true, "user_id": userUuid, "username": userName}).Error; err != nil {
			logx.Error("创建容器-更新GPU状态失败", zap.Error(err))
			return nil, err
		}
	}
	respId, err := strconv.Atoi(rsp.ID)
	return &types.CreateContainerResp{
		GpuNum: int64(respId),
	}, nil
}

func (l *CreateContainerLogic) GetAvailableGPU(node int32, gpuNum int) (gpus []string, err error) {
	// 工具方法：查找可用GPU 根据节点id、所需数量
	// 1.查找可用GPU used=0 为未使用
	if err = l.svcCtx.DB.Where("node_id = ? and used = ?", node, 0).Model(&common_models.GpuMonitor{}).Pluck("uuid", &gpus).Error; err != nil {
		logx.Error("查找可用GPU失败", zap.Error(err))
		return nil, err
	}
	// 2.判断数量是否足够
	if len(gpus) < gpuNum {
		logx.Error("可用GPU数量不足", zap.Error(err))
		return nil, errors.New("可用GPU数量不足")
	}
	// 3.返回指定数量的GPU的uuid
	return gpus[:gpuNum], nil
}
