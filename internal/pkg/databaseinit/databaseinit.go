package databaseinit

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/OPYS/internal/pkg/types"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	configPath = "../../configs/"
	ConnVar    *sqlx.DB
)

func InitDataBase() (*sqlx.DB, error) {

	var dbData types.DBData
	slbyte, _ := ioutil.ReadFile(configPath + "dbDataAccess.json")

	err := json.Unmarshal(slbyte, &dbData) // decoding JSON data into the struct
	if err != nil {
		// loging err
		return nil, errors.New("Can not decode from JSON file!")
	}

	//  dbInfo := dbData.DataBaseUser + `://` + dbData.DataBaseUser + `:` + dbData.DataBasePassword + `@192.168.202.10/` + dbData.DataBaseName
	dbInfo := `user=postgres password=admin dbname=postgres sslmode=disable`

	ConnVar = sqlx.MustConnect("postgres", dbInfo)

	if err = ping(); err != nil {
		return nil, err
	}
	fmt.Println("\n PINGED!")

	return ConnVar, nil
}

func Disconnect() {
	fmt.Println("DISCONNECTED!")
	ConnVar.Close()
}

func ping() error {
	return ConnVar.Ping()
}
