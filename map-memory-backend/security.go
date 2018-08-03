package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"map-memory-backend/config"
	"map-memory-backend/entity"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func login(c *gin.Context) {

	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Parameters can't be empty"})
		return
	}

	var user entity.User
	config.RDB_CONN.First(&user, "username = ?", username)

	hashedPassword := user.Password

	matched := CheckPasswordHash(password, hashedPassword)

	if matched {
		session.Set("user", username) // In real world usage you'd set this to the users ID
		session.Set("uid", user.ID)   // In real world usage you'd set this to the users ID
		err := session.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate session token"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
	}
}

func RandStringWithCharset(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func ConfirmEmail(ctx *gin.Context) {
	x, _ := ioutil.ReadAll(ctx.Request.Body)
	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(x), &mapResult); err != nil {
		fmt.Println(err)
	}
	confirmCode := mapResult["confirm_code"].(string)

	// if confirmCode empty, return error with message
	if len(strings.TrimSpace(confirmCode)) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "空的验证码", "data": confirmCode})
		return
	}

	// if confirm code not exist in redis, return error with message
	rKey := "mail_confirm_code_" + confirmCode
	rVal := getCacheByKey(config.REDIS_CONN, rKey)
	if rVal == "" {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "无效或过期的验证码", "data": confirmCode})
		return
	}

	// fetch value in redis, get the 2 parts by splitting, email and enc password
	results := strings.SplitN(rVal, ";", 2)
	if len(results) < 2 {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "不正确的验证码", "data": confirmCode})
		return
	}
	email := results[0]
	encPassword := results[1]

	// do the real registration stuff, save user in database
	user := entity.User{
		Username: email,
		Password: encPassword,
		RegDate:  time.Now(),
		Email:    email,
		Nickname: email,
		Phone:    nil,
	}

	rs := config.RDB_CONN.Create(&user)

	if rs.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": rs.Error.(*mysql.MySQLError).Message, "data": user.ID})
	}

	// evict item in redis
	deleteCacheItemByKey(config.REDIS_CONN, rKey)

	// return succeed message
	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "注册成功", "data": email})
}

func UpdateFavilocUsedTime(ctx *gin.Context) {
	fid := ctx.Param("id")

	var favloc entity.FavoriteLocation
	i, err := strconv.ParseInt(fid, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "不合法的ID", "data": err})
	}
	favloc.ID = i
	config.RDB_CONN.Model(&favloc).UpdateColumns(entity.FavoriteLocation{LastUsed: time.Now()})

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功更新常用位置使用时间", "data": favloc.ID})
}

func UpdateMemoryPosById(ctx *gin.Context) {
	mid := ctx.Param("id")

	x, _ := ioutil.ReadAll(ctx.Request.Body)
	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(x), &mapResult); err != nil {
		fmt.Println(err)
	}

	longitude := mapResult["longitude"].(float64)
	latitude := mapResult["latitude"].(float64)

	// update memory id, with title and content

	var memory entity.Memory
	i, err := strconv.ParseInt(mid, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "不合法的ID", "data": err})
	}
	memory.ID = i
	config.RDB_CONN.Model(&memory).UpdateColumns(entity.Memory{Longitude: longitude, Latitude: latitude})

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功更新记忆点位置", "data": memory.ID})

}

func DeleteFavoriteLocationById(ctx *gin.Context) {
	fid := ctx.Param("id")
	i, err := strconv.ParseInt(fid, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "不合法的ID", "data": err})
	}
	config.RDB_CONN.Where("id = ?", i).Delete(entity.FavoriteLocation{})

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功删除常用位置", "data": i})

}

func DeleteMemoryById(ctx *gin.Context) {
	mid := ctx.Param("id")

	i, err := strconv.ParseInt(mid, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "不合法的ID", "data": err})
	}

	config.RDB_CONN.Where("id = ?", i).Delete(entity.Memory{})

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功删除记忆点", "data": i})

}

func Add2FavoriteLocation(ctx *gin.Context) {
	x, _ := ioutil.ReadAll(ctx.Request.Body)
	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(x), &mapResult); err != nil {
		fmt.Println(err)
	}
	// get 5 parameters: longitude, latitude, title, content, icon
	name := mapResult["name"].(string)
	longitude := mapResult["longitude"].(float64)
	latitude := mapResult["latitude"].(float64)
	session := sessions.Default(ctx)
	current_uid := session.Get("uid").(int64)

	favLoc := entity.FavoriteLocation{
		Name:      name,
		Longitude: longitude,
		Latitude:  latitude,
		UserID:    current_uid,
		LastUsed:  time.Now(),
	}

	config.RDB_CONN.Create(&favLoc)

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功添加常用位置", "data": favLoc.ID})

}

func CreateMemoryPoint(ctx *gin.Context) {
	x, _ := ioutil.ReadAll(ctx.Request.Body)
	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(x), &mapResult); err != nil {
		fmt.Println(err)
	}
	// get 5 parameters: longitude, latitude, title, content, icon
	longitude := mapResult["longitude"].(float64)
	latitude := mapResult["latitude"].(float64)
	title := mapResult["title"].(string)
	content := mapResult["content"].(string)
	icon := mapResult["icon"].(string)

	// create 2 parameters: current timestamp, and login user
	current_time := time.Now()
	session := sessions.Default(ctx)
	current_uid := session.Get("uid").(int64)

	// use the 7 parameter all together to create a new record or memory point
	memory_record := entity.Memory{
		Title:     title,
		Content:   content,
		Longitude: longitude,
		Latitude:  latitude,
		Icon:      icon,
		UserID:    current_uid,
		CreatedAt: current_time,
	}

	createResult := config.RDB_CONN.Create(&memory_record)

	if createResult.Error != nil {
		// error
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": createResult.Error.(*mysql.MySQLError).Message, "data": memory_record.ID})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功创建记忆点", "data": memory_record.ID})
	}

}

func RegisterByEmail(ctx *gin.Context) {
	env := getEnv()
	var baseUrl string

	if env == "prod" {
		baseUrl = config.BASE_URL_PROD
	} else if env == "qa" {
		baseUrl = config.BASE_URL_QA
	} else {
		baseUrl = config.BASE_URL_DEV
	}

	x, _ := ioutil.ReadAll(ctx.Request.Body)
	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(x), &mapResult); err != nil {
		fmt.Println(err)
	}
	email := mapResult["email"].(string)
	password := mapResult["password"].(string)

	// check if email is already registered
	var user entity.User
	//config.RDB_CONN.First(&user, "email = ?", email)
	count := 0
	config.RDB_CONN.Model(&user).Where("email = ?", email).Count(&count)
	if count > 0 {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "该邮箱已注册", "data": email})
		return
	}

	encPassword, _ := HashPassword(password)

	// generate random confirm string 16 length
	randConfirmCode := RandStringWithCharset(16)
	// set the confirmstring to redis, wiht key confirm_sessionid and value email;passowrd(encrypted), expire in 1 hour
	rKey := "mail_confirm_code_" + randConfirmCode
	rValue := email + ";" + encPassword
	setCacheItemVal(config.REDIS_CONN, rKey, rValue, 3600)
	// make the confirm url, make the confirm email body and send to the user's mailbox
	confirmUrl := baseUrl + "/mailconfirm/" + randConfirmCode
	confirmMailBody := `<table>
	<tr height="40">
	  <td style="padding-left:25px;padding-right:25px;font-size:18px;font-family:'微软雅黑','黑体',arial;">
		<A href="">尊敬的` + email + `</A>，您好,
	  </td>
	</tr>
	<tr height="15">
	  <td></td>
	</tr>
	<tr height="30">
	  <td style="padding-left:55px;padding-right:55px;font-family:'微软雅黑','黑体',arial;font-size:14px;">
		感谢您使用地图记忆服务。
	  </td>
	</tr>
	<tr height="30">
	  <td style="padding-left:55px;padding-right:55px;font-family:'微软雅黑','黑体',arial;font-size:14px;">
		请点击以下链接进行邮箱验证，链接一小时内有效，以便开始使用您的地图记忆服务账户：
	  </td>
	</tr>
	<tr height="60">
	  <td style="padding-left:55px;padding-right:55px;font-family:'微软雅黑','黑体',arial;font-size:14px;">
		<a href="` + confirmUrl + `" target="_blank" style="display: inline-block;color:#fff;line-height: 40px;background-color: #1989fa;border-radius: 5px;text-align: center;text-decoration: none;font-size: 14px;padding: 1px 30px;">
		  马上验证邮箱
		</a>
	  </td>
	</tr>
	<tr height="10">
	  <td></td>
	</tr>
	<tr height="20">
	  <td style="padding-left:55px;padding-right:55px;font-family:'微软雅黑','黑体',arial;font-size:12px;">
		如果您无法点击以上链接，请复制以下网址到浏览器里直接打开：
	  </td>
	</tr>
	<tr height="30">
	  <td style="padding-left:55px;padding-right:65px;font-family:'微软雅黑','黑体',arial;line-height:18px;">
		<a style="color:#0c94de;font-size:12px;" href="` + confirmUrl + `">
		  ` + confirmUrl + `
		</a>
	  </td>
	</tr>
	<tr height="20">
	  <td style="padding-left:55px;padding-right:55px;font-family:'微软雅黑','黑体',arial;font-size:12px;">
		如果您并未申请地图记忆服务账户，可能是其他用户误输入了您的邮箱地址。请忽略此邮件，或者<a href="mailto:elitewant@163.com">联系我们</a>。
	  </td>
	</tr>
	<tr height="20">
	  <td></td>
	</tr>
  </table>`
	sendMail(email, "地图记忆注册确认服务", confirmMailBody)
	// return succeed message
	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功发送注册确认邮件", "data": email})

}

func logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
	} else {
		log.Println(user)
		session.Delete("user")
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
	}
}

func whoami(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user")

	var user entity.User
	config.RDB_CONN.First(&user, "username = ?", username)

	var displayName string

	if len(user.Nickname) > 0 {
		displayName = user.Nickname
	} else {
		displayName = username.(string)
	}

	if username == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
	} else {
		c.String(http.StatusOK, displayName)
	}
}
