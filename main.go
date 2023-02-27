package main

import (
	"database/sql"
	"database/sql/driver"
	"log"

	"github.com/labstack/echo"

	"github.com/elin/application-server/user"

	mysql "github.com/go-sql-driver/mysql"
)

func main() {
	e := echo.New()

	conn := getDBConn()
	db := sql.OpenDB(conn)
	defer db.Close()

	ug := e.Group("/users")
	ur := user.NewUserRepository(db)
	_ = user.NewUserController(ug, ur)

	e.Logger.Fatal(e.Start(":8888"))
}

func getDBConn() driver.Connector {
	mysqlConn, err := mysql.NewConnector(&mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return mysqlConn
}
