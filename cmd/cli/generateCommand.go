package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"portfolio-generator/cmd/structures"
	"portfolio-generator/cmd/views"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

var fileModeReadOnly fs.FileMode = 0644
var fileModeReadWrite fs.FileMode = 0755

var generateCommand = &cobra.Command{
	Use:   "generate",
	Short: "Generates html files based on input",
	Run: func(cmd *cobra.Command, args []string) {
		// Get current working directory and crash if that is not
		// available
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Could not read current working directory: %v\n", err)
		}

		// Initializing Program struct and reading flag values into it
		program := Program{
			Input:     ".",
			InputType: none,
			Output:    ".",
		}

		input := cmd.Flag("input").Value.String()
		program.Input = checkInput(input, cwd)

		output := cmd.Flag("output").Value.String()
		program.Output = checkOutput(output, cwd)

		// check input flag validity and if it's a file or directory
		fileInfo, err := os.Stat(string(program.Input))
		if err != nil {
			log.Fatalf("Error reading fileinfo of %s\n%e\n", program.Input, err)
		}

		if fileInfo.Mode().IsRegular() {
			program.InputType = file
		} else if fileInfo.Mode().IsDir() {
			program.InputType = dir
		} else {
			log.Fatalf("%s is not a file nor directory\n", program.Input)
		}

		jsonFiles := getInputFiles(program)

		// The results from the goroutines need to go somewhere
		results := make(chan string, len(jsonFiles))
		wg := sync.WaitGroup{}
		for _, fileName := range jsonFiles {
			wg.Add(1)
			go func(fileName string, channel chan<- string) {
				defer wg.Done()
				fileData, err := marshalFile(fileName)
				if err != nil {
					log.Printf("Could not open file %s\n", fileName)
					return
				}

				outputPath := filepath.Join(string(program.Output), fileData.FileName)
				outputFile, err := openDumpFile(outputPath)
				if err != nil {
					log.Printf("Could not open file %s, skipping %s\n", outputPath, fileName)
					return
				}
				defer outputFile.Close()

				err = views.Page(*fileData).Render(context.Background(), outputFile)
				if err != nil {
					log.Printf("Could render %s\n", fileName)
					return
				}
				channel <- outputPath
			}(fileName, results)
		}

		wg.Wait()
		close(results)

		// loop over results to print them to stdout
		_results := []string{}
		for result := range results {
			_results = append(_results, result)
		}

		fmt.Printf("Generated the following files:\n%s\n", strings.Join(_results, "\n"))
	},
}

// Opens a file and tries to unmarshal it into a structures.File
func marshalFile(fileName string) (*structures.File, error) {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	var fileData structures.File
	if err := json.Unmarshal(data, &fileData); err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &fileData, nil
}

// This opens a file and returns the handle to the caller
// This does not close the file!
func openDumpFile(fileName string) (*os.File, error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fileModeReadOnly)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// Check if the input is valid and return a full path if it is
func checkInput(input, cwd string) Input {
	var returnValue Input
	if input != "." {
		// Check if the input exists and is a file or not
		inputFileInfo, err := os.Stat(string(input))
		if err != nil || os.IsNotExist(err) {
			log.Fatalf("%s is not a file nor directory\n", input)
		}

		if inputFileInfo.IsDir() {
			returnValue = Input(path.Join(cwd, string(input)))
		} else {
			returnValue = Input(path.Join(cwd, string(input)))
		}
	} else {
		returnValue = Input(cwd)
	}

	return returnValue
}

// Check if output is valid and return a full path if it is
func checkOutput(output, cwd string) Output {
	var returnValue Output
	if output != "." {
		ouputFileInfo, err := os.Stat(string(output))
		if err != nil && os.IsNotExist(err) {
			if err := os.MkdirAll(string(output), fileModeReadWrite); err != nil {
				log.Fatalf("Error creating directory: %v\n", err)
			}
			ouputFileInfo, _ = os.Stat(string(output))
		}

		if !ouputFileInfo.Mode().IsDir() {
			log.Fatalf("%s is not a file nor directory\n", output)
		}

		returnValue = Output(path.Join(cwd, string(output)))
	} else {
		returnValue = Output(cwd)
	}

	return returnValue
}

func getInputFiles(program Program) []string {
	jsonFiles := []string{}
	switch program.InputType {

	case dir:
		dir, err := os.Open(string(program.Input))
		if err != nil {
			log.Fatalf("Error opening directory:%e", err)
		}
		defer dir.Close()

		files, err := dir.Readdir(-1)
		if err != nil {
			log.Fatalf("Error reading directory contents: %e", err)
		}

		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
				fileName := path.Join(string(program.Input), file.Name())
				jsonFiles = append(jsonFiles, fileName)
			}
		}
	case file:
		jsonFiles = append(jsonFiles, string(program.Input))
	}

	return jsonFiles
}
