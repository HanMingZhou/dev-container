package docker

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-container/app/docker/cmd/api/internal/svc"
	"go-zero-container/app/docker/cmd/api/internal/types"
	common_models "go-zero-container/common/global/models"
	"go-zero-container/common/utils"
	"go-zero-container/common/utils/aes"
	"go-zero-container/common/utils/mapset"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

	// 1 从jwt中解析用户信息:  即l.ctx中获取
	userName := fmt.Sprintf("%s", l.ctx.Value("Username"))
	userUuid := fmt.Sprintf("%s", l.ctx.Value("UUID"))
	logx.Error("userNmame", userName)
	logx.Error("userUuid", userUuid)

	// uid, conID暂时不用取值
	//uid := fmt.Sprintf("%s", l.ctx.Value("ID"))
	//conID := r.URL.Query().Get("conId")

	// 节点ID, gpuNum 通过http.request时自行加入参数
	nodeInt, err := strconv.Atoi(r.URL.Query().Get("nodeId"))
	gpuNum, err := strconv.Atoi(r.URL.Query().Get("gpuNum"))
	logx.Error("nodeInt:", nodeInt, " gpuNum:", gpuNum)

	// 判断username是否是管理员
	if userName == "admin" {
		return nil, errors.New("管理员不允许创建容器")
	}

	// 2 查询container name 是否已经创建
	var dataCon common_models.Container
	if !errors.Is(l.svcCtx.DB.Where("containe_name = ? and user_uuid = ?", req.Name, userUuid).First(&dataCon).Error, gorm.ErrRecordNotFound) {
		logx.Error("创建容器-容器名称已创建", zap.Error(err))
		logx.Error("容器名称已创建：", " containe_name:", req.Name, " user_uuid:", userUuid)
		return nil, errors.New("容器名称已创建")
	}

	// 3 判断ENV——“USER_PASSWD”是否为空，否则 创建随机密码
	env := req.Env
	var passwd string
	var userPasswd = l.svcCtx.Config.DockerAccount.UserPasswd
	statu, word := mapset.InSlice(env, userPasswd)
	if statu {
		//密码校验
		if !utils.RegexpPlas("^(?=.*[a-zA-Z])(?=.*\\d)[a-zA-Z\\d]{6,18}$", word) {
			logx.Error("创建容器-容器密码校验失败")
			return nil, errors.New("包含数字和字母6-18位")
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

	// 4 创建容器-创建nfs卷 TODO

	// 5 根据GPU个数绑定GPU, gpu为0,默认http.request请求时,gpuNum为0
	gpus := make([]string, 0)
	if gpuNum > 0 {
		gpus, err = l.GetAvailableGPU(int32(nodeInt), gpuNum)
		if err != nil {
			logx.Error("创建容器-获取可用GPU失败", zap.Error(err))
			return nil, err
		}
	}

	// 6 container_name = username + portainer + 容器名
	ContainerName := req.Name
	req.Name = fmt.Sprintf("%s-%s-%s", l.svcCtx.Config.DockerAccount.ConPrefix, userName, ContainerName)
	args["name"] = req.Name

	// 7 Encode writes the JSON encoding of v to the stream, NewEncoder returns a new encoder that writes to w
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(req)
	if err != nil {
		logx.Error("json解析失败", err)
		return nil, err
	}
	// 8 连接远程服务器
	//client, err := container.NewContainer()
	client := l.svcCtx.Portainer
	/*if err != nil {
		logx.Error("Portainer认证失败", zap.Error(err))
		return nil, err
	}*/
	// 9 开始创建容器
	_, err, rsp := client.CreateContainer(int32(nodeInt), &buf, args)
	if err != nil {
		logx.Error("创建容器失败-portainer", zap.Error(err))
		return nil, err
	}

	// 10 容器信息写入数据库
	gpuStr, _ := json.Marshal(gpus)
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
		logx.Error("创建容器失败-DB", zap.Error(err))
		return nil, errors.New(err.Error())
	}

	// 11 GPU使用情况更新
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
	l.svcCtx.DB.Where("node_id = ? and used = ?", node, 0).Model(&common_models.GpuMonitor{}).Pluck("uuid", &gpus)
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
