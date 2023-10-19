/**
  @creator: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @github: https://github.com/OpencvLZG
  @since: 2023/10/10
  @desc: //TODO
**/

package main

import (
	"goDnParse/util/aes"
	"strconv"
	"time"
)

func main() {
	//url := "https://www.bing.com/images/search?view=detailV2&ccid=kaFOfG4Q&id=453890841DEF3E9462E4D11F97798ED948B8E8A4&thid=OIP.kaFOfG4QtRH3DMGj0K_pYQHaFP&mediaurl=https%3a%2f%2fwework.qpic.cn%2fwwpic%2f39117_rUd3ALkfSBWKy4o_1646964249%2f0&exph=794&expw=1123&q=https%3a%2f%2fwework.qpic.cn%2fwwpic%2f450425_duhVzgLoTzuFx5P_1650445417%2f0.lll&simid=607988948188074634&FORM=IRPRST&ck=AA2B30E4D8A054B070EDB63AE94AB055&selectedIndex=2"
	//data := path.Ext(url)
	//println(data)
	//data2 := filepath.Ext(url)
	//println(data2)
	//tid, err := GenerateTId("")
	//if err != nil {
	//	return
	//}
	//println(tid)
	//
	//// 解密
	//
	tid := "D/33BeSalrrGvkSInhIWpg=="
	decrypted, _ := aes.DeEncryptTId(tid)
	println(string(decrypted))
	expirationTime, _ := strconv.Atoi(string(decrypted))
	currentTime := time.Now()
	println(currentTime.Unix())
	println(int64(expirationTime) < currentTime.Unix())
	if int64(expirationTime) < currentTime.Unix() {
		println(1)
	}
	//
	//// 打印解密结果
	//println(string(decrypted))

}
