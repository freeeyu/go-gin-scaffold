package tools

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

//MD5 md5加密,支持字符串,整数
func MD5(obj interface{}) string {
	var str string = ""
	//change to string
	switch obj.(type) {
	case string:
		str = obj.(string)
	case int:
		str = strconv.Itoa(obj.(int))
		break
	case int64:
		str = strconv.FormatInt(obj.(int64), 10)
		break
	}
	m := md5.Sum(bytes.NewBufferString(str).Bytes())
	return hex.EncodeToString(m[:])
}

//HTTPGet http get 请求
func HTTPGet(url string) []byte {
	// uuu := proxyAddr.TYPE + "://" + proxyAddr.HOST + ":" + strconv.Itoa(proxyAddr.PORT)
	// proxy, _ := url.Parse(uuu)
	client := &http.Client{
		// Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil) //建立一个请求
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}
	response, err := client.Do(req)
	if err != nil {
		// fmt.Println(err.Error())
		// r := rand.New(rand.NewSource(time.Now().Unix()))
		// index := r.Intn(len(proxyList))
		// return getAddr(addr, proxyList[index])
		fmt.Println("Fatal error ", err.Error())
		os.Exit(0)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return body
}
