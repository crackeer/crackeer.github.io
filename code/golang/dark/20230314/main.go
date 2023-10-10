package main

func main() {
	/*
		result := getFullOrder("aabb")
		for _, value := range result {
			fmt.Println(value)
		}
		fmt.Println(len(result))*/
}

func getSeries(list []int) int {
	if len(list) < 1 {
		return 0
	}
}

func getFullOrder(str string) []string {
	if len(str) < 1 {
		return nil
	}
	if len(str) < 2 {
		return []string{str}
	}
	bytes := []byte(str)
	result := []string{}
	for i := 0; i < len(bytes); i++ {
		if i-1 > -1 && bytes[i] == bytes[i-1] {
			continue
		}
		bytes[0], bytes[i] = bytes[i], bytes[0]
		orders := getFullOrder(string(bytes[1:]))
		for _, item := range orders {
			tmp := make([]byte, len(item)+1)

			copy(tmp, bytes[0:1])
			copy(tmp[1:], item)
			result = append(result, string(tmp))
		}
		bytes[i], bytes[0] = bytes[0], bytes[i]
	}
	return result
}
