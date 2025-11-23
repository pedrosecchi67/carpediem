package dbase

import (
	sql "database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// get dbase path from env. variables or local directory
func dbasePath() string {
	// env. variable
	path := os.Getenv("CARPEDIEM_PATH")
	if path != "" {
		return path
	}

	// relative path to executable
	executable, err := os.Executable()
	if err != nil {
		log.Fatal("Error finding executable file for database location")
	}
	path = filepath.Join(filepath.Dir(executable), "dbase")

	dpath := filepath.Join(path, "poetry-database.sqlite3")
	_, err = os.Stat(dpath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal(
				fmt.Sprintf("File %s does not exist.\nSet env. variable CARPEDIEM_PATH to current database path or reinstall with binary in database directory, as intended.", dpath))
		} else {
			log.Fatal(err)
		}
	}

	return dpath
}

// get path and connect to database
func Connect() *sql.DB {
	conn, err := sql.Open("sqlite3", dbasePath())
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

// define SQL query for a given author and title
func defineQuery(title string, author string, isQuery bool) (string, []any) {
	query := "SELECT * FROM Poems WHERE 1=1"

	params := make([]any, 0, 0)

	if len(title) > 0 {
		query += " AND title_noncase LIKE ?"
		params = append(params,
			fmt.Sprintf("%%%s%%", strings.ToLower(title)))
	}

	if len(author) > 0 {
		query += " AND author_noncase LIKE ?"
		params = append(params,
			fmt.Sprintf("%%%s%%", strings.ToLower(author)))
	}

	if !isQuery {
		query += " ORDER BY RANDOM() LIMIT 1"
	}

	return query, params
}

// struct to define a poem
type Poem struct {
	title  string
	author string
	poem   string
}

// run query
func QueryPoems(db *sql.DB, title string, author string, isQuery bool) []Poem {
	query, params := defineQuery(title, author, isQuery)

	rows, err := db.Query(query, params...)
	if err != nil {
		log.Fatal(err)
	}

	poems := make([]Poem, 0, 0)
	for rows.Next() {
		var poem Poem
		var id int
		var titleNonCase string
		var authorNoncase string

		if err := rows.Scan(&id, &poem.title, &poem.author, &poem.poem,
			&titleNonCase, &authorNoncase); err != nil {
			log.Fatal(err)
		}

		poems = append(poems, poem)
	}

	return poems
}

// print line for poem table summary
func (poem Poem) PrintTable() {
	fmt.Printf("%20s\t%50s\n", poem.author, poem.title)
}

// print the poem
func (poem Poem) PrintPoem() {
	fmt.Printf(
		"%s\nBy %s\n\n", poem.title, poem.author,
	)
	fmt.Println(poem.poem)
}
