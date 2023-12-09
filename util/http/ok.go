/**
  @creator: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @github: https://github.com/OpencvLZG
  @since: 2023/10/15
  @desc: //TODO
**/

package http

import (
	"goDnParse/util/aes"
	"net/http"
	"os"
	"strconv"
	"time"
)

// 重定向
func reDirect(host string) string {
	scriptHeader := "<script>"
	scriptBody := "location.replace(\"" + host + "\");"
	scriptEnd := "</script>"
	scriptData := scriptHeader + scriptBody + scriptEnd
	return scriptData
}

// Auth 认证
func Auth(h http.Handler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// 实现认证逻辑
		// 实现认证逻辑
		token, err := r.Cookie("token")
		if token == nil {
			w.WriteHeader(403)
			_, err = w.Write([]byte(reDirect("/static/err/403.html")))
			if err != nil {
				return
			}
			return
		} else if err != nil {
			w.WriteHeader(403)
			_, err = w.Write([]byte(reDirect("/static/err/403.html")))
			if err != nil {
				return
			}
			return
		}
		deEncryptToken, err := aes.DeEncryptTId(token.Value)
		if err != nil {
			w.WriteHeader(403)
			_, err = w.Write([]byte(reDirect("/static/err/403.html")))
			if err != nil {
				return
			}
			return
		}
		expirationTime, _ := strconv.Atoi(string(deEncryptToken))
		currentTime := time.Now()
		if int64(expirationTime) < currentTime.Unix() {
			w.WriteHeader(403)
			_, err = w.Write([]byte(reDirect("/static/err/403.html")))
			if err != nil {
				return
			}
			return
		}

		// 鉴权通过,调用后续handler
		h.ServeHTTP(w, r)
	}
}

// AuthLoading 认证页面
func AuthLoading(h http.Handler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// 实现认证逻辑
		token, err := r.Cookie("token")
		if token == nil {
			_, err = w.Write([]byte(reDirect("/static/err/loading.html")))
			if err != nil {
				return
			}
			return
		} else if err != nil {
			_, err = w.Write([]byte(reDirect("/static/err/loading.html")))
			if err != nil {
				return
			}
			return
		}
		deEncryptToken, err := aes.DeEncryptTId(token.Value)
		if err != nil {
			_, err = w.Write([]byte(reDirect("/static/err/loading.html")))
			if err != nil {
				return
			}
			return
		}
		expirationTime, _ := strconv.Atoi(string(deEncryptToken))
		currentTime := time.Now()
		if int64(expirationTime) < currentTime.Unix() {
			_, err = w.Write([]byte(reDirect("/static/err/loading.html")))
			if err != nil {
				return
			}
			return
		}

		// 鉴权通过,调用后续handler
		h.ServeHTTP(w, r)
	}
}

// StatusOk Ok
func StatusOk(w http.ResponseWriter, data string) {
	w.WriteHeader(200)
	_, err := w.Write([]byte(data))
	if err != nil {
		return
	}
	return
}

// DefaultHandler 默认处理
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	indexHTML, _ := os.ReadFile("./static/err/404.html")

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "text/html")
	_, err := w.Write(indexHTML)
	if err != nil {
		return
	}
}

// DefaultIconHandler 默认图标
func DefaultIconHandler(w http.ResponseWriter, r *http.Request) {
	indexHTML, _ := os.ReadFile("./static/cat.ico")

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "image/x-icon")
	_, err := w.Write(indexHTML)
	if err != nil {
		return
	}
}

// StatusFatal Fatal
func StatusFatal(w http.ResponseWriter, data string) {
	w.WriteHeader(400)
	_, err := w.Write([]byte(data))
	if err != nil {
		return
	}
	return
}
