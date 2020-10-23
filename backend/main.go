package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/moominzie/user-record/controllers"
	_ "github.com/moominzie/user-record/docs"
	"github.com/moominzie/user-record/ent"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Facultys struct {
	Faculty []Faculty
}

type Faculty struct {
	Fname string
}

type Branchs struct {
	Branch []Branch
}

type Branch struct {
	Brname string
	id     int
}

type Buildings struct {
	Building []Building
}

type Building struct {
	Buname string
}

type Rooms struct {
	Room []Room
}

type Room struct {
	Rname string
	id    int
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:user_record.db?&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewFacultyController(v1, client)
	controllers.NewBranchController(v1, client)
	controllers.NewBuildingController(v1, client)
	controllers.NewRoomController(v1, client)
	controllers.NewUserController(v1, client)

	// Set Facultys Data
	facultys := Facultys{
		Faculty: []Faculty{
			Faculty{"Engineering"},
		},
	}
	for _, f := range facultys.Faculty {
		client.Faculty.
			Create().
			SetFname(f.Fname).
			Save(context.Background())
	}

	// Set Branchs Data
	branchs := Branchs{
		Branch: []Branch{
			Branch{"Computer Engineering", 1},
			Branch{"Civil Engineering", 1},
			Branch{"Electrical Engineering", 1},
			Branch{"Aeronautical Engineering", 2},
			Branch{"Engineering Management", 2},
			Branch{"Chemical Engineering", 2},
		},
	}
	for _, br := range branchs.Branch {
		client.Branch.
			Create().
			SetBrname(br.Brname).
			Save(context.Background())
	}

	// Set Buildings Data
	buildings := Buildings{
		Building: []Building{
			Building{"Studying building 1"},
			Building{"Studying building 2"},
		},
	}
	for _, bu := range buildings.Building {
		client.Building.
			Create().
			SetBuname(bu.Buname).
			Save(context.Background())
	}

	// Set Rooms Data
	rooms := Rooms{
		Room: []Room{
			Room{"B1120", 2001},
			Room{"B1112", 2001},
			Room{"B1215", 2001},
			Room{"B1217", 2001},
			Room{"B2502", 2002},
			Room{"B2510", 2002},
			Room{"B2512", 2002},
			Room{"B2515", 2002},
		},
	}
	for _, r := range rooms.Room {
		client.Room.
			Create().
			SetRname(r.Rname).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
