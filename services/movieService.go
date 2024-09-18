package services

import (
	"fmt"
	"movie-rest-api/models"
	"movie-rest-api/config"
	"database/sql"
)

var db *sql.DB
var err error

func init() {
	db, err = connectDb.DbConnection()
	if err != nil {
		fmt.Println("Error connecting to database:", err.Error())
		panic(err)
	}
}

func ListMoviesHandler() ([]models.Movie, error) {
	if err != nil {
		return nil, fmt.Errorf("database connection error: %w", err)
	}

	results, err := db.Query("SELECT * FROM movies")
	if err != nil {
		return nil, fmt.Errorf("query execution error: %w", err)
	}
	defer results.Close()

	movieItems := []models.Movie{}
	for results.Next() {
		var movieFound models.Movie
		err = results.Scan(&movieFound.ID, &movieFound.Title, &movieFound.Director)
		if err != nil {
			return nil, fmt.Errorf("error scanning movie: %w", err)
		}
		movieItems = append(movieItems, movieFound)
	}

	return movieItems, nil
}

func CreateMovieHandler(movieItem models.Movie) error {
	insert, err := db.Exec("INSERT INTO movies (id, title, director) VALUES (?, ?, ?)",
		movieItem.ID, movieItem.Title, movieItem.Director)

	if err != nil {
		return fmt.Errorf("error inserting movie: %w", err)
	}

	rowsAffected, err := insert.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return fmt.Errorf("no rows affected: %w", err)
	}

	fmt.Println("Movie inserted successfully")
	return nil
}

func GetMovieById(id string) (*models.Movie, error) {
	movieItem := &models.Movie{}

	results, err := db.Query("SELECT * FROM movies WHERE id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("query execution error: %w", err)
	}
	defer results.Close()

	if results.Next() {
		err := results.Scan(&movieItem.ID, &movieItem.Title, &movieItem.Director)
		if err != nil {
			return nil, fmt.Errorf("error scanning movie: %w", err)
		}
	}
	return movieItem, nil
}
