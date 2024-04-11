package container

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-zero-container/common/global/models"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
)

func NewPortainer(c *Portainer) Portainer {
	return Portainer{
		Config: c.Config,
		Token:  c.Token,
		ApiURL: fmt.Sprintf("%s://%s:%d%s", c.Config.Schema, c.Config.Host, c.Config.Port, c.Config.URL),
		Addr:   fmt.Sprintf("%s:%d", c.Config.Host, c.Config.Port),
	}
}

type temp struct {
	Authentication bool
	Name           string
	URL            string
	Type           int
}

/**
 * @Author Flamingo
 * @Description //创建http请求
 * @Date 2023/1/30 15:51
 **/
func (p *Portainer) makeRequest(t string, url string, body io.Reader, args map[string]string) (*http.Response, error) {
	urlargs := "?"
	for k, v := range args {
		urlargs += fmt.Sprintf("%s=%s&", k, v)
	}
	if urlargs == "?" {
		urlargs = ""
	} else {
		urlargs = urlargs[:len(urlargs)-1]
	}
	req, err := http.NewRequest(t, p.ApiURL+url+urlargs, body)
	logx.Error("requestUrl:", p.ApiURL+url+urlargs)
	if err != nil {
		logx.Error("创建http请求 error", zap.Error(err))
		return nil, err
	}

	// request header 增加 authorization, x-api-key, content-type
	logx.Info("创建http请求时,p.token:", p.Token)
	req.Header.Add("Authorization", "Bearer "+p.Token)
	req.Header.Add("X-API-Key", p.Token)
	logx.Error("x-api-key:", p.Token)
	logx.Error("requestBody:", body)
	req.Header.Add("Content-Type", "application/json")
	c := &http.Client{}

	// Do sends an HTTP request and returns an HTTP response
	return c.Do(req)
}

func (p *Portainer) makeRequestToken(t string, url string, body io.Reader, args map[string]string) (*http.Response, error) {
	urlargs := "?"
	for k, v := range args {
		urlargs += fmt.Sprintf("%s=%s&", k, v)
	}
	if urlargs == "?" {
		urlargs = ""
	} else {
		urlargs = urlargs[:len(urlargs)-1]
	}
	req, err := http.NewRequest(t, p.ApiURL+url+urlargs, body)
	if err != nil {
		logx.Error("创建http请求 error", zap.Error(err))
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+p.AuthToken)
	logx.Error("ExecCreateRequestToken:", p.Token)
	logx.Error("ExecCreateRequestAuthToken:", p.AuthToken)
	//req.Header.Add("X-API-Key", p.Token)
	req.Header.Add("Content-Type", "application/json")
	c := &http.Client{}
	return c.Do(req)
}

/**
 * @Author lhf
 * @Description //portainer 认证
 * @Date 2023/1/30 15:44
 **/
func (p *Portainer) Auth() error {
	authData := make(map[string]string)
	// 获取portainer username password
	authData["Username"] = p.Config.User
	authData["Password"] = p.Config.Password
	// Marshal用于将数据结构转换为 JSON 格式的字节序列
	payload, err := json.Marshal(&authData)
	// post请求：
	res, err := http.Post(p.ApiURL+"/auth", "application/json", bytes.NewReader(payload))
	if err != nil {
		log.Error("访问/auth失败", zap.Error(err))
		return err
	}
	// 判断http请求的结果
	if res.StatusCode != http.StatusOK {
		return errors.New("unauthorized")
	}
	// http请求的body convert to 切片
	jwtString, err := ioutil.ReadAll(res.Body)
	logx.Info("jwt string: ", jwtString)
	_ = res.Body.Close()
	if err != nil {
		log.Error("jwt转换失败", zap.Error(err))
		return err
	}

	//	parse json-encoded data to map
	logx.Info("before parse json-encoded data to map, p.token=", p.Token)
	jwtData := make(map[string]string)
	_ = json.Unmarshal(jwtString, &jwtData)
	p.Token = jwtData["jwt"]
	logx.Info("after parse json-encoded data to p.token", p.Token)

	return err
}

/**
 * @Author Flamingo
 * @Description //Container Exec
 * @Date 2023/2/9 9:53
 **/
func (p *Portainer) Exec(args map[string]string) (*websocket.Conn, error) {
	execurl := fmt.Sprintf("ws://%s/api/websocket/exec", p.Addr)
	urlargs := "?"
	for k, v := range args {
		urlargs += fmt.Sprintf("%s=%s&", k, v)
	}
	if urlargs == "?" {
		urlargs = ""
	}
	urlargs = strings.TrimRight(urlargs, "&")
	url := execurl + urlargs
	log.Println(url)

	backendWS, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		logx.Error("Container Exec", zap.Error(err))
		return nil, err
	}

	return backendWS, nil
}

/**
 * @Author Flamingo
 * @Description //Container ExecCreate
 * @Date 2023/2/7 18:04
 **/
func (p *Portainer) ExecCreate(e int32, id string, CreateReq io.Reader) (*CreateExecRsp, error) {
	url := fmt.Sprintf("/endpoints/%d/docker/containers/%s/exec", e, id)
	res, err := p.makeRequestToken("POST", url, CreateReq, nil)
	logx.Error("CreateContainerByExecUrl:", url)
	logx.Error("CreateContainerByExecCreateReq:", CreateReq)
	if err != nil {
		logx.Error("Container console请求失败", zap.Error(err))
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	_ = res.Body.Close()

	var containerRes ExecConRsp
	err = json.Unmarshal(data, &containerRes)
	if err != nil {
		return nil, err
	}
	rsp := CreateExecRsp{
		Token:  p.Token,
		ExecId: containerRes.ID,
	}

	switch res.StatusCode {
	case http.StatusCreated:
		return &rsp, nil
	case http.StatusInternalServerError:
		return nil, errors.New(fmt.Sprintf("InternalServerError: (%s)", url))
	case http.StatusNotFound:
		return nil, errors.New(fmt.Sprintf("Not found: (%s)", url))
	default:
		return nil, errors.New(fmt.Sprintf("UnhandledError %d: (%s)", res.StatusCode, url))
	}
}

func (p *Portainer) ExecStart(e int32, id string, StartReq io.Reader) (string, error) {
	url := fmt.Sprintf("/endpoints/%d/docker/exec/%s/start", e, id)
	res, err := p.makeRequestToken("POST", url, StartReq, nil)
	if err != nil {
		logx.Error("Container console请求失败", zap.Error(err))
		return "", err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	_ = res.Body.Close()

	return string(data), nil
	//var containerRes ExecConRsp
	//err = json.Unmarshal(data, &containerRes)
	//if err != nil {
	//	return nil, err
	//}
	//rsp := StartExecRsp{
	//	Token:  p.Token,
	//	ExecId: containerRes.ID,
	//}
	//
	//switch res.StatusCode {
	//case http.StatusCreated:
	//	return &rsp, nil
	//case http.StatusInternalServerError:
	//	return nil, errors.New(fmt.Sprintf("InternalServerError: (%s)", url))
	//case http.StatusNotFound:
	//	return nil, errors.New(fmt.Sprintf("Not found: (%s)", url))
	//default:
	//	return nil, errors.New(fmt.Sprintf("UnhandledError %d: (%s)", res.StatusCode, url))
	//}
}

/**
 * @Author lhf
 * @Description //展示所有节点信息
 * @Date 2023/1/30 15:44
 **/
func (p *Portainer) ListEndpoints() ([]Endpoint, error) {
	url := "/endpoints"
	res, err := p.makeRequest("GET", url, nil, nil)
	if err != nil {
		logx.Error("展示所有节点信息失败", zap.Error(err))
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		logx.Error("展示所有节点信息读取body失败", zap.Error(err))
		return nil, err
	}
	var endpoints []Endpoint
	err = json.Unmarshal(data, &endpoints)
	if err != nil {
		logx.Error("Endpoints unmarshaling error", zap.Error(err))
		return nil, err
	}
	return endpoints, err
}

/**
 * @Author Flamingo
 * @Description //展示某个节点所有容器
 * @Date 2023/1/30 15:49
 **/
func (p *Portainer) ListContainers(e int32) ([]Container, error) {
	url := fmt.Sprintf("/endpoints/%d/docker/containers/json", e)
	urlargs := make(map[string]string)
	urlargs["all"] = "1"
	res, err := p.makeRequest("GET", url, nil, urlargs)
	if err != nil {
		logx.Error("展示某个节点所有容器 error", zap.Error(err))
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		logx.Error("展示某个节点所有容器 error", zap.Error(err))
		return nil, err
	}
	var containers []Container
	err = json.Unmarshal(data, &containers)
	return containers, nil
}

/**
 * @Author Flamingo
 * @Description //暂停容器
 * @Date 2023/1/30 15:50
 **/
func (p *Portainer) StopContainer(e int32, id string) error {
	url := fmt.Sprintf("/endpoints/%d/docker/containers/%s/stop", e, id)
	res, err := p.makeRequest("POST", url, nil, nil)
	if err != nil {
		logx.Error("暂停容器", zap.Error(err))
		return err
	}
	_ = res.Body.Close()
	switch res.StatusCode {
	case http.StatusNoContent:
		return nil
	case http.StatusInternalServerError:
		return errors.New(fmt.Sprintf("InternalServerError: (%s)", url))
	case http.StatusNotFound:
		return errors.New(fmt.Sprintf("Not found: (%s)", url))
	case http.StatusNotModified:
		return errors.New(fmt.Sprintf("容器已经关机"))
	default:
		return errors.New(fmt.Sprintf("UnhandledError %d: (%s)", res.StatusCode, url))
	}
}

/**
 * @Author Flamingo
 * @Description //启动容器
 * @Date 2023/1/30 15:50
 **/
func (p *Portainer) StartContainer(e int32, id string) error {
	url := fmt.Sprintf("/endpoints/%d/docker/containers/%s/start", e, id)
	res, err := p.makeRequest("POST", url, nil, nil)
	logx.Error("StartContainerUrl:", url)
	if err != nil {
		logx.Error("启动容器 error", zap.Error(err))
		return err
	}
	_ = res.Body.Close()
	switch res.StatusCode {
	case http.StatusNoContent:
		return nil
	case http.StatusInternalServerError:
		return errors.New(fmt.Sprintf("InternalServerError: (%s)", url))
	case http.StatusNotFound:
		return errors.New(fmt.Sprintf("Not found: (%s)", url))
	default:
		return errors.New(fmt.Sprintf("UnhandledError %d: (%s)", res.StatusCode, url))
	}
}

/**
 * @Author Flamingo
 * @Description //容器重启
 * @Date 2023/2/1 13:39
 **/
func (p *Portainer) RestartContainer(e int32, id string) error {
	url := fmt.Sprintf("/endpoints/%d/docker/containers/%s/restart", e, id)
	res, err := p.makeRequest("POST", url, nil, nil)
	if err != nil {
		logx.Error("容器重启 error", zap.Error(err))
		return err
	}
	_ = res.Body.Close()
	switch res.StatusCode {
	case http.StatusNoContent:
		return nil
	case http.StatusInternalServerError:
		return errors.New(fmt.Sprintf("InternalServerError: (%s)", url))
	case http.StatusNotFound:
		return errors.New(fmt.Sprintf("Not found: (%s)", url))
	default:
		return errors.New(fmt.Sprintf("UnhandledError %d: (%s)", res.StatusCode, url))
	}
}

/**
 * @Author Flamingo
 * @Description //更新容器名称
 * @Date 2023/2/1 13:39
 **/
func (p *Portainer) RenameContainer(e int32, id string, params map[string]string) error {
	url := fmt.Sprintf("/endpoints/%d/docker/containers/%s/rename", e, id)
	res, err := p.makeRequest("POST", url, nil, params)
	if err != nil {
		logx.Error("更新容器名称 error", zap.Error(err))
		return err
	}
	_ = res.Body.Close()
	switch res.StatusCode {
	case http.StatusNoContent:
		return nil
	case http.StatusInternalServerError:
		return errors.New(fmt.Sprintf("InternalServerError: (%s)", url))
	case http.StatusNotFound:
		return errors.New(fmt.Sprintf("Not found: (%s)", url))
	default:
		return errors.New(fmt.Sprintf("UnhandledError %d: (%s)", res.StatusCode, url))
	}
}

/**
 * @Author Flamingo
 * @Description //更新容器
 * @Date 2023/2/1 13:39
 **/
func (p *Portainer) UpdateContainer(e int32, id string, updateBody io.Reader) error {
	url := fmt.Sprintf("/endpoints/%d/docker/containers/%s/update", e, id)
	res, err := p.makeRequest("POST", url, updateBody, nil)
	logx.Error("UpdateContainerUrl:", url)
	logx.Error("UpdateContainerupdateBody:", updateBody)
	if err != nil {
		logx.Error("更新容器 error", zap.Error(err))
		return err
	}
	_ = res.Body.Close()
	switch res.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusInternalServerError:
		return errors.New(fmt.Sprintf("InternalServerError: (%s)", url))
	case http.StatusNotFound:
		return errors.New(fmt.Sprintf("Not found: (%s)", url))
	default:
		return errors.New(fmt.Sprintf("UnhandledError %d: (%s)", res.StatusCode, url))
	}
}

/**
 * @Author Flamingo
 * @Description //容器删除
 * @Date 2023/2/1 13:42
 **/
func (p *Portainer) DeleteContainer(e int32, id string) error {
	url := fmt.Sprintf("/endpoints/%d/docker/containers/%s", e, id)
	args := make(map[string]string)
	args["force"] = "true"
	logx.Error("delete_container_url:", url, " args:", args)
	res, err := p.makeRequest("DELETE", url, nil, args)
	if err != nil {
		logx.Error("容器删除 error", zap.Error(err))
		return err
	}
	switch res.StatusCode {
	case http.StatusNoContent:
		return nil
	case http.StatusInternalServerError:
		return errors.New(fmt.Sprintf("InternalServerError: (%s)", url))
	case http.StatusNotFound:
		return errors.New(fmt.Sprintf("Not found: (%s)", url))
	default:
		return errors.New(fmt.Sprintf("UnhandledError %d: (%s)", res.StatusCode, url))
	}
}

/**
 * @Author Flamingo
 * @Description //容器详情
 * @Date 2023/2/1 14:15
 **/
func (p *Portainer) GetContainers(e int32, id string) (*Container, error) {
	var containers Container
	url := fmt.Sprintf("/endpoints/%d/docker/containers/%s/json", e, id)
	urlargs := make(map[string]string)
	urlargs["all"] = "1"
	res, err := p.makeRequest("GET", url, nil, urlargs)
	if err != nil {
		logx.Error("容器详情 error", zap.Error(err))
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		logx.Error("容器详情 error", zap.Error(err))
		return nil, err
	}
	err = json.Unmarshal(data, &containers)
	return &containers, nil
}

/**
 * @Author Flamingo
 * @Description //创建容器
 * @Date 2023/1/30 15:54
 **/
func (p *Portainer) CreateContainer(e int32, CreateReq io.Reader, createArgs map[string]string) (int, error, *CreateConRsp) {
	url := fmt.Sprintf("/endpoints/%d/docker/containers/create", e)
	logx.Error("url:", url)
	res, err := p.makeRequest("POST", url, CreateReq, createArgs)
	if err != nil {
		logx.Error("创建容器 error", zap.Error(err))
		return 0, err, nil
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, err, nil
	}
	_ = res.Body.Close()

	var containerRes CreateConRsp
	err = json.Unmarshal(data, &containerRes)
	if err != nil {
		return res.StatusCode, err, nil
	}

	switch res.StatusCode {
	case http.StatusOK:
		return res.StatusCode, nil, &containerRes
	case http.StatusNoContent:
		return res.StatusCode, nil, nil
	case http.StatusInternalServerError:
		return res.StatusCode, errors.New(fmt.Sprintf("InternalServerError: (%s)", url)), nil
	case http.StatusNotFound:
		// 失败案例：{"message":"No such image: xxx:latest"}
		var resErr struct {
			Message string `json:"message"`
		}
		logx.Error("errorMessage:", resErr.Message)
		_ = json.Unmarshal(data, &resErr)
		if strings.Contains(resErr.Message, "No such image") {
			return res.StatusCode, errors.New(resErr.Message), nil
		}
		return res.StatusCode, errors.New(resErr.Message), nil
	case http.StatusNotModified:
		return res.StatusCode, errors.New(fmt.Sprintf("Already started: (%s)", url)), nil
	default:
		return res.StatusCode, errors.New(fmt.Sprintf("UnhandledError %d: (%s)", res.StatusCode, url)), nil
	}
}

/**
 * @Author Wuenze
 * @Description
 * @Date 15:34 2023/3/27
 * @Param
 * @return
 **/
func (p *Portainer) InspectContainer(e int32, id string) (*Inspect, error) {
	var inspect Inspect
	url := fmt.Sprintf("/endpoints/%d/docker/containers/%s/json", e, id)
	//urlargs := make(map[string]string)
	res, err := p.makeRequest("GET", url, nil, nil)
	if err != nil {
		logx.Error("容器inspect error", zap.Error(err))
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		logx.Error("容器inspect error", zap.Error(err))
		return nil, err
	}
	// Json-encoded数据解析到inspect结构体中
	err = json.Unmarshal(data, &inspect)
	return &inspect, nil
}

/**
 * @Author Wuenze
 * @Description 获取容器日志
 * @Date 9:23 2023/3/28
 * @Param 节点id 容器id
 * @return []string日志
 **/
func (p *Portainer) GetLogs(e int32, id string, req *models.ContainerLogReq) ([]string, error) {
	url := fmt.Sprintf("/endpoints/%d/docker/containers/%s/logs", e, id)
	urlargs := make(map[string]string)
	if req.Follow {
		urlargs["follow"] = strconv.Itoa(1)
	} else {
		urlargs["follow"] = strconv.Itoa(0)
	}

	urlargs["since"] = strconv.Itoa(req.Since)
	urlargs["utils"] = strconv.Itoa(req.Utils)
	if req.Stderr {
		urlargs["stderr"] = strconv.Itoa(1)
	} else {
		urlargs["stderr"] = strconv.Itoa(0)
	}
	if req.Stdout {
		urlargs["stdout"] = strconv.Itoa(1)
	} else {
		urlargs["stdout"] = strconv.Itoa(0)
	}

	urlargs["tail"] = req.Tail
	if req.Timestamp {
		urlargs["timestamps"] = strconv.Itoa(1)
	} else {
		urlargs["timestamps"] = strconv.Itoa(0)
	}
	res, err := p.makeRequest("GET", url, nil, urlargs)
	if err != nil {
		logx.Error("查看容器日志 error", zap.Error(err))
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		logx.Error("查看容器日志 error", zap.Error(err))
		return nil, err
	}

	// 按照“标题开始SOH”分割
	// 代码中定义了一个匿名函数作为分隔符，该函数接受一个 rune 参数（即 Unicode 字符）
	// 如果该字符是 ASCII 码为 1 的字符（即标题开始符号 SOH），则返回 true，表示应该在此处分割字符串
	// 否则返回 false，表示不分割
	logs := strings.FieldsFunc(string(data), func(c rune) bool {
		if c == 1 {
			return true
		} else {
			return false
		}
	})
	return logs, nil
}

/**
 * @Author Wuenze
 * @Description 列出某个节点的指定容器
 * @Date 10:55 2023/4/6
 * @Param
 * @return
 **/
func (p *Portainer) ListTargetContainers(e string, conIds []string) ([]Container, error) {
	url := fmt.Sprintf("/endpoints/%v/docker/containers/json", e)
	urlargs := make(map[string]string)
	urlargs["all"] = "1"
	conIdsBytes, _ := json.Marshal(conIds)
	urlargs["filters"] = "{\"id\":" + string(conIdsBytes) + "}"
	res, err := p.makeRequest("GET", url, nil, urlargs)
	if err != nil {
		logx.Error("展示某个节点指定容器 error", zap.Error(err))
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		logx.Error("展示某个节点指定容器 error", zap.Error(err))
		return nil, err
	}
	var containers []Container
	err = json.Unmarshal(data, &containers)
	return containers, nil
}

/**
 * @Author Wuenze
 * @Description 保存镜像
 * @Date 9:41 2023/5/5
 * @Param 节点id 容器id 镜像仓库+镜像名+tag
 * @return
 **/
func (p *Portainer) SaveImage(e int32, id, repo string) (string, error) {
	// http://10.101.0.45:9000/api/endpoints/60/docker/commit?container=0a337e2f3bf512194c15f36f08e4c16b32e289caa94e46bda3ff4cec3d13dad3&repo=10.101.0.45%2Fwez%2Fhelloworld%2Fpink-venom:1.0
	url := fmt.Sprintf("/endpoints/%d/docker/commit", e)
	urlargs := make(map[string]string)
	urlargs["container"] = id
	urlargs["repo"] = repo
	res, err := p.makeRequest("POST", url, nil, urlargs)
	if err != nil {
		logx.Error("保存镜像 error", zap.Error(err))
		return "", err
	}
	data, _ := ioutil.ReadAll(res.Body)
	// {"Id":"sha256:2d6be6865fcb3b09201795cde38403109b33feaa9e3f199f7c0d925d8d773049"}
	// data的json转string

	switch res.StatusCode {
	case 201:
		break
	case 500:
		return "", errors.New("保存镜像失败" + string(data))
	default:
		return "", errors.New("保存镜像失败" + string(data))
	}

	var imageSHA256 struct {
		Id string `json:"Id"`
	}
	err = json.Unmarshal(data, &imageSHA256)
	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	_ = res.Body.Close()
	return imageSHA256.Id, nil
}
