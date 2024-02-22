package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"portfolio-generator/cmd/structures"

	"github.com/spf13/cobra"
)

// This command generates a new json config file for convenience
var newCommand = &cobra.Command{
	Use:   "new",
	Short: "Generates a json file with dummy data",
	Run: func(cmd *cobra.Command, args []string) {
		// open a new file
		file, err := os.OpenFile("new.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fileModeReadOnly)
		if err != nil {
			log.Fatalf("Could not open new.json: %e", err)
		}
		// close file on function return
		defer file.Close()

		// generate new config file struct
		dummyData := structures.CreateExampleFile()

		// marshal struct to json
		jsonData, err := json.MarshalIndent(dummyData, "", "    ")
		if err != nil {
			log.Fatalf("Error marshaling JSON: %e", err)
		}

		// write json data to file
		_, err = file.Write(jsonData)
		if err != nil {
			log.Fatalf("Could not write to new.json: %e", err)
		}

		fmt.Println("Generated new.json")
	},
}
