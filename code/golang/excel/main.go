package main

import (
	"fmt"

	excel "github.com/szyhf/go-excel"
)

type Raw struct {
	// use field name as default column name
	OuterEdgeGroupID string `xlsx:"column(outer_edge_group_id)"`
}

// Run ...
func main() {

	fileName := "./test.xlsx"

	conn := excel.NewConnecter()
	err := conn.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	rd, err := conn.NewReader("Sheet1")
	if err != nil {
		panic(err)
	}
	defer rd.Close()

	for rd.Next() {
		raw := &Raw{}
		// Read a row into a struct.
		err := rd.Read(raw)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(raw)
		}
	}
}
