package context

import (
	serverClient "promise/server/client"
	"promise/server/db"
	"promise/server/object/model"
)

// Base is the base Server context.
type Base struct {
	ErrorHandler
	CredentialHandler
	ServerClient serverClient.ServerClientInterface
	DB           db.Server
}

// CreateServerContext Create server context by server.
func CreateServerContext(server *model.Server) *Base {
	var context Base
	context.DB.TemplateImpl = new(db.Server)
	context.ServerClient = serverClient.GetServerClient(server)
	if context.ServerClient == nil {
		return nil
	}
	return &context
}
