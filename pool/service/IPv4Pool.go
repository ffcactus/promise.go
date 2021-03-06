package service

import (
	"promise/base"
	"promise/pool/db"
	"promise/pool/object/dto"
)

var (
	ipv4PoolDB = &db.IPv4Pool{
		DB: base.DB{
			TemplateImpl: new(db.IPv4Pool),
		},
	}
)

// IPv4Pool is the concrete service.
type IPv4Pool struct {
}

// Category returns the category of this service.
func (s *IPv4Pool) Category() string {
	return base.CategoryPoolIPv4
}

// Response creates a new response DTO.
func (s *IPv4Pool) Response() base.GetResponseInterface {
	return new(dto.GetIPv4PoolResponse)
}

// DB returns the DB implementation.
func (s *IPv4Pool) DB() base.DBInterface {
	return ipv4PoolDB
}
