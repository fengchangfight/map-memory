package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"map-memory-backend/config"
	"map-memory-backend/entity"
	"map-memory-backend/model"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(ctx *gin.Context) {
	session := sessions.Default(ctx)
	current_uid := session.Get("uid").(int64)
	var result model.UserVO
	config.RDB_CONN.Table("mp_user").Select("id, nickname, email, phone").Where("id = ? ", current_uid).Scan(&result)
	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功获取个人信息", "data": result})
}

func GetPersonMemory(ctx *gin.Context) {

	params := ctx.Request.URL.Query()

	pageStr := params["page"]

	title_queryArr := params["title_query"]

	var title_query string
	if title_queryArr != nil && len(title_queryArr) > 0 {
		title_query = title_queryArr[0]
	} else {
		title_query = ""
	}

	var page int

	var err error

	if pageStr != nil && len(pageStr) > 0 {
		page, err = strconv.Atoi(pageStr[0])
		if err != nil {
			page = 1
		}
	} else {
		page = 1
	}

	// get current uid
	session := sessions.Default(ctx)
	current_uid := session.Get("uid").(int64)

	var result []model.MemoryVO
	var count int

	if len(title_query) > 0 {
		config.RDB_CONN.Table("mp_memory").Offset((page-1)*config.ITEMS_PER_PAGE).Limit(config.ITEMS_PER_PAGE).Order("created_at desc").Select("id, title, content, longitude, latitude, icon, created_at").Where("user_id = ? and title like ? ", current_uid, "%"+title_query+"%").Scan(&result)
		config.RDB_CONN.Model(&entity.Memory{}).Where("user_id = ? and title like ? ", current_uid, "%"+title_query+"%").Count(&count)
	} else {
		config.RDB_CONN.Model(&entity.Memory{}).Where("user_id = ?", current_uid).Count(&count)
		config.RDB_CONN.Table("mp_memory").Offset((page-1)*config.ITEMS_PER_PAGE).Limit(config.ITEMS_PER_PAGE).Order("created_at desc").Select("id, title, content, longitude, latitude, icon, created_at").Where("user_id = ? ", current_uid).Scan(&result)
	}

	// // fmt.Println(result)
	pageObj := model.PaginationVO{
		TotalCount:  count,
		PageSize:    len(result),
		CurrentPage: page,
		Data:        result,
	}

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功获取个人的记忆点列表", "data": pageObj})
}

func GetMyFavLoc(ctx *gin.Context) {
	session := sessions.Default(ctx)
	current_uid := session.Get("uid").(int64)

	var result []model.FavoriteLocVO
	config.RDB_CONN.Table("mp_favorite_location").Order("last_used desc").Select("id, name, longitude, latitude").Where("user_id = ?", current_uid).Scan(&result)
	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功获取常用位置列表", "data": result})

}
func UpdateMemoryContent(ctx *gin.Context) {

	id := ctx.PostForm("id")
	content := ctx.PostForm("content")

	var memory entity.Memory
	idint, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "不合法的ID", "data": err})
	}
	memory.ID = idint
	config.RDB_CONN.Model(&memory).UpdateColumns(entity.Memory{Content: content})

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功更新记忆点", "data": memory.ID})
}

func UpdateMemoryById(ctx *gin.Context) {
	mid := ctx.Param("id")

	x, _ := ioutil.ReadAll(ctx.Request.Body)
	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(x), &mapResult); err != nil {
		fmt.Println(err)
	}

	title := mapResult["title"].(string)
	content := mapResult["content"].(string)
	icon := mapResult["icon"].(string)

	// update memory id, with title and content

	var memory entity.Memory
	i, err := strconv.ParseInt(mid, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "不合法的ID", "data": err})
	}
	memory.ID = i
	config.RDB_CONN.Model(&memory).UpdateColumns(entity.Memory{Title: title, Content: content, Icon: icon})

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功更新记忆点", "data": memory.ID})

}

func QueryMemoryById(ctx *gin.Context) {
	mid := ctx.Param("id")

	//var memory entity.Memory

	//config.RDB_CONN.First(&memory, "id = ?", mid)

	var result model.MemoryDetailVO
	config.RDB_CONN.Table("mp_memory").Select("mp_memory.id, title, content, longitude, latitude, icon, mp_user.username, mp_user.nickname, created_at").Joins("inner join mp_user on mp_user.id = mp_memory.user_id").Where("mp_memory.id = ?", mid).Scan(&result)

	// result := model.MemoryDetailVO{
	// 	ID:        memory.ID,
	// 	Title:     memory.Title,
	// 	Content:   memory.Content,
	// 	Longitude: memory.Longitude,
	// 	Latitude:  memory.Latitude,
	// 	Icon:      memory.Icon,
	// 	Username:  memory.User.Username,
	// 	CreatedAt: memory.CreatedAt.String(),
	// }

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功获取记忆详情", "data": result})

}
func UpdateUserInfo(ctx *gin.Context) {
	x, _ := ioutil.ReadAll(ctx.Request.Body)
	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(x), &mapResult); err != nil {
		fmt.Println(err)
	}

	nickName := mapResult["nick_name"].(string)
	phone := mapResult["phone"].(string)
	email := mapResult["email"].(string)

	session := sessions.Default(ctx)
	username := session.Get("user").(string)

	var user entity.User
	config.RDB_CONN.First(&user, "username = ?", username)

	mp2Update := make(map[string]interface{})

	if len(nickName) > 0 {
		mp2Update["nickname"] = nickName
	}
	if len(phone) > 0 {
		mp2Update["phone"] = phone
	}
	if len(email) > 0 {
		mp2Update["email"] = email
	}

	config.RDB_CONN.Model(&user).Updates(mp2Update)

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功更新用户信息", "data": user.ID})
}

func UpdatePassword(ctx *gin.Context) {
	x, _ := ioutil.ReadAll(ctx.Request.Body)
	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(x), &mapResult); err != nil {
		fmt.Println(err)
	}

	oldPassword := mapResult["old_password"].(string)
	newPassword := mapResult["new_password"].(string)

	// check if old password match
	session := sessions.Default(ctx)
	username := session.Get("user").(string)

	var user entity.User
	config.RDB_CONN.First(&user, "username = ?", username)
	hashedPassword := user.Password
	matched := CheckPasswordHash(oldPassword, hashedPassword)

	// if not match, return with fail message
	if !matched {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "原密码不正确", "data": matched})
	} else {
		encPassword, _ := HashPassword(newPassword)

		config.RDB_CONN.Model(&user).UpdateColumns(entity.User{Password: encPassword})
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功更新密码", "data": user.ID})
	}
}

func GetPersonMemoryInBound(ctx *gin.Context) {
	x, _ := ioutil.ReadAll(ctx.Request.Body)
	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(x), &mapResult); err != nil {
		fmt.Println(err)
	}

	// get bounds by 4 points
	south_west_x := mapResult["south_west_x"].(float64)
	south_west_y := mapResult["south_west_y"].(float64)
	north_east_x := mapResult["north_east_x"].(float64)
	north_east_y := mapResult["north_east_y"].(float64)

	// get current uid
	session := sessions.Default(ctx)
	current_uid := session.Get("uid").(int64)
	// select memory by uid and within bound

	var result []model.MemoryVO
	config.RDB_CONN.Table("mp_memory").Select("id, title, content, longitude, latitude, icon").Where("user_id = ? and longitude > ? and longitude < ? and latitude > ? and latitude < ?", current_uid, south_west_x, north_east_x, south_west_y, north_east_y).Scan(&result)

	// fmt.Println(result)

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功获取范围内个人的记忆点", "data": result})
}

func GetBaseUrl(ctx *gin.Context) {
	env := getEnv()
	var baseServiceUrl string

	if env == "prod" {
		baseServiceUrl = config.BASE_SERVICE_URL_PROD
	} else if env == "qa" {
		baseServiceUrl = config.BASE_SERVICE_URL_QA
	} else {
		baseServiceUrl = config.BASE_SERVICE_URL_DEV
	}
	ctx.String(http.StatusOK, baseServiceUrl)
}
