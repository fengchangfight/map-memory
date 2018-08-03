package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"map-memory-backend/config"
	"net"
	"net/mail"
	"net/smtp"
	"time"

	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func newRedisClient() *redis.Client {
	env := getEnv()
	var redisServer string
	var redisAuth string
	if env == "dev" {
		redisServer = config.REDIS_SERVER_DEV
		redisAuth = config.REDIS_AUTH_DEV
	} else if env == "qa" {
		redisServer = config.REDIS_SERVER_QA
		redisAuth = config.REDIS_AUTH_QA
	} else if env == "prod" {
		redisServer = config.REDIS_SERVER_PROD
		redisAuth = config.REDIS_AUTH_PROD
	} else {
		redisServer = config.REDIS_SERVER_DEV
		redisAuth = config.REDIS_AUTH_DEV
	}
	client := redis.NewClient(&redis.Options{
		Addr:     redisServer + ":" + config.REDIS_PORT,
		Password: redisAuth,
		DB:       0, // use default DB
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic("Error connecting redis")
	}

	return client
}

func setCacheItemVal(client *redis.Client, key string, val string, timeinsecs int) bool {
	err := client.Set(key, val, time.Duration(timeinsecs)*time.Second).Err()
	if err != nil {
		return false
	}
	return true
}

func deleteCacheItemByKey(client *redis.Client, key string) bool {
	client.Del(key)
	return true
}

func getCacheByKey(client *redis.Client, key string) string {
	val, err := client.Get(key).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		panic(err)
	} else {
		return val
	}
}

func sendMail(toAddr string, subject string, body string) {
	diyiMail := config.DIYI_MAIL
	from := mail.Address{config.MAILNAME, diyiMail}
	to := mail.Address{"", toAddr}

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	message += "Content-Type: text/html; charset=UTF-8\r\n"
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := config.SMTP_SERVER + ":" + config.SMTP_PORT
	host, _, _ := net.SplitHostPort(servername)
	auth := smtp.PlainAuth("", diyiMail, config.MAIL_AUTH_CODE, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()
}
