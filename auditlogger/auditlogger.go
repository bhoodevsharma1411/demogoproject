package mycode

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func GetConfig() {
	File, _ := os.OpenFile("./audit.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	new_logger := logger.Config{
		Output: File,
		Format: "${status} - ${time} - ${method} - ${latency} - ${url} \n",
	}
	fmt.Printf("new logger type:- %T", new_logger)
}

func Create_config() logger.Config {
	fmt.Println("Hello audit Logger")
	File, _ := os.OpenFile("./audit.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	audit_logger := logger.Config{Output: File, Format: "${status} - ${time} - ${method} - ${latency} - ${url} \n"}
	return audit_logger
}
