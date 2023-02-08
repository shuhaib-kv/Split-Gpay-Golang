package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/routes"
)

var g = gin.Default()

func main() {
	routes.RoutesOfApi(g)
}
