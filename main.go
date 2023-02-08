package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/db"
	"github.com/shuhaib-kv/Split-Gpay-Golang.git/pkg/routes"
)

var g = gin.Default()

func init() {
	db.ConnectTODatabase()
}

func main() {
	routes.RoutesOfApi(g)
	g.Run(":8082")
}
