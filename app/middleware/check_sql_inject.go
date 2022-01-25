package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lkeme/QSearch/app/response"
	"net/http"
	"net/http/httputil"
	"reflect"
	"regexp"
)

var skipParamsForSQLInjectMap = map[string]int{
	"file":           1,
	"encrypt_data":   1,
	"openid":         1,
	"iv":             1,
	"name":           1,
	"link":           1,
	"head_pic_link":  1,
	"pic_link":       1,
	"id_front_pic":   1,
	"id_reverse_pic": 1,
	"education_pic":  1,
}

//FilteredSQLInject sql注入判断
func FilteredSQLInject(oldStr string) bool {
	ok := true
	str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	re, err := regexp.Compile(str)
	if err != nil {
		//panic(err.Error())
		return false
	}
	ok = re.MatchString(oldStr)
	if ok {
		return false
	} else {
		return true
	}
}

// CheckSQLInject 检查Sql注入
func CheckSQLInject() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := httputil.DumpRequest(c.Request, true)
		if err != nil {
			err = fmt.Errorf("parse request body failed, err: %s", err)
			response.InternalServerError(c, err.Error())
			return
		}
		//fmt.Printf("%s\n", string(body))
		//fmt.Printf("%#v\n", c.Request)
		if bytes.Index(body, []byte("application/json")) != -1 {
			//检查json参数
			//过滤http头部,获取body
			body = body[bytes.Index(body, []byte("User-Agent")):]
			index := bytes.Index(body, []byte("{"))
			if index != -1 {
				//判读是否找到body
				body = body[index:]
			} else {

				body = nil
			}
			if len(body) != 0 {
				//检查request json参数
				fmt.Println(string(body))
				m := make(map[string]interface {
				})
				err = json.Unmarshal(body, &m)
				if err != nil {
					err = fmt.Errorf("parse request body failed, err: %s", err.Error())
					response.InternalServerError(c, err.Error())
					return
				}
				for k, v := range m {
					if skipParamsForSQLInjectMap[k] == 1 {
						continue
					}
					if reflect.TypeOf(v).String() == "string" {
						if FilteredSQLInject(v.(string)) {
							err := fmt.Errorf("sql注入攻击 %s", v)
							response.InternalServerError(c, err.Error())
							return
						}
					}
				}
			}
		}
		if c.Request.Form == nil {
			//检查get和post中的参数
			err := c.Request.ParseMultipartForm(32 << 20)
			if err != nil {
				response.InternalServerError(c, err.Error())
				return
			}
		}
		for k, arr := range c.Request.Form {
			if c.Request.Method != http.MethodGet {
				fmt.Printf("%s=%v&", k, arr)
			}
			if skipParamsForSQLInjectMap[k] == 1 {
				continue
			}
			for _, v := range arr {
				if FilteredSQLInject(v) {
					err := errors.New("sql注入攻击")
					response.InternalServerError(c, err.Error())
					return
				}
			}
		}
		c.Next()
	}
}
