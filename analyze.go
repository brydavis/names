package main

import (
	"database/sql"
	"fmt"
	"log"
	"sort"
)

import _ "github.com/go-sql-driver/mysql"

func Min(array interface{}) float64 {
	a := Floatize(array)
	low := a[0]
	for _, val := range a[1:] {
		if low > val {
			low = val
		}
	}
	return low
}

func Floatize(input interface{}) (output []float64) {
	switch t := input.(type) {
	case []int:
		for _, val := range t {
			output = append(output, float64(val))
		}
	case int:
		output = append(output, float64(t))
	case []float64:
		output = t
	case float64:
		output = append(output, t)
	case []interface{}:
		for _, val := range t {
			switch tt := val.(type) {
			case int:
				output = append(output, float64(tt))
			case float64:
				output = append(output, tt)
			}
		}
	}
	return
}

func Max(array interface{}) float64 {
	var high float64
	for _, val := range Floatize(array) {
		if high < val {
			high = val
		}
	}
	return high
}

func Sum(array interface{}) (total float64) {
	for _, val := range Floatize(array) {
		total += val
	}
	return
}

func Avg(array interface{}) float64 {
	a := Floatize(array)
	return Sum(a) / float64(len(a))
}

func Med(array interface{}) float64 {
	a := Floatize(array)
	ln := len(a)
	sort.Float64s(a)
	if ln%2 == 0 {
		hi := a[ln/2]
		lo := a[(ln/2)-1]
		return Avg([]float64{hi, lo})
	} else {
		return a[ln/2]
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

	mnames := make(map[string]NameSet)

	for _, nset := range output {
		m, exists := mnames[nset.Name]

		if exists {
			if nset.Count > m.Count {
				mnames[nset.Name] = nset
			}
		} else {
			mnames[nset.Name] = nset
		}
	}

	for _, nset := range mnames {
		results = append(results, nset)
	}

	// j, _ := json.Marshal(results)
	// fmt.Printf("%v\n\n", string(j))

	return
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

	mnames := make(map[string]NameSet)

	for _, nset := range output {
		m, exists := mnames[nset.Name]

		if exists {
			if nset.Count < m.Count {
				mnames[nset.Name] = nset
			}
		} else {
			mnames[nset.Name] = nset
		}
	}

	for _, nset := range mnames {
		results = append(results, nset)
	}

	// j, _ := json.Marshal(results)
	// fmt.Printf("%v\n\n", string(j))

	return
}

func Seqs(series []int) {
	sort.Ints(series)
	start := series[0]
	current := series[0]

	for _, val := range series {
		if val == current {
			current += 1
		} else {
			fmt.Printf("start %d => end %d\n", start, current-1)
			start = val
			current = val + 1
		}
	}
	fmt.Printf("start %d => end %d\n", start, current-1)
}

func Gaps(series []int) {
	sort.Ints(series)
	current := series[0]

	for _, val := range series {
		if val == current {
			current += 1
		} else {
			fmt.Printf("start %d => end %d\n", current, val-1)
			current = val + 1
		}
	}
}

func Upper(series []int) float64 {
	sort.Ints(series)
	floats := Floatize(series)
	return floats[len(floats)-1]
}

func Lower(series []int) float64 {
	sort.Ints(series)
	floats := Floatize(series)
	return floats[0]
}

func Distinct(series []int) (unique []float64) {
	sort.Ints(series)
	floats := Floatize(series)
	latest := floats[0]
	unique = append(unique, latest)

	for _, s := range floats {
		if s > latest {
			unique = append(unique, s)
			latest = s
		}
	}
	return
}
