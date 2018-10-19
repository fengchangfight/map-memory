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
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
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

	queryArr := params["query"]

	sortByArr := params["sort_by"]

	var query string
	if queryArr != nil && len(queryArr) > 0 {
		query = queryArr[0]
	} else {
		query = ""
	}

	var sort_by string
	if sortByArr != nil && len(sortByArr) > 0 {
		sort_by = sortByArr[0]
	} else {
		sort_by = "created_at"
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

	if len(query) > 0 {
		config.RDB_CONN.Table("mp_memory").Offset((page-1)*config.ITEMS_PER_PAGE).Limit(config.ITEMS_PER_PAGE).Order(sort_by+" desc").Select("id, title, longitude, latitude, icon, created_at, last_update, locked, is_public").Where("user_id = ? and (title like ? or content like ?) ", current_uid, "%"+query+"%", "%"+query+"%").Scan(&result)
		config.RDB_CONN.Table("mp_memory").Where("user_id = ? and (title like ? or content like ?) ", current_uid, "%"+query+"%", "%"+query+"%").Count(&count)
	} else {
		config.RDB_CONN.Table("mp_memory").Offset((page-1)*config.ITEMS_PER_PAGE).Limit(config.ITEMS_PER_PAGE).Order(sort_by+" desc").Select("id, title, longitude, latitude, icon, created_at, last_update, locked, is_public").Where("user_id = ? ", current_uid).Scan(&result)
		config.RDB_CONN.Table("mp_memory").Where("user_id = ?", current_uid).Count(&count)
	}

	// // fmt.Println(result)
	pageObj := model.PaginationVO{
		TotalCount:  count,
		PageSize:    config.ITEMS_PER_PAGE,
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

	current_time := time.Now()
	config.RDB_CONN.Model(&memory).UpdateColumns(entity.Memory{Content: content, LastUpdate: current_time})

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功更新记忆点", "data": memory.ID})
}
func UpdateMemoryLockById(ctx *gin.Context) {
	mid := ctx.Param("id")

	x, _ := ioutil.ReadAll(ctx.Request.Body)
	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(x), &mapResult); err != nil {
		fmt.Println(err)
	}

	locked := mapResult["locked"].(bool)
	read_code := mapResult["read_code"]

	if locked == false {
		// this is to unlock, need to check the read_code
		session := sessions.Default(ctx)
		username := session.Get("user").(string)
		var user entity.User
		config.RDB_CONN.First(&user, "username = ?", username)
		if read_code == nil || user.ReadCode != read_code {
			ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "阅读密码不正确", "data": read_code})
			return
		}
	}
	// update memory id, with title and content

	var memory entity.Memory
	i, err := strconv.ParseInt(mid, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "不合法的ID", "data": err})
	}
	memory.ID = i

	///lockedInt := map[bool]int{true: 1, false: 0}[locked]
	config.RDB_CONN.Model(&memory).UpdateColumns(map[string]interface{}{"locked": locked})

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功更新记忆锁", "data": memory.ID})

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

	//记忆的状态 是公开的还是私密的 0私密1 公开
	is_public := mapResult["is_public"].(bool)

	// create 2 parameters: current timestamp, and login user
	current_time := time.Now()
	session := sessions.Default(ctx)
	current_uid := session.Get("uid").(int64)

	// use the 7 parameter all together to create a new record or memory point
	memory_record := entity.Memory{
		Title:      title,
		Content:    content,
		Longitude:  longitude,
		Latitude:   latitude,
		Icon:       icon,
		UserID:     current_uid,
		CreatedAt:  current_time,
		LastUpdate: current_time,
		Locked:     false,
		IsPublic:   is_public,
	}

	createResult := config.RDB_CONN.Create(&memory_record)

	if createResult.Error != nil {
		// error
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": createResult.Error.(*mysql.MySQLError).Message, "data": memory_record.ID})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功创建记忆点", "data": memory_record.ID})
	}

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
	is_public := mapResult["is_public"].(bool)

	// update memory id, with title and content

	var memory entity.Memory
	i, err := strconv.ParseInt(mid, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "不合法的ID", "data": err})
	}
	memory.ID = i

	current_time := time.Now()
	config.RDB_CONN.Model(&memory).UpdateColumns(map[string]interface{}{"title": title, "content": content, "icon": icon, "is_public": is_public, "last_update": current_time})

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功更新记忆点", "data": memory.ID})

}

func QueryMemoryById(ctx *gin.Context) {
	mid := ctx.Param("id")

	var result model.MemoryDetailVO
	config.RDB_CONN.Table("mp_memory").Select("mp_memory.id, title, content, longitude, latitude, icon, mp_user.username, mp_user.nickname, created_at, last_update, locked, is_public").Joins("inner join mp_user on mp_user.id = mp_memory.user_id").Where("mp_memory.id = ?", mid).Scan(&result)

	session := sessions.Default(ctx)
	username := session.Get("user").(string)
	iamowner := (result.Username == username)
	result.IamOwner = iamowner
	if result.Locked == true {
		// check if read code is provided
		params := ctx.Request.URL.Query()
		read_code_arr := params["read_code"]

		if read_code_arr == nil || len(read_code_arr) < 1 {
			ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "未提供阅读密码", "data": false})
		} else {
			read_code := read_code_arr[0]

			var user entity.User
			config.RDB_CONN.First(&user, "username = ?", username)
			if user.ReadCode == read_code {
				ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功获取记忆详情", "data": result})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "阅读密码不匹配", "data": false})
			}
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功获取记忆详情", "data": result})
	}

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
func UpdateReadCode(ctx *gin.Context) {
	x, _ := ioutil.ReadAll(ctx.Request.Body)
	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(x), &mapResult); err != nil {
		fmt.Println(err)
	}

	oldReadCode := mapResult["old_read_code"].(string)
	newReadCode := mapResult["new_read_code"].(string)

	// check if old password match
	session := sessions.Default(ctx)
	username := session.Get("user").(string)

	var user entity.User
	config.RDB_CONN.First(&user, "username = ?", username)
	if oldReadCode != user.ReadCode {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "原阅读密码不正确", "data": false})
		return
	} else {
		config.RDB_CONN.Model(&user).UpdateColumns(entity.User{ReadCode: newReadCode})
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功更新阅读密码", "data": user.ID})
	}

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

	//0 仅查看自己的 1 查看所有人的
	view_public, err := strconv.Atoi(mapResult["view_public"].(string))
	if err != nil {
		view_public = 0
	}

	// get current uid
	session := sessions.Default(ctx)
	current_uid := session.Get("uid").(int64)
	// select memory by uid and within bound

	var result []model.MemoryVO
	//config.RDB_CONN.Table("mp_memory").Select("id, title, content, longitude, latitude, icon, locked").Where("user_id = ? and longitude > ? and longitude < ? and latitude > ? and latitude < ?", current_uid, south_west_x, north_east_x, south_west_y, north_east_y).Scan(&result)
	DB := config.RDB_CONN
	DB = DB.Table("mp_memory").
		Select("id, title, content, longitude, latitude, icon, locked").
		Where("longitude > ? and longitude < ? and latitude > ? and latitude < ?", south_west_x, north_east_x, south_west_y, north_east_y)
	if view_public == 0 {
		DB = DB.Where("user_id = ?", current_uid)
	} else {
		DB = DB.Where("is_public = ? or user_id = ?", 1, current_uid)
	}
	DB.LogMode(true)
	DB.Scan(&result)

	// fmt.Println(result)

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功获取范围内个人的记忆点", "data": result})
}

func IsFirstDayRegUser(ctx *gin.Context) {
	session := sessions.Default(ctx)
	username := session.Get("user")

	var user entity.User
	config.RDB_CONN.First(&user, "username = ?", username)

	regDate := user.RegDate
	currentDate := time.Now()

	diff := currentDate.Sub(regDate)

	fmt.Println(diff.Hours())
	if diff.Hours() > 3 {
		ctx.String(http.StatusOK, "false")
	} else {
		ctx.String(http.StatusOK, "true")
	}

}

func GetBaseServiceUrl(ctx *gin.Context) {
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

func GetLocationById(ctx *gin.Context) {
	// time.Sleep(2 * time.Second)
	params := ctx.Request.URL.Query()

	memidstr := params["id"]

	var memid string
	if memidstr != nil && len(memidstr) > 0 {
		memid = memidstr[0]
	} else {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "记忆id不正确", "data": ""})
		return
	}

	var memory *entity.Memory
	result := make(map[string]float64)
	DB := config.RDB_CONN
	DB = DB.Table("mp_memory").Where("id = ? ", memid).First(&memory)

	if memory == nil {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "该记忆已失效", "data": ""})
		return
	}
	result["longitude"] = memory.Longitude
	result["latitude"] = memory.Latitude

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功获取记忆位置", "data": result})
}

func GetBaseUrl(ctx *gin.Context) {
	env := getEnv()
	var baseUrl string

	if env == "prod" {
		baseUrl = config.BASE_URL_PROD
	} else if env == "qa" {
		baseUrl = config.BASE_URL_QA
	} else {
		baseUrl = config.BASE_URL_DEV
	}

	ctx.String(http.StatusOK, baseUrl)
}
