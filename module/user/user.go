package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
	D "go_api/lib/database"
	G "go_api/lib/global"
	"go_api/lib/response"
	T "go_api/lib/tools"
	V "go_api/lib/valid"
	"time"
)

var valid = V.Validate{}

//当前表操作db
var db = D.DBT("user")

//table 收藏表
const table = "user"

//Get 登录
func Get(c *gin.Context) {
	mobile := c.PostForm("mobile")
	password := c.PostForm("password")
	message, ok := valid.CheckList([]V.Rule{
		V.Rule{Name: "手机号码", Input: mobile, Rule: []string{"require", "numeric", "mobile"}},
		V.Rule{Name: "密码", Input: password, Rule: []string{"require", "minsize:6"}},
	})
	if !ok {
		c.JSON(response.HTTPStatusFaild, G.Json(message, nil))
		return
	}
	user, err := db.Where("mobile", mobile).Where("password", T.MD5(password)).Fields("token,expired_at").First()
	if err != nil {
		c.JSON(response.HTTPStatusFaild, err.Error())
		return
	}
	if len(user) < 1 {
		c.JSON(response.HTTPStatusOK, G.Json("手机号码或密码不正确", nil))
		return
	}
	c.JSON(response.HTTPStatusOK, G.Json("获取数据成功", user))
}

//Post 注册
func Post(c *gin.Context) {
	mobile := c.PostForm("mobile")
	name := c.PostForm("name")
	password := c.PostForm("password")
	price := c.PostForm("price")
	message, ok := valid.CheckList([]V.Rule{
		V.Rule{Name: "手机号码", Input: mobile, Rule: []string{"require", "numeric", "mobile"}},
		V.Rule{Name: "用户名", Input: name, Rule: []string{"require"}},
		V.Rule{Name: "密码", Input: password, Rule: []string{"require", "minsize:6"}},
		V.Rule{Name: "商品价格", Input: price, Rule: []string{"require", "float"}},
	})
	if !ok {
		c.JSON(response.HTTPStatusFaild, G.Json(message, nil))
		return
	}
	//增加
	t, _ := time.ParseDuration("1h")
	expire := time.Now().Add(t).Unix()
	res, err := db.Data(G.MakeData{"mobile": mobile, "nickname": name, "password": T.MD5(password), "token": T.MD5(time.Now()), "expired_at": expire}).Insert()
	if err != nil {
		c.JSON(response.HTTPStatusFaild, G.Json(err.Error(), nil))
	}
	c.JSON(response.HTTPStatusOK, G.Json("数据更新成功", res))

}

//Put 更新接口
func Put(c *gin.Context) {
	name := c.PostForm("name")
	if message, ok := valid.Check("用户名", name, []string{"require"}); !ok {
		c.JSON(response.HTTPStatusFaild, G.Json(message, nil))
		return
	}
	//更新
	sql := fmt.Sprintf("update %s set nickname = '%s' where id = %d", table, name, G.User.UID)
	res, err := D.DB().Query(sql)
	if err != nil {
		c.JSON(response.HTTPStatusFaild, G.Json(err.Error(), nil))
		return
	}
	c.JSON(response.HTTPStatusFaild, G.Json("数据更新成功", res))
	return
}

//Delete 删除接口
func Delete(c *gin.Context) {

}
