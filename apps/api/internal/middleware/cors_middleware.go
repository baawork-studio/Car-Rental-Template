package middleware

import (
	"strings"
	"github.com/gin-contrib/cors"
)

func Cors(origins string) cors.Config { return cors.Config{AllowOrigins: strings.Split(origins, ","), AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}, AllowHeaders: []string{"Origin", "Content-Type", "Authorization"}} }
