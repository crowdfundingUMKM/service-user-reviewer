package log

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func InitLog() {
	f, err := os.Create("./log/gin.log")
	if err != nil {
		log.Fatal("cannot create open gin.log", err)
	}
	gin.DefaultWriter = io.MultiWriter(f)
}
