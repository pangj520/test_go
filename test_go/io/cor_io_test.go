/**
 *@Author: pangj
 *@Description:
 *@File: cor_io
 *@Version:
 *@Date: 2022/04/15/16:23
 */

package io

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestHttpGet(t *testing.T) {
	rsp, err := http.Get("")
	if err != nil {
		t.Errorf("get err :%v\n", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			t.Errorf("close file err :%v\n", err)
		}
	}(rsp.Body)

	body, err := ioutil.ReadAll(rsp.Body)
	t.Logf("body len :%v, read err :%v\n", len(body), err)
}

func TestReadAll(t *testing.T) {
	file, err := os.Open("D:\\极客时间\\毛大\\4.降级&重试.mp4")
	if err != nil {
		t.Errorf("open err :%v\n", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	_, err = ioutil.ReadAll(file)
	if err != nil {
		t.Errorf("read all failed :%v\n", err.Error())
		return
	}
}
