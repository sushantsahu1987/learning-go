package main

import (
	"fmt"
	"sort"
)

func main() {
	// integers must be sorted
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice1)
	idx5 := sort.SearchInts(slice1, 5)
	fmt.Println("search for 5 ", idx5)

	// strings must be sorted
	slice2 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	fmt.Println(slice2)
	idxc := sort.SearchStrings(slice2, "c")
	fmt.Println("search for c ", idxc)

	slice3 := []int{6, 3, 5, 2, 8, 3, 4, 0, 9}
	sort.Ints(slice3)
	fmt.Println("sort slice 3", slice3)

	slice4 := []string{"x", "s", "a", "d", "f", "g"}
	sort.Strings(slice4)
	fmt.Println("sort slice 4", slice4)

	type Person struct {
		Name string
		Age  int
	}

	persons := []Person{
		{"Sushant", 23},
		{"Mary", 21},
		{"Spike", 25},
		{"Jane", 19},
	}
	sort.SliceStable(persons, func(i, j int) bool { return persons[i].Name < persons[j].Name })
	fmt.Println("sort persons", persons)

	searchByName := "Sushant"
	idx := sort.Search(len(persons), func(i int) bool { return persons[i].Name >= searchByName })
	if idx < len(persons) && persons[idx].Name == searchByName {
		fmt.Printf("found person with name %s at %d", searchByName, idx)
	} else {
		fmt.Println("not found person with name", searchByName)
	}

}
