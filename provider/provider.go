// Copyright 2018 Kirill Lapshin. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Package provider provides TODO: Docs.
package provider

import (
	"github.com/goldensplit/gs-server/model"
	"github.com/jinzhu/gorm"
)

// New TODO: Docs.
func New() *Provider {
	// TODO: Refactor hardcode.
	db, err := gorm.Open("postgres", "host=db port=5432 user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		// TODO: Is it good enough to return nil if error?
		return nil
	}
	return &Provider{db}
}

// Provider TODO: Docs.
type Provider struct {
	db *gorm.DB
}

// Create TODO: Docs.
func (p *Provider) Create(m model.Model) error {
	// TODO: Update model.
	p.db.Create(m)
	if err := p.db.Error; err != nil {
		return err
	}
	return nil
}

// Read TODO: Docs.
func (p *Provider) Read(m model.Model) error {
	return nil
}

// Update TODO: Docs.
func (p *Provider) Update(m model.Model) error {
	return nil
}

// Delete TODO: Docs.
func (p *Provider) Delete(m model.Model) error {
	return nil
}
