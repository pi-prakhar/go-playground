package main

import "fmt"

func main() {
	fmt.Println("Welcome to session on maps")
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println(languages)
	fmt.Println("Js is ", languages["JS"])

	delete(languages, "RB")

	fmt.Println("List of all languages: ", languages)

	for key, value := range languages {
		fmt.Println(key, " : ", value)
	}

}
