package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/database"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/routes"
)

var g = gin.Default()

func init() {
	database.ConnectTODatabase()
}

func main() {
	routes.RoutesOfApi(g)
	g.Run(":8089")
}
