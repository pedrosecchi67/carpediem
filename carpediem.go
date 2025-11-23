package main

import (
	dbase "carpediem/dbase"
	"flag"
	"fmt"
)

func main() {
	conn := dbase.Connect()
	defer conn.Close()

	author := flag.String("author", "", "Author name for querying")
	title := flag.String("title", "", "Title for querying")
	query := flag.Bool("query", false, "If provided, makes a query for poem titles instead of printing a random poem. Organizes them in table format.")

	flag.Parse()

	poems := dbase.QueryPoems(conn, *title, *author, *query)

	if *query {
		fmt.Printf(
			"%d poems found.\n", len(poems),
		)

		if len(poems) > 0 {
			fmt.Printf("%20s\t%50s\n", "POET", "TITLE")

			for _, poem := range poems {
				poem.PrintTable()
			}
		}
	} else {
		if len(poems) == 0 {
			fmt.Println("No poems found!")
		} else {
			poems[0].PrintPoem()
		}
	}
}
