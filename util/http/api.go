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
	"net/http"
	"path/filepath"
	"sync"
	"time"
)

var routerLock sync.RWMutex

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
