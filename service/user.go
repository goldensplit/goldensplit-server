// Copyright 2018 Kirill Lapshin. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package service provides TODO: Docs.
package service

import (
	"github.com/goldensplit/gs-api/model"
	"github.com/goldensplit/gs-api/provider"
)

// NewUser TODO: Docs.
func NewUser() *User {
	return &User{provider.New()}
}

// User TODO: Docs.
type User struct {
	p *provider.Provider
	// Here will be some context data.
}

// Create TODO: Docs.
func (u User) Create(m model.Model) error {
	return u.p.Create(m)
}

// Read TODO: Docs.
func (u User) Read(m model.Model) error {
	return u.p.Read(m)
}

// Update TODO: Docs.
func (u User) Update(m model.Model) error {
	return u.p.Update(m)
}

// Delete TODO: Docs.
func (u User) Delete(m model.Model) error {
	return u.p.Delete(m)
}
