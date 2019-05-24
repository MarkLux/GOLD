package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const GOLD_DIR = "/Users/hanxiao/WorkPlace/GOLD/"
const GOLD_FILE_DIR = GOLD_DIR + "serving/wrapper/gold/gold_biz.go"
const SH_SCRIPT_DIR = GOLD_DIR + "serving/preview/update.sh"
const WRAPPER_DIR  = GOLD_DIR + "serving/wrapper/"
func main() {
	r := gin.Default()
	r.Use(Cors())

	r.POST("/preview", preview)

	r.Run(":8082")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func preview(ctx *gin.Context) {
	code := ctx.Request.FormValue("code")
	funcServicePidFile, err := os.Open("/Users/hanxiao/funcServicePid")
	if err != nil {
		log.Fatal("fail to open funcServicePidFile")
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1002,
			"err_msg": err.Error(),
		})
	}
	defer funcServicePidFile.Close()
	oldServicePid, err := ioutil.ReadAll(funcServicePidFile)

	err = exec.Command("sh", SH_SCRIPT_DIR, code, GOLD_FILE_DIR, string(oldServicePid), WRAPPER_DIR).Run()
	if err != nil {
		//log.Fatal("fail to exec sh script",err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":     0,
		"res":      err,
		"code_str": code,
	})
}
