package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	//Connection properties
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	//Get database handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Database Connected!")

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", createAlbum)
	router.GET("/album/:id", getAlbumById)
	router.Run("localhost:8000")
}

// API get all
func getAlbums(c *gin.Context) {
	//Initiate return data
	var albums []album

	//Query
	rows, err := db.Query("SELECT * FROM album")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": err})
		return
	}
	//Close connection after finished
	defer rows.Close()

	//Loop through rows
	for rows.Next() {
		var alb album
		//using Rows.Scan to assign each row’s column values to Album struct fields
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": err})
			return
		}
		//Append Data
		albums = append(albums, alb)
	}

	c.IndentedJSON(http.StatusOK, albums)
}

// API create album
func createAlbum(c *gin.Context) {
	var newAlbum album
	//Call bindJSON to bind received JSON
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", newAlbum.Title, newAlbum.Artist, newAlbum.Price)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": err})
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": err})
	}

	c.IndentedJSON(http.StatusOK, id)
}

// API get specific album by id
func getAlbumById(c *gin.Context) {
	//Get parameter
	param := c.Param("id")
	id, _ := strconv.ParseInt(param, 10, 64)

	// Var
	var alb album
	//Query single data
	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)

	// Scan data to copy column to struct
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		// Check no data found
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"Message": err})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": err})
	}

	c.IndentedJSON(http.StatusOK, alb)
}
