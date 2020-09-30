package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/scaleshift/scaleshift/api/src/db"
	"github.com/scaleshift/scaleshift/api/src/generated/restapi"
	"github.com/scaleshift/scaleshift/api/src/generated/restapi/operations"
	logs "github.com/scaleshift/scaleshift/api/src/log"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
)

func main() {
	logs.Debug("start scaleshift", nil, nil)
	db.Initialize()

	// ----------------------------------------------------------------------------------------
	//  Copied from api/src/generated/cmd/scale-shift-server/main.go
	// ----------------------------------------------------------------------------------------
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewScaleShiftAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "ScaleShift"
	parser.LongDescription = "A platform for machine learning & high performance computing\n"

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
	// ----------------------------------------------------------------------------------------
}
