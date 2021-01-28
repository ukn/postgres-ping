package main

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	var db *sql.DB
	var err error
	connectionFailed := true
	errorMessage := errors.New("")
	verbose := false
	var interval time.Duration = 2

	user := os.Getenv("user")
	if user == "" {
		user = "postgres"
	}

	pass := os.Getenv("pass")

	host := os.Getenv("host")
	if host == "" {
		host = "127.0.0.1"
	}

	dbName := os.Getenv("dbname")
	if dbName == "" {
		dbName = "postgres"

	}

	dbPort := os.Getenv("port")
	if dbPort == "" {
		dbPort = "5432"

	}

	sslMode := os.Getenv("sslmode")
	if sslMode == "" {
		sslMode = "disable"
	}

	connectTimeout := os.Getenv("timeout")
	if connectTimeout == "" {
		connectTimeout = "2"
	}

	if _, exists := os.LookupEnv("verbose"); exists {
		verbose, err = strconv.ParseBool(os.Getenv("verbose"))
		if err != nil {
			log.Fatalln("verbose variable has to be set to true, false, 1 or 0.\nError:", err)
		}
	}

	if _, exists := os.LookupEnv("interval"); exists {
		intInterval, err := strconv.Atoi(os.Getenv("interval"))
		interval = time.Duration(intInterval)
		if err != nil {
			log.Fatalln("interval variable has to be an integer.\nError:", err)
		}
	}

	// connection string
	connStr := "postgres://" + user + ":" + pass + "@" + host + ":" + dbPort + "/" + dbName + "?sslmode=" + sslMode + "&connect_timeout=" + connectTimeout
	log.Println("Connection string:", "postgres://"+user+":***@"+host+"/"+dbName+"?sslmode="+sslMode+"&connect_timeout="+connectTimeout)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	for {
		if err = db.Ping(); err != nil {
			if (connectionFailed && err.Error() != errorMessage.Error()) || verbose {
				log.Println("Connection failed to", host, "ERROR:", err)
			}
			connectionFailed = true
			errorMessage = err

		} else {
			if (connectionFailed && err == nil) || verbose {
				log.Println("Postrgres ping", host, "successful")
			}
			connectionFailed = false
		}
		time.Sleep(interval * time.Second)

	}

}
