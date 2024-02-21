package main

import "portfolio-generator/cmd/cli"

func main() {
	// str := structures.CreateExampleFile()
	// jsonData, err := json.MarshalIndent(str, "", "    ")
	// if err != nil {
	// 	fmt.Println("Error marshaling JSON:", err)
	// 	return
	// }
	//
	// // Print the JSON data
	// fmt.Println(string(jsonData))
	cli.Execute()
}
