package main

var base62 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func getShortURL(length int, id int) string {
	currLen := 1
	shortURL := ""
	for currLen <= length {
		remainder := id % 62
		shortURL = string(base62[remainder]) + shortURL
		id = id / 62
		currLen++
	}

	return shortURL
}

func main() {

}
