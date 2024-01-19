package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)


func main() {
	var rootCmd = &cobra.Command{Use: "golite" ,Short: "A command-line tool for GoLite framework",}

	var createMiddlewareCmd = &cobra.Command{
		Use: "make:middleware [name]",
		Short :"Create a new middleware file",
		Args:  cobra.ExactArgs(1),
        Run: func(cmd *cobra.Command, args []string) {
            // Get the middleware name from the command arguments
            name := args[0]

            // Create the middleware file with the given name
            createMiddleware(name)
        },
	}

	rootCmd.AddCommand(createMiddlewareCmd)

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}


func createMiddleware(name string) {

	//read template
	templateFile := "../../internal/middleware/template.go"

	content,err := os.ReadFile(templateFile)
	fmt.Println(content)

	if err != nil {
		fmt.Println("Error occured while reading template")
		return
	}

	
	 // Define data for template
	 data := struct {
        MiddlewareName string
    }{
        MiddlewareName: name,
    }

	tmpl,err := template.New("middleware").Parse(string(content))

	if err != nil {
		fmt.Println("Error parsing middleware template:", err)
        return
	}

	//create middleware
	fileDir  := "../../middleware/" 

	outputPath := fileDir + name + ".go"

	err = os.MkdirAll(fileDir,os.ModePerm)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error creating middleware path", fileDir)
		return
	}

	file,err := os.Create(outputPath)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error creating middleware", name)
		return
	}

	defer file.Close()

	if err := tmpl.Execute(file,data); err != nil {
		fmt.Println("Error generating middleware file:", err)
        return
	}

    fmt.Printf("Created middleware file: %s\n", outputPath)
}