package main

func main() {
	tt := make(map[string]string, 0)
	tt["11"] = "a"
	tt["22"] = "b"
	tt["33"] = "c"
	tt["44"] = "d"

	for key,v := range tt {
		println(key)
		println(v)
	}
}
