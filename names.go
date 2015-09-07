package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

import _ "github.com/go-sql-driver/mysql"

var connStr = ""

var query string

type NameSet struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Year   int    `json:"year"`
	Count  int    `json:"count"`
}

type Criteria struct {
	Name,
	Gender string
	Exact bool
}

func Upload(folder string) error {
	dir, err := ioutil.ReadDir(folder)
	if err != nil {
		return err
	}

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return err
	}

	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("insert into names(nm, yr, cnt, gndr) values(?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	for _, file := range dir {
		year, _ := strconv.Atoi(file.Name()[3:7])
		content, _ := os.Open(folder + "/" + file.Name())
		defer content.Close()

		reader := csv.NewReader(content)
		reader.FieldsPerRecord = -1 // see the Reader struct information below

		raw, err := reader.ReadAll()
		if err != nil {
			return err
		}

		for _, each := range raw {
			count, _ := strconv.Atoi(each[2])

			_, err := stmt.Exec(
				each[0], // name
				year,
				count,
				each[1], // gender
			)

			if err != nil {
				return err
			}
		}
	}

	tx.Commit()
	return nil
}

func Search(crit Criteria) error {
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return err
	}

	defer db.Close()

	if crit.Exact {
		query = fmt.Sprintf("select distinct nm, yr, cnt, gndr from `names` where nm = '%s' and gndr like '%%%s%%'", crit.Name, crit.Gender)
	} else {
		query = fmt.Sprintf("select distinct nm, yr, cnt, gndr from `names` where nm like '%%%s%%' and gndr like '%%%s%%'", crit.Name, crit.Gender)
	}

	rows, err := db.Query(query)
	if err != nil {
		return err
	}

	defer rows.Close()

	var results []NameSet

	for rows.Next() {
		var nset NameSet
		rows.Scan(&nset.Name, &nset.Year, &nset.Count, &nset.Gender)
		results = append(results, nset)
	}

	for _, v := range results {
		fmt.Printf("name: %s, gender: %s, count: %d, year: %d\n", v.Name, v.Gender, v.Count, v.Year)
	}

	return nil
}
