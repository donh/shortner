package storage

import (
	"strconv"

	"github.com/donh/shortner/pkg/models"
	"github.com/donh/shortner/pkg/util"
	"github.com/jmoiron/sqlx"

	// A MySQL driver for Go's database/sql package
	_ "github.com/go-sql-driver/mysql"
)

func setDatabaseConnection() (*sqlx.DB, error) {
	databaseConfig := util.Config().Database
	account := databaseConfig.Account
	database := databaseConfig.Database
	hostname := databaseConfig.Hostname
	password := databaseConfig.Password
	port := strconv.Itoa(databaseConfig.Port)
	connection := account + ":" + password + "@(" + hostname + ":" + port + ")/" + database
	db, err := sqlx.Connect("mysql", connection)
	return db, err
}

// Query queries user table
func Query(field, value string) (a int, b, c string, e error) {
	url := ""
	short := ""

	db, err := setDatabaseConnection()
	if err != nil {
		return 0, "", "", models.ErrDatabaseError
	}

	where := "WHERE `" + field + "`=?"
	statement := "SELECT * FROM `tree`.`items` " + where + " LIMIT 1;"
	item := models.Item{}
	err = db.Get(&item, statement, value)
	if err != nil {
		return 0, "", "", err
	}
	id := int(item.ID)
	url = item.URL
	if item.Short.Valid {
		short = item.Short.String
	}
	return id, url, short, nil
}

// Insert inserts a record into user table
func Insert(url string) int {
	db, err := setDatabaseConnection()
	if err != nil {
		return 0
	}

	statement := "INSERT INTO `tree`.`items` (url) VALUES (?);"
	result, err := db.Exec(statement, url)
	if err != nil {
		return 0
	}

	id, _ := result.LastInsertId()
	if id > 0 {
		return int(id)
	}
	return 0
}

// Update updates a record in user table
func Update(id int, short string) error {
	db, err := setDatabaseConnection()
	if err != nil {
		return models.ErrDatabaseError
	}

	statement := "UPDATE `tree`.`items` SET `short`=? WHERE `id`=?;"
	result, err := db.Exec(statement, short, id)
	if err != nil {
		return models.ErrDatabaseError
	}
	RowsAffected, _ := result.RowsAffected()
	if RowsAffected > 0 {
		return nil
	}
	return models.ErrDatabaseError
}
