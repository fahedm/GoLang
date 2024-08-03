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

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}	

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

	albums, err := albumsByArtist("John Coltrane") // Call the albumsByArtist function
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)
}

// "go get ." to add the github.com/go-sql-driver/mysql module as a dependency for your own module. 
// Use a dot argument to mean “get dependencies for code in the current directory.”

// In command prompt, set the DBUSER and DBPASS environment variables for use
// export DBUSER=username and export DBPASS=password

// run the program, you code is connected to Database

// albumByArtist queries for almbums that have the specified artists

func albumsByArtist(name string) ([]Album, error) {
	// An album slice will hold the data from return row
	var albums []Album // albums slice of the Album type
	// This will hold data from returned rows. Struct field names and types correspond to database column names and types.

	rows, err := db.Query("select * from album where artist = ?", name) // to execute a SELECT statement 
	/* Query’s first parameter is the SQL statement. 
	After the parameter, you can pass zero or more parameters of any type. 
	These provide a place for you to specify the values for parameters in your SQL statement. 
	By separating the SQL statement from parameter values (rather than concatenating them with, say, fmt.Sprintf), 
	you enable the database/sql package to send the values separate from the SQL text, removing any SQL injection risk. */
	
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close() // Defer closing rows so that any resources it holds will be released when the function exits.
	// Lopp through rows, using Scan to assign column data to struct fields
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price ); err != nil {
			/* Scan is used to assign each row’s column values to Album struct fields.
			Scan takes a list of pointers to Go values, where the column values will be written. 
			Here, you pass pointers to fields in the alb variable, created using the & operator. 
			Scan writes through the pointers to update the struct fields. 
			err != nil > check for an error from scanning column values */
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb) // append the new alb to the albums slice.
	}
	if err := rows.Err(); err != nil {
		/* check for an error from the overall query, using rows.Err. 
		Note that if the query itself fails, checking for an error here is the only way to find out 
		that the results are incomplete. */
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}