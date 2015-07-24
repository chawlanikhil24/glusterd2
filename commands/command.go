// The command package defines the command interfaces that need to be implemented by the GlusterD commands
package commands

import (
	"github.com/gorilla/mux"
	"github.com/kshlm/glusterd2/commands/hello"
	"github.com/kshlm/glusterd2/context"
)

type Command interface {
	// SetRoutes should setup the REST API endpoints and handlers for the command
	SetRoutes(r *mux.Router, ctx *context.GDContext) error
	// SetTransactionHandlers will setup the transaction handlers for the command
	//SetTransactionHandlers() error
}

var Commands = []Command{
	&hello.HelloCommand{},
}