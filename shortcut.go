package serco

import (
	"github.com/aledenza/serco/client"
	"github.com/aledenza/serco/database"
	requestid "github.com/aledenza/serco/requestId"
)

var (
	NewClient       = client.NewClient
	NewDatabase     = database.NewDatabase
	NewDbConnection = database.NewConnection
	SetRequestId    = requestid.Set
	GetRequestId    = requestid.Get
)
