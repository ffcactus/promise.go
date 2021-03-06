package entity

import (
	"promise/server/object/model"
)

// Board is board hardware
type Board struct {
	ServerRef string
	EmbeddedResource
	ProductInfo
	CardNo          *int
	DeviceLocator   *string
	DeviceType      *string
	Location        *string
	CPLDVersion     *string
	PCBVersion      *string
	BoardName       *string
	BoardID         *string
	ManufactureDate *string
}

// ToModel will create a new model from entity.
func (e *Board) ToModel() *model.Board {
	m := new(model.Board)
	createResourceModel(&e.EmbeddedResource, &m.Resource)
	createProductInfoModel(&e.ProductInfo, &m.ProductInfo)
	m.CardNo = e.CardNo
	m.DeviceLocator = e.DeviceLocator
	m.DeviceType = e.DeviceType
	m.Location = e.Location
	m.CPLDVersion = e.CPLDVersion
	m.PCBVersion = e.PCBVersion
	m.BoardName = e.BoardName
	m.BoardID = e.BoardID
	m.ManufactureDate = e.ManufactureDate
	return m
}

// Load will load data from model.
func (e *Board) Load(m *model.Board) {
	updateResourceEntity(&e.EmbeddedResource, &m.Resource)
	updateProductInfoEntity(&e.ProductInfo, &m.ProductInfo)
	e.CardNo = m.CardNo
	e.DeviceLocator = m.DeviceLocator
	e.DeviceType = m.DeviceType
	e.Location = m.Location
	e.CPLDVersion = m.CPLDVersion
	e.PCBVersion = m.PCBVersion
	e.BoardName = m.BoardName
	e.BoardID = m.BoardID
	e.ManufactureDate = m.ManufactureDate
}
