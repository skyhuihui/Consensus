package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//声明nodeInfo节点，代表各个小国家
type nodeInfo struct {
	//节点名称
	id string
	//节点路径
	path string
	//http响应
	writer http.ResponseWriter
}

//创建map，存储各个国家的ip地址,也就是url
var nodeTable = make(map[string]string)

//当http服务器，接收到网络请求并且/req 则回调request
func (node *nodeInfo) request(writer http.ResponseWriter, request *http.Request) {
	//该命令允许request请求参数
	request.ParseForm() //解析完毕，打包返回。所以该方法不需要返回值。
	if len(request.Form["warTime"]) > 0 {
		node.writer = writer
		fmt.Println("主节点接收到的参数信息为", request.Form["warTime"][0])
		//fmt.Println(request.Form["warTime"])打印出来是个数组，如果warTime=1111&2222等等，等号后边的是数组，warTime是map的key
		//激活主节点后，向其他的节点发送广播
		node.broadcast(request.Form["warTime"][0], "/prePrepare")
	}
}

//节点发送广播的方法
func (node *nodeInfo) broadcast(msg string, path string) {
	fmt.Println("广播", path)
	//遍历所有的节点
	for nodeId, url := range nodeTable {
		if nodeId == node.id {
			continue
		}
		//使当前节点以外的节点做响应
		http.Get("http://" + url + path + "?warTime=" + msg + "&nodeId=" + node.id)
	}
}

//处理广播后接收到的数据
func (node *nodeInfo) prePrepare(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Println("接收到的广播为", request.Form["warTime"][0])
	if len(request.Form["warTime"]) > 0 {
		node.broadcast(request.Form["warTime"][0], "/prepare")
	}
}

//接收子节点的广播
func (node *nodeInfo) prepare(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	//打印消息
	fmt.Println("接收到的子节点的广播", request.Form["warTime"][0])
	//校验
	if len(request.Form["warTime"]) > 2/3*len(nodeTable) {
		node.authentication(request)
	}
}

var authenticationNodeMap = make(map[string]string)
var authenticationSuceess = false

//校验拜占庭
func (node *nodeInfo) authentication(request *http.Request) {
	if !authenticationSuceess {
		if len(request.Form["nodeId"]) > 0 {
			authenticationNodeMap[request.Form["nodeId"][0]] = "OK"
			//如果有两个国家节点正确的返回了结果
			if len(authenticationNodeMap) > len(nodeTable)/3 {
				authenticationSuceess = true
				node.broadcast(request.Form["warTime"][0], "/commit")
			}
		}
	}
}

//返回成功响应
func (node *nodeInfo) commit(writer http.ResponseWriter, request *http.Request) {
	if writer != nil {
		fmt.Println("拜占庭校验成功")
		//在网页上显示ok
		io.WriteString(node.writer, "ok")
	}
}

// 测试
//   ./main Apple

//   ./main MS

//   ./main Google

//   ./main IBM

//    curl -H "Content-Type: applicaton/json" -X POST -d '{"clientID":"ahnhwi","operation":"GetMyName","timestamp":859381532}' http://localhost:1111/req

func main() {
	//接受终端参数
	userId := os.Args[1]
	fmt.Println(userId)
	//存储4个国家的IP地址
	nodeTable = map[string]string{
		"Apple":  "localhost:1111",
		"MS":     "localhost:1112",
		"Google": "localhost:1113",
		"IBM":    "localhost:1114",
	}

	//创建国家对象
	node := nodeInfo{id: userId, path: nodeTable[userId]}

	//http协议的回调函数
	http.HandleFunc("/req", node.request)
	http.HandleFunc("/prePrepare", node.prePrepare)
	http.HandleFunc("/prepare", node.prepare)
	http.HandleFunc("/commit", node.commit)
	//启动服务器
	if err := http.ListenAndServe(node.path, nil); err != nil {
		fmt.Println(err)
	}
}
