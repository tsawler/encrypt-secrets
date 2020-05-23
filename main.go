package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"io/ioutil"
	"log"
	"os"
)

var app *application
var key string

type application struct {
	db *sql.DB
}

// openDB opens a database connection for postgres
func openDB(dsn, dbType string) (*sql.DB, error) {
	if dbType == "postgres" {
		d, err := sql.Open("pgx", dsn)
		if err != nil {
			panic(err)
		}
		return d, err
	} else {
		d, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}
		return d, err
	}
}

// main is main app function
func main() {
	// create loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	dbHost := "127.0.0.1"
	dbPort := "5432"

	// read flags
	dbUser := flag.String("u", "", "DB Username")
	dbPass := flag.String("p", "", "DB Password")
	dbSsl := flag.String("s", "disable", "SSL Settings")
	databaseName := flag.String("db", "", "Database name")
	dbType := flag.String("dbtype", "", "Database type (postgres or mysql")
	keyPtr := flag.String("key", "", "Secret key (32 chars)")
	flag.Parse()

	if *dbUser == "" || *databaseName == "" || *dbType == "" || *dbSsl == "" {
		fmt.Println("Missing required flags.")
		os.Exit(1)
	}

	if *keyPtr != "" {
		key = *keyPtr
	} else {
		// create 32 char signer secret
		key = RandomString(32)
	}

	// write it to stdout, and save it to a file
	infoLog.Println("Key is", key)
	err := ioutil.WriteFile("./urlSignerSecret.txt", []byte(key), 0644)
	if err != nil {
		errorLog.Fatal(err)
	}

	dsn := ""

	if *dbType == "postgres" {
		if *dbPass == "" {
			dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s timezone=UTC connect_timeout=5", dbHost, dbPort, *dbUser, *databaseName, *dbSsl)
		} else {
			dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s timezone=UTC connect_timeout=5", dbHost, dbPort, *dbUser, *dbPass, *databaseName, *dbSsl)
		}
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&tls=%s&collation=utf8_unicode_ci&timeout=5s&readTimeout5s", *dbUser, *dbPass, dbHost, "3306", *databaseName, *dbSsl)
	}

	// open connection to db
	db, err := openDB(dsn, *dbType)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		errorLog.Fatal(err)
	}
	infoLog.Println("Pinged database successfully!")

	// populate config
	app = &application{
		db: db,
	}

	secrets := []string {
		"do-secret",
		"mailchimp-key",
		"pusher-secret",
		"recaptcha-secret",
		"stripe-secret",
		"stripe-test-secret",
		"github-secret",
		"facebook-secret",
		"google-secret",
		"recaptcha-secret",
	}

	for _, x := range secrets {
		if *dbType == "postgres" {
			infoLog.Println("Doing", x)
			err = app.updateSecret(x)
			if err != nil {
				errorLog.Fatal(err)
			}
		} else {
			infoLog.Println("Doing", x)
			err = app.updateSecretMysql(x)
			if err != nil {
				errorLog.Fatal(err)
			}
		}
	}


	infoLog.Println("Done!")
}

