package connection

import (
	"database/sql"
	"fmt"

	"com.derso/curso_creuto/gorm/util"
)

var (
	// Conexão
	Db *sql.DB
)

func ConnectDB() {
	host := util.GetEnv("DEMO_DBHOST", "localhost")
	dbPort := util.GetEnv("DEMO_DBPORT", "5432")
	dbUser := util.GetEnv("DEMO_DBUSER", "postgres")
	password := util.GetEnv("DEMO_DBPASSWORD", "mysecretpassword")
	dbName := util.GetEnv("DEMO_DBNAME", "postgres")
	sslMode := util.GetEnv("DEMO_DBSSL", "disable")
	connectString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, dbPort, dbUser, password, dbName, sslMode)

	var err error
	Db, err = sql.Open("postgres", connectString)

	if err != nil {
		panic(err)
	}

	defer Db.Close()

	if err = Db.Ping(); err != nil {
		panic(err)
	}
}
