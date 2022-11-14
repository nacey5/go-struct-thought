package main

import "log"

func handlerMoreResult() {
	rows, err := db.Query("SELECT * from album; SELECT * from song;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Loop through the first result set.
	for rows.Next() {
		// Handle result set.
	}

	// Advance to next result set.
	rows.NextResultSet()

	// Loop through the second result set.
	for rows.Next() {
		// Handle second set.
	}

	// Check for any error in either result set.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
