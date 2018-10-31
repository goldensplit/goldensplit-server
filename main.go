// Copyright 2018 Kirill Lapshin. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goldensplit/gs-api/router"
)

func main() {
	server := gin.Default()
	api := server.Group("/api")
	for _, route := range router.Registered() {
		route.Route(api)
	}
	server.Run(":80")
}
