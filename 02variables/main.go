package main

import "fmt"

//jwtToken := 30000 // walarus operators not allowed outside the method
const LoginToken string = "kfgmsmfskmfok" //First letter of variable name is capital which means its public variable
var jwtToken string = "ijdoeofkokfok"

func main() {
	var username string = "prakhar"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variables is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 254
	fmt.Println(smallVal)
	fmt.Printf("variables is of type: %T \n", smallVal)

	var smallFloat float32 = 255.4554451132232
	fmt.Println(smallFloat)
	fmt.Printf("variables is of type: %T \n", smallFloat)

	var float float64 = 255.4554451132232
	fmt.Println(float)
	fmt.Printf("variables is of type: %T \n", float)

	var intVariable int = 7
	fmt.Println(intVariable)
	fmt.Printf("variables is of type: %T \n", intVariable)

	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("variables is of type: %T \n", defaultInt)

	//implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)
	fmt.Printf("variables is of type: %T \n", website)

	//no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("variables is of type: %T \n", LoginToken)

	fmt.Println(jwtToken)
	fmt.Printf("variables is of type: %T \n", jwtToken)

}
