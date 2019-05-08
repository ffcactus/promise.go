package redfish

// CollectionElement DTO.
type CollectionElement struct {
	Id string `json:"@odata.id"`
}

// Collection DTO.
type Collection struct {
	Count   int                 `json:"Members@odata.count"`
	Members []CollectionElement `json:"Members"`
}
