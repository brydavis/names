package main

// package names

import (
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

import _ "github.com/go-sql-driver/mysql"

const connStr = ""

type NameSet struct {
	Name,
	Gender string
	Year,
	Count int
}

// type NameDict []NameSet

type Criteria struct {
	Name,
	Gender string
}

func Upload(folder string) { // (data NameDict, err error) {
	dir, err := ioutil.ReadDir(folder)
	if err != nil {
		fmt.Println(err)
	}

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into names(nm, yr, cnt, gndr) values(?, ?, ?, ?)")

	if err != nil {
		log.Fatal(err, stmt)
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
			fmt.Println(err)
		}

		for _, each := range raw {
			count, _ := strconv.Atoi(each[2])

			// fmt.Printf("name: %s, gender: %s, count: %s, year: %s\n", each[0], each[1], each[2], year)
			// name := NameSet{each[0], each[1], year, count}
			// data = append(data, name)

			_, err := stmt.Exec(
				each[0], // name
				year,
				count,
				each[1], // gender
			)

			if err != nil {
				log.Fatal(err)
			}

		}
	}

	tx.Commit()
}

func Search(name string) {
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := fmt.Sprintf("select nm, yr, cnt, gndr from `names` where nm like '%%%s%%'", name)
	fmt.Println(query)

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var results []NameSet

	for rows.Next() {
		var nset NameSet

		rows.Scan(&nset.Name, &nset.Year, &nset.Count, &nset.Gender)
		results = append(results, nset)

	}

	for _, v := range results {
		// if v.Name == name {
		fmt.Printf("name: %s, gender: %s, count: %d, year: %d\n", v.Name, v.Gender, v.Count, v.Year)
		// }
	}
}

func High(crit Criteria) (top NameSet) {

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := fmt.Sprintf("select nm, yr, cnt, gndr from `names` where nm like '%%%s%%' and gndr like '%%%s%%'", crit.Name, crit.Gender)
	fmt.Println(query)

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var results []NameSet

	for rows.Next() {
		var nset NameSet

		rows.Scan(&nset.Name, &nset.Year, &nset.Count, &nset.Gender)
		results = append(results, nset)

	}

	// for _, v := range results {
	// 	// if v.Name == name {
	// 	fmt.Printf("name: %s, gender: %s, count: %d, year: %d\n", v.Name, v.Gender, v.Count, v.Year)
	// 	// }
	// }

	// var male NameSet
	// var female NameSet
	// name := crit.Name
	// gender := crit.Gender

	for _, v := range results {
		// if v.Name == name {
		// fmt.Printf("name: %s, gender: %s, count: %s, year: %s\n", v.Name, v.Gender, v.Count, v.Year)
		// switch {
		// case v.Gender == "M" && gender == "M":
		// if v.Count > male.Count {
		// 	male.Name = v.Name
		// 	male.Gender = v.Gender
		// 	male.Count = v.Count
		// 	male.Year = v.Year
		// }

		if v.Count > top.Count {
			top.Name = v.Name
			top.Gender = v.Gender
			top.Count = v.Count
			top.Year = v.Year
		}
		// case v.Gender == "F" && gender == "F":
		// 	if v.Count > female.Count {
		// 		female.Name = v.Name
		// 		female.Gender = v.Gender
		// 		female.Count = v.Count
		// 		female.Year = v.Year
		// 	}
		// }

		// }
	}

	// 	return []NameSet{male, female}
	return
}

// func (dataset NameDict) Low(crit Criteria) []NameSet {
// 	var male NameSet
// 	var female NameSet
// 	name := crit.Name
// 	gender := crit.Gender

// 	for _, v := range dataset {
// 		if v.Name == name {
// 			switch {
// 			case v.Gender == "M" && gender == "M":
// 				if male.Count != 0 {
// 					if v.Count < male.Count {
// 						male.Name = v.Name
// 						male.Gender = v.Gender
// 						male.Count = v.Count
// 						male.Year = v.Year
// 					}
// 				} else {
// 					male.Name = v.Name
// 					male.Gender = v.Gender
// 					male.Count = v.Count
// 					male.Year = v.Year
// 				}
// 			case v.Gender == "F" && gender == "F":
// 				if female.Count != 0 {
// 					if v.Count < female.Count {
// 						female.Name = v.Name
// 						female.Gender = v.Gender
// 						female.Count = v.Count
// 						female.Year = v.Year
// 					}
// 				} else {
// 					female.Name = v.Name
// 					female.Gender = v.Gender
// 					female.Count = v.Count
// 					female.Year = v.Year
// 				}
// 			}
// 		}
// 	}

// 	return []NameSet{male, female}
// }

func main() {
	// Upload("./data")

	// Search(os.Args[1])

	n := flag.String("name", "Ruby", "Enter search name (e.g. \"John\")")
	g := flag.String("gender", "F", "Enter search gender (F)emale or (M)ale")

	flag.Parse()

	crit := Criteria{*n, *g}

	high := High(crit)
	// low := Low(crit)

	fmt.Println(high)
	// fmt.Println(low)

}
