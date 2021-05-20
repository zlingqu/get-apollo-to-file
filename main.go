package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	configServerUrl = flag.String("configServerUrl", "", "Apollo的http地址")
	appId           = flag.String("appId", "", "Apollo应用的appId")
	clusterName     = flag.String("clusterName", "", "Apollo的集群名")
	namespaceName   = flag.String("namespaceName", "", "Apollo的namespaceName")
	destFilePath    = flag.String("destFilePath", "", "目的文件路径")
)

func httpGet(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil
	}

	fmt.Println(string(body))

	return body
}

func main() {
	flag.Parse()
	var apolloResponse map[string]string

	requestApolloUrl := *configServerUrl + "/configfiles/json/" + *appId + "/" + *clusterName + "/" + *namespaceName

	resp := httpGet(requestApolloUrl)
	if resp != nil {
		e := json.Unmarshal(resp, &apolloResponse)
		if e != nil {
			log.Fatalln("对apollo的请求结构进行反序列化失败！")
		}
	}

	// log.Println("apollo返回的结果数据为：", apolloResponse)

	f, err := os.OpenFile(*destFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开目标文件错误")
	}
	defer f.Close()

	var (
		Len     = len(apolloResponse)
		i       = 0
		lineStr string
	)
	if Len < 1 {
		log.Println("配置中心没有没有数据")
		return
	}

	w := bufio.NewWriter(f)
	fmt.Fprintln(w, "\nENV \\")
	for k, v := range apolloResponse {
		if i < Len-1 {
			lineStr = fmt.Sprintf("%s=%s \\", k, v)
		} else {
			lineStr = fmt.Sprintf("%s=%s", k, v)
		}
		i++
		fmt.Fprintln(w, lineStr)

	}
	w.Flush()

}