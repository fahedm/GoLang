package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
) // Import the MySQL driver github.com/go-sql-driver/mysql.

var db *sql.DB /* Making db a global variable simplifies this example. 
				In production, you’d avoid the global variable, 
				such as by passing the variable to functions that need it or by wrapping it in a struct.*/

func main() {
	// Capture connection properties.
	cfg := mysql.Config{  // to collect connection properties
		User : os.Getenv("DBUSER"),
		Passwd : os.Getenv("DBPASS"),
		Net : "tcp",
		Addr : "127.0.0.1:3306",
		DBName : "recordings",
	}
	//The Config struct makes for code that’s easier to read than a connection string would be.

	// Database Handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	// format properties into a DSN for a connection string
	// initialize the db variable, passing the return value of FormatDSN.
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping() // DB.Ping to confirm that connecting to the database works
	if pingErr!= nil {
		log.Fatal(pingErr)
	}
	fmt.Println("connected")
}

// "go get ." to add the github.com/go-sql-driver/mysql module as a dependency for your own module. 
// Use a dot argument to mean “get dependencies for code in the current directory.”

// In command prompt, set the DBUSER and DBPASS environment variables for use
// export DBUSER=username and export DBPASS=password

// run the program, you code is connected to Database