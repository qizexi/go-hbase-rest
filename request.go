package rest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//请求对象
type req struct {
	Accept      string      //接受来自服务器的编码
	ContentType string      //发送给服务器的编码
	RespHeader  http.Header //响应的头信息
}

//初始化请求对象
//accept 接受来自服务器的编码
//ctype 发送给服务器的编码
func NewRequest(accept, ctype string) *req {
	rq := new(req)
	if accept == "" {
		accept = "text/xml"
	}
	if ctype == "" {
		ctype = "text/xml"
	}

	rq.Accept = accept
	rq.ContentType = ctype

	return rq
}

//发起Get请求
//url 请求的地址
func (rq *req) Get(url string) (string, error) {
	rurl := url

	reqt, err := http.NewRequest("GET", rurl, nil)
	if err != nil {
		return "", err
	}
	reqt.Header.Add("Accept", rq.Accept)
	reqt.Header.Add("Content-Type", rq.ContentType)
	reqt.Header.Add("Connection", "Close")

	client := &http.Client{}

	resp, err := client.Do(reqt)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

//发起Post请求
//url 请求的地址
//body post请求body
func (rq *req) Post(url string, body string) (string, error) {
	rurl := url

	reqt, err := http.NewRequest("POST", rurl, strings.NewReader(body))
	if err != nil {
		return "", err
	}
	reqt.Header.Add("Accept", rq.Accept)
	reqt.Header.Add("Content-Type", rq.ContentType)
	reqt.Header.Add("Content-Length", fmt.Sprintf("%d", len(body)))
	reqt.Header.Add("Connection", "Close")

	client := &http.Client{}

	resp, err := client.Do(reqt)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

//发起Put请求
//url 请求的地址
//body put请求body
func (rq *req) Put(url string, body string) (string, error) {
	rurl := url

	reqt, err := http.NewRequest("PUT", rurl, strings.NewReader(body))
	if err != nil {
		return "", err
	}
	reqt.Header.Add("Accept", rq.Accept)
	reqt.Header.Add("Content-Type", rq.ContentType)
	reqt.Header.Add("Content-Length", fmt.Sprintf("%d", len(body)))
	reqt.Header.Add("Connection", "Close")

	client := &http.Client{}

	resp, err := client.Do(reqt)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	rq.RespHeader = resp.Header

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

//发起Delete请求
//url 请求的地址
func (rq *req) Delete(url string) (string, error) {
	rurl := url

	reqt, err := http.NewRequest("DELETE", rurl, nil)
	if err != nil {
		return "", err
	}
	reqt.Header.Add("Accept", rq.Accept)
	reqt.Header.Add("Content-Type", rq.ContentType)
	reqt.Header.Add("Connection", "Close")

	client := &http.Client{}

	resp, err := client.Do(reqt)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}
