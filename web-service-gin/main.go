package main // A standalone program (as opposed to a library) is always in package main.

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	/* Struct tags such as json:"artist" specify what a field’s name should be when the struct’s contents are serialized into JSON. 
	Without them, the JSON would use the struct’s capitalized field names – a style not as common in JSON. */
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
    router := gin.Default() // Initialize a Gin router using Default.
    router.GET("/albums", getAlbums) // to associate the GET HTTP method and /albums path with a handler function.
	// passing the name of the getAlbums function. This is different from passing the result of the function
    router.GET("/albums/:id", getAlbumByID)
	// In Gin, the colon preceding an item in the path signifies that the item is a path parameter.
	router.POST("/albums", postAlbums)
	// With Gin, you can associate a handler with an HTTP method-and-path combination. 
	// In this way, you can separately route requests sent to a single path based on the method the client is using.

	router.Run("localhost:8080") // to attach the router to an http.Server and start the server.
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	// gin.Context is the most important part of Gin. 
	// It carries request details, validates and serializes JSON, and more. 
    c.IndentedJSON(http.StatusOK, albums) // to serialize the struct into JSON and add it to the response.
	/* The function’s first argument is the HTTP status code you want to send to the client. 
	Here, you’re passing the StatusOK constant from the net/http package to indicate 200 OK.
	Note that you can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON. 
	In practice, the indented form is much easier to work with when debugging 
	and the size difference is usually small.

	*/
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum) // Append the album struct initialized from the JSON to the albums slice.
    c.IndentedJSON(http.StatusCreated, newAlbum)
	// Add a 201 status code to the response, along with JSON representing the album you added.
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}