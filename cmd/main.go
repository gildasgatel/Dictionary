package main

import (
	"Dictionary/models/rows"
	"Dictionary/pkg/dictionary"
	"flag"
	"fmt"
	"log"
)

func main() {
	list := flag.Bool("list", false, "descrpition")
	add := flag.Bool("add", false, "descrpition")
	get := flag.Bool("get", false, "descrpition")
	update := flag.Bool("update", false, "descrpition")
	delete := flag.Bool("delete", false, "descrpition")

	flag.Parse()

	var rows rows.Rows
	if items := flag.Args(); len(items) > 0 {
		rows.Key = append(rows.Key, []byte(items[0])...)
	}
	if items := flag.Args(); len(items) > 1 {
		rows.Desc = append(rows.Desc, []byte(items[1])...)
	}

	app, err := dictionary.New()
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case *list:
		datas, err := app.List()
		if err != nil {
			log.Println(err)
		}
		fmt.Println("* * * Contents * * * ")
		for _, d := range datas {
			fmt.Println(d.String())
		}
		fmt.Println("* * * * * * * * * * * ")

	case *add:
		err := app.Add(&rows)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("* * * Add succed * * * ")
	case *get:
		err := app.Get(&rows)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("* * * Get * * * ")
		fmt.Println(rows.String())
	case *update:
		err := app.Update(&rows)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("* * * Update succed * * * ")

	case *delete:
		err := app.Delete(&rows)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("* * * Delete succed * * * ")
	default:
		fmt.Println(help)
	}

}

var help = `
Call an action to use Dictionary.
 * * * * * * * * * * * * 
-list                     // list all data saved.
-add <key> "description"  // save data
-get <key>                // get data by key
-Update <key>             // update by key
-Delete <key>             // delete by key
* * * * * * * * * * * * 
`
