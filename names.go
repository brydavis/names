package main

// package names

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

import _ "github.com/go-sql-driver/mysql"

const connStr = ""

var query string

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
	Exact bool
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

func Search(crit Criteria) {
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if crit.Exact {
		query = fmt.Sprintf("select distinct nm, yr, cnt, gndr from `names` where nm = '%s' and gndr like '%%%s%%'", crit.Name, crit.Gender)
	} else {
		query = fmt.Sprintf("select distinct nm, yr, cnt, gndr from `names` where nm like '%%%s%%' and gndr like '%%%s%%'", crit.Name, crit.Gender)
	}

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

func High(crit Criteria) (results []NameSet) {

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if crit.Exact {
		query = fmt.Sprintf("select distinct nm, yr, cnt, gndr from `names` where nm = '%s' and gndr like '%%%s%%'", crit.Name, crit.Gender)
	} else {
		query = fmt.Sprintf("select distinct nm, yr, cnt, gndr from `names` where nm like '%%%s%%' and gndr like '%%%s%%'", crit.Name, crit.Gender)
	}

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var output []NameSet

	for rows.Next() {
		var nset NameSet

		rows.Scan(&nset.Name, &nset.Year, &nset.Count, &nset.Gender)
		output = append(output, nset)

	}

	// for _, v := range output {
	// 	// if v.Name == name {
	// 	fmt.Printf("name: %s, gender: %s, count: %d, year: %d\n", v.Name, v.Gender, v.Count, v.Year)
	// 	// }
	// }

	// var male NameSet
	// var female NameSet
	// name := crit.Name
	// gender := crit.Gender

	mnames := make(map[string]NameSet)

	for _, nset := range output {

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

		m, exists := mnames[nset.Name]

		if exists {
			if nset.Count > m.Count {
				mnames[nset.Name] = nset
			}
		} else {

			mnames[nset.Name] = nset
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

	for _, nset := range mnames {
		results = append(results, nset)
	}

	j, _ := json.Marshal(results)
	fmt.Printf("%v\n\n", string(j))

	return results
}

func Low(crit Criteria) (results []NameSet) {

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if crit.Exact {
		query = fmt.Sprintf("select distinct nm, yr, cnt, gndr from `names` where nm = '%s' and gndr like '%%%s%%'", crit.Name, crit.Gender)
	} else {
		query = fmt.Sprintf("select distinct nm, yr, cnt, gndr from `names` where nm like '%%%s%%' and gndr like '%%%s%%'", crit.Name, crit.Gender)
	}

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var output []NameSet

	for rows.Next() {
		var nset NameSet

		rows.Scan(&nset.Name, &nset.Year, &nset.Count, &nset.Gender)
		output = append(output, nset)

	}

	// for _, v := range output {
	// 	// if v.Name == name {
	// 	fmt.Printf("name: %s, gender: %s, count: %d, year: %d\n", v.Name, v.Gender, v.Count, v.Year)
	// 	// }
	// }

	// var male NameSet
	// var female NameSet
	// name := crit.Name
	// gender := crit.Gender

	mnames := make(map[string]NameSet)

	for _, nset := range output {

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

		m, exists := mnames[nset.Name]

		if exists {
			if nset.Count < m.Count {
				mnames[nset.Name] = nset
			}
		} else {

			mnames[nset.Name] = nset
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

	for _, nset := range mnames {
		results = append(results, nset)
	}

	j, _ := json.Marshal(results)
	fmt.Printf("%v\n\n", string(j))

	return results
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
	b := flag.Bool("exact", false, "")

	flag.Parse()

	crit := Criteria{*n, *g, *b}

	high := High(crit)
	low := Low(crit)

	fmt.Println(high)
	fmt.Println(low)

}
