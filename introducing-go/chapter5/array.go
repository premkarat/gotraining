package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 83

	var total float64
	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	fmt.Println("average= ", total/float64(len(y)))
	total = 0
	for _, value := range y {
		total += value
	}
	fmt.Println("range average =", total/float64(len(y)))
	a := [5]float64{50, 10, 3, 2, 56}
	fmt.Println(a)
	//slice operation
	slice1 := make([]float64, 5, 6)
	fmt.Println("lenghth of slice1=", len(slice1))
	slice1 = append(slice1, 1, 2, 3)
	fmt.Println("lenghth of slice1=", len(slice1))
	fmt.Printf("%p\n", &slice1)
	slice1 = append(slice1, 4, 5, 5, 6, 7, 8, 9, 10, 11, 12, 12)
	fmt.Println("lenghth of slice1=", len(slice1))
	fmt.Printf("%p\n", &slice1)
	fmt.Println(slice1)
	//slice operations
	slice2 := []int{9, 8, 7, 5}
	slice3 := make([]int, 0)
	slice3 = append(slice3, 1, 2, 3, 4, 5)
	fmt.Println("slice3=", slice3)
	copy(slice2, slice3)
	fmt.Println(slice2, slice3)
	//maps
	m := make(map[int]int)
	m[1] = 10
	fmt.Println("map value=", m[1])
	//maps example 2
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	fmt.Println(elements["Li"])
	name, ok := elements["Un"]
	fmt.Println("name=", name, "ok=", ok)
	elements1 := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
	}
	if name, ok := elements1["un"]; !ok {
		fmt.Println("name=", name, "not found")
	}
	elements2 := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "Active",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "Active",
		},
	}
	if ele, ok := elements2["N"]; ok {
		fmt.Println(ele["name"], ele["state"])
	}
	elements3 := map[string]map[string]string{
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
	}
	if element, ok := elements3["O"]; ok {
		fmt.Println(element["name"], element["state"])
	}
	//storing map of string to slice of ints
	elements4 := map[string][]int{
		"marks": []int{
			12,
			25,
			14,
		},
	}
	if smarks, ok := elements4["marks"]; ok {
		fmt.Println(smarks)
	}
}
