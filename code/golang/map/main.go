package main

import "fmt"

func main() {
	var tmp map[string]interface{} = map[string]interface{}{
		"foo":    "sss",
		"SSS":    99,
		"SJSAHS": 362,
		"6763":   "SHKAJHS",
		"SHASH":  787,
	}

	for key := range tmp {
		fmt.Println(key)
	}
	fmt.Println("--------")
	for key := range tmp {
		fmt.Println(key)
	}

}
