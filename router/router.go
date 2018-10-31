// Copyright 2018 Kirill Lapshin. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package router provides TODO: Docs.
package router

import (
	"github.com/gin-gonic/gin"
)

// Router TODO: Docs.
type Router interface {
	Route(*gin.RouterGroup)
}

// Registered TODO: Docs.
func Registered() []Router {
	return []Router{
		&User{},
	}
}
