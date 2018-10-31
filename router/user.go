// Copyright 2018 Kirill Lapshin. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/goldensplit/gs-server/model"
	"github.com/goldensplit/gs-server/service"
)

// User TODO: Docs.
type User struct {
	// Here will be some context data.
}

// Route TODO: Docs.
func (u User) Route(router *gin.RouterGroup) {
	user := router.Group("/user")

	user.Handle("POST", "/", u.Create)
}

// Create TODO: Docs.
func (u *User) Create(c *gin.Context) {
	// c.JSON(500, "Create User: Not Implemented Yet")
	service.User{}.Create(model.User{Name: "TEST"})
}

// Read TODO: Docs.
func (u *User) Read(c *gin.Context) {
	c.JSON(500, "Read User: Not Implemented Yet")
	// service.User{}.Read(model.User{Name: "TEST"})
}

// Update TODO: Docs.
func (u *User) Update(c *gin.Context) {
	c.JSON(500, "Update User: Not Implemented Yet")
	// service.User{}.Update(model.User{Name: "TEST"})
}

// Delete TODO: Docs.
func (u *User) Delete(c *gin.Context) {
	c.JSON(500, "Delete User: Not Implemented Yet")
	// service.User{}.Delete(model.User{Name: "TEST"})
}
