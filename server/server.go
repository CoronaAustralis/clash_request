package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func ginServer() {
	r := gin.New()
	r.Use(gin.LoggerWithWriter(log.StandardLogger().Out), gin.RecoveryWithWriter(log.StandardLogger().Out))
	r.GET("/",func(c *gin.Context) {
		for k,v := range SubsrciptionContent.Header {
			// if k == "Content-Disposition"{
			// }
			if k == "Content-Encoding"{
				continue
			}
			c.Header(k,v[0])
		}
		c.Data(200, "text/html; charset=utf-8",SubsrciptionContent.Body)
	})

	r.Run(":"+strconv.Itoa(Config.Port))
}

