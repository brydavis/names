package main

import "fmt"

func main() {

	// connStr = os.Getenv("CONN_STR")

	// Upload("./data")

	// Search(os.Args[1])

	// n := flag.String("name", "Ruby", "Enter search name (e.g. \"John\")")
	// g := flag.String("gender", "F", "Enter search gender () (F)emale or (M)ale")
	// b := flag.Bool("exact", false, "")
	// flag.Parse()

	// crit := Criteria{*n, *g, *b}

	// report := make(map[string][]NameSet)

	// report["high"] = High(crit)
	// report["low"] = Low(crit)

	// j, _ := json.Marshal(report)

	// fmt.Println(string(j))

	// a := []float64{5.3, 14.9, 1.4, 2., 3.2}
	// sort.Float64s(a)
	// fmt.Println(a)
	// fmt.Println(Min(a))
	// fmt.Println(Max(a))
	// fmt.Println(Sum(a))
	// fmt.Println(Avg(a))
	// fmt.Println(Med(a))

	// b := []int{34, 63, 25, 67, 1}
	// fmt.Println(Floatize(b), Min(b)+3.4)
	// fmt.Println(Min(b))
	// fmt.Println(Max(b))
	// fmt.Println(Sum(b))
	// fmt.Println(Avg(b))
	// fmt.Println(Med(b))

	// c := []interface{}{34, 11.1, 63.3, 25, 67.1}
	// fmt.Println(Floatize(c), Max(c)+10)
	// fmt.Println(Min(c))
	// fmt.Println(Max(c))
	// fmt.Println(Sum(c))
	// fmt.Println(Avg(c))
	// fmt.Println(Med(c))

	// d := []interface{}{65.8, 63.3, 60.0, 67.1}
	// fmt.Println(Floatize(d), Max(d)+10)
	// fmt.Println(Min(d))
	// fmt.Println(Max(d))
	// fmt.Println(Sum(d))
	// fmt.Println(Avg(d))
	// fmt.Println(Med(d))

	yrs := []int{2010, 2000, 2001, 2002, 2025, 2000, 2004, 2000, 2000, 2004, 2004, 2007, 2011, 2002, 2000, 2001, 2011, 2008, 2009, 2010, 2020, 2021, 2003, 2009, 2008, 2016, 2017, 2019, 2020, 2021}

	// cp := make([]int, len(yrs))
	// copy(cp, yrs)
	// fmt.Println(cp)

	fmt.Println(Lower(yrs))
	fmt.Println(Upper(yrs))

}
