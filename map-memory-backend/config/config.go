/* Copyright Fengchang Xie & Huifeng Zhang, 2018 All Rights Reserved */
package config

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

/**
 * Created by Fengchang on 2018/5/30 12:53.
 */

const (
	// dev
	DB_HOST              = "127.0.0.1"
	RDB_USER             = "root"
	RDB_PASSWORD         = "Passw0rd"
	REDIS_SERVER_DEV     = "127.0.0.1"
	REDIS_AUTH_DEV       = "???"
	BASE_URL_DEV         = "http://localhost:8043"
	RDB_NAME_DEV         = "map_memory"
	BASE_SERVICE_URL_DEV = "http://localhost:8042"

	// qa
	REDIS_SERVER_QA     = "127.0.0.1"
	REDIS_AUTH_QA       = "???"
	BASE_URL_QA         = "http://www.ditujiyi.com"
	RDB_USER_QA         = "??"
	RDB_PASSWORD_QA     = "??"
	RDB_NAME_QA         = "map_memory"
	BASE_SERVICE_URL_QA = "http://www.ditujiyi.com/service"

	//prod
	RDB_USER_PROD         = "???"
	RDB_PASSWORD_PROD     = "???"
	REDIS_SERVER_PROD     = "127.0.0.1"
	REDIS_AUTH_PROD       = "???"
	BASE_URL_PROD         = "http://www.ditujiyi.com"
	RDB_NAME_PROD         = "map_memory"
	BASE_SERVICE_URL_PROD = "http://www.ditujiyi.com/service"

	// common
	DB_PORT        = "9080"
	GIN_PORT       = ":8042"
	REDIS_PORT     = "6379"
	DIYI_MAIL      = "ditujiyi@163.com"
	SMTP_SERVER    = "smtp.163.com"
	SMTP_PORT      = "465"
	MAIL_AUTH_CODE = "x5bdfhd29hu"

	ITEMS_PER_PAGE = 100

	MAILNAME = "地图记忆"
)

var REDIS_CONN *redis.Client
var RDB_CONN *gorm.DB
