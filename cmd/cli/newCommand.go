package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"portfolio-generator/cmd/structures"

	"github.com/spf13/cobra"
)

var newCommand = &cobra.Command{
	Use:   "new",
	Short: "Generates a json file with dummy data",
	Run: func(cmd *cobra.Command, args []string) {
		// open a new file
		file, err := os.OpenFile("new.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fileModeReadOnly)
		if err != nil {
			log.Fatalf("Could not open new.json: %e", err)
		}
		defer file.Close()

		dummyData := structures.CreateExampleFile()
		jsonData, err := json.MarshalIndent(dummyData, "", "    ")
		if err != nil {
			log.Fatalf("Error marshaling JSON: %e", err)
		}

		_, err = file.Write(jsonData)
		if err != nil {
			log.Fatalf("Could not write to new.json: %e", err)
		}

		fmt.Println("Generated new.json")
	},
}
