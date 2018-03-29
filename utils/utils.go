package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"crypto/rand"
	"time"

	"github.com/gorilla/mux"
	"strings"
	"sort"
	"encoding/hex"
	"crypto/md5"
	"regexp"
)


// 获取当前时间戳
func NowTimestamp() int64 {
	t := time.Now()
	timestamp := t.UTC().UnixNano() / 1000000000
	return timestamp
}

// 必须是时间戳
func StrtoTimestamp(data string) int64 {
	 if data == ""{
	 	return 0
	 }
	 dlist := strings.Split(data,"")
	 if len(dlist) != 10{
	 	return 0
	 }
	 ret,err:= strconv.ParseInt(data,10,64)
	 if err != nil{
	 	return 0
	 }
	 return ret
}

// genToken 生成随机字符串
func GenToken(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func GetRequestValue(r *http.Request, key string) int64 {
	val := r.FormValue(key)
	if val == "" {
		var has bool
		val, has = mux.Vars(r)[key]
		if !has || val == "" {
			val = "0"
		}
	}
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		fmt.Println("[GetRequestValue]获取", key, "错误", err, "i:", i)
	}
	return i
}

func GetRequestValueStr(r *http.Request, key string) string {
	val := r.FormValue(key)
	if val == "" {
		var has bool
		val, has = mux.Vars(r)[key]
		if !has || val == "" {
			val = ""
		}
	}
	return val
}

func StrToint(str string) int64 {
	if str != ""{
		new := strings.Replace(str," ","",-1)
		i,err := strconv.ParseInt(new,10,64)
		if err != nil{
			fmt.Println("[StrToint] 字符串转换数字失败,",err)
		}
		return i
	}
	return 0
}

func StrMd5(str string)  string {
	h := md5.New()
	base_str := str+"unicloud.17710633221"
	tmp := strings.Split(base_str,"")
	sort.Strings(tmp)
	str = strings.Join(tmp,"")
	h.Write([]byte(str)) //
	cipherStr := h.Sum(nil)
	md5_key:= fmt.Sprintf("%s", hex.EncodeToString(cipherStr)) // 输出加密结果
	return md5_key
}
//必须输数字
func IsNumber(data string) bool {
	b,_ := regexp.MatchString("^[0-9]+$", data)
	return b
}
// 必须输字母
func IsLetter(data string) bool {
	b,_ := regexp.MatchString("^[a-zA-Z]+$", data)
	return b
}

// 以字母开头数字字母下划线结尾
func IsLetterNumber(data string) bool {
	b,_ := regexp.MatchString("^[a-zA-Z]+[_0-9a-zA-Z]+$", data)
	return b
}