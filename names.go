// package main
package names

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	// "strings"
)

type NameSet struct {
	Name,
	Gender,
	Year,
	Count string
}

type NameDict []NameSet

func Upload(folder string) (dataset NameDict, err error) {
	dir, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	for _, v := range dir {
		year := v.Name()[3:7]
		file, _ := os.Open(folder + "/" + v.Name())
		defer file.Close()

		reader := csv.NewReader(file)
		reader.FieldsPerRecord = -1 // see the Reader struct information below

		raw, err := reader.ReadAll()
		if err != nil {
			return nil, err
		}

		for _, each := range raw {
			// fmt.Printf("name: %s, gender: %s, count: %s, year: %s\n", each[0], each[1], each[2], year)
			name := NameSet{each[0], each[1], year, each[2]}
			dataset = append(dataset, name)
		}
	}
	return
}

func (dataset NameDict) Search(name string) {
	for _, v := range dataset {
		if v.Name == name {
			fmt.Printf("name: %s, gender: %s, count: %s, year: %s\n", v.Name, v.Gender, v.Count, v.Year)
		}
	}
}

// func main() {
// 	data, err := Upload("./data")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	data.Search(os.Args[1])

// 	// fmt.Println()
// }
