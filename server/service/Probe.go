package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"promise/server/client"
	"promise/server/object/dto"
	"promise/server/object/model"
)

// Probe will try to probe the server.
func Probe(request *dto.DiscoverServerRequest) (*model.ServerBasicInfo, error) {
	c := client.FindBestClient(request.Vender, request.Hostname, request.Username, request.Password)
	if c == nil {
		log.WithFields(log.Fields{
			"hostname": request.Hostname,
		}).Warn("Probe server failed, can not find client.")
		return nil, fmt.Errorf("failed to get server client")
	}

	serverBasicInfo, err := c.GetBasicInfo()
	if err != nil {
		log.WithFields(log.Fields{
			"hostname": request.Hostname,
			"error":    err,
		}).Warn("Probe server failed, can not get basic info.")
		return nil, fmt.Errorf("failed to get server basic info")
	}

	serverBasicInfo.Hostname = request.Hostname
	serverBasicInfo.OriginUsername = &request.Username
	serverBasicInfo.OriginPassword = &request.Password
	if request.Name != nil {
		serverBasicInfo.Name = *request.Name
	} else {
		serverBasicInfo.Name = request.Hostname
	}
	if request.Description != nil {
		serverBasicInfo.Description = *request.Description
	} else {
		serverBasicInfo.Description = ""
	}
	return serverBasicInfo, nil
}
