/**
  @creator: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @github: https://github.com/OpencvLZG
  @since: 2023/10/16
  @desc: //TODO
**/

package http

import (
	"goDnParse/util/aes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var routerLock sync.RWMutex

// GenerateDnLink 解析下载
func GenerateDnLink(w http.ResponseWriter, r *http.Request) {
	// 获取原始文件URL参数
	url := r.URL.Query().Get("url")
	ext := filepath.Ext(url)
	timer := time.Now()
	currentTime := timer.Format("0601020505.000")
	dnDeal := func(w http.ResponseWriter, r *http.Request) {
		// 请求文件

		resp, err := http.Get(url)

		if err != nil {
			w.WriteHeader(500)
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				return
			}
		}(resp.Body)

		// 设置响应头
		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		w.Header().Set("Content-Disposition", "attachment; filename="+currentTime+ext)

		// 复制body到响应
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			return
		}
	}
	tId, err := aes.GenerateTId(ext)
	if err != nil {
		return
	}
	routerLock.Lock()
	http.HandleFunc("/dn/"+currentTime+tId+ext, dnDeal)
	routerLock.Unlock()
	StatusOk(w, currentTime+tId+ext)
}

// GenerateToken  生成密钥
func GenerateToken(w http.ResponseWriter, r *http.Request) {
	token, err := aes.GenerateToken()
	if err != nil {
		return
	}
	cookie := http.Cookie{
		Name:  "token",
		Value: token,
	}
	http.SetCookie(w, &cookie)
	if err != nil {
		return
	}
	_, err = w.Write([]byte(token))
}

// FileUploadHandle 文件上传
func FileUploadHandle(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token != "www.cilang.buzz" {
		StatusFatal(w, "身份校验失败")
		return
	}
	file, handle, err := r.FormFile("file")
	if err != nil {
		StatusFatal(w, "获取失败")
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	targetFile, err := os.Create("static/file/" + handle.Filename)
	if err != nil {
		StatusFatal(w, "创建失败")
		return
	}
	defer func(targetFile *os.File) {
		err := targetFile.Close()
		if err != nil {
			return
		}
	}(targetFile)
	_, err = io.Copy(targetFile, file)
	if err != nil {
		StatusFatal(w, "写入失败")
		return
	}
	StatusOk(w, "上传成功")
}
