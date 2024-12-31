package jsonlib

import (
	"fmt"
	"log"
	"os"
)

func ExampleReadJSONFile() {
    // Define a struct to hold the JSON data
    type JSONData struct {
        Name string `json:"name"`
        Age uint `json:"age"`
        Email string `json:"email"`
    }

    // Create a temporary JSON file for the example
    jsonData := `{"name": "Alice", "age": 30, "email": "alice@example.com"}`
    
    tmpFile, err := os.CreateTemp("", "example-*.json")
    if err != nil {
        log.Fatalf("error creating temp file: %v", err)
        return
    }
    defer os.Remove(tmpFile.Name()) // Clean up the file after the example

    _, _ = tmpFile.WriteString(jsonData)
    _ = tmpFile.Close()

    // Use the ReadJSONFile function to read the JSON data
    var data JSONData
    if err := ReadJSONFile(tmpFile.Name(), &data); err != nil {
        log.Fatalf("error reading JSON file: %v", err)
        return
    }

    // Print the result
    fmt.Println("Name:", data.Name)
    fmt.Println("Age:", data.Age)
    fmt.Println("Email:", data.Email)

    // Output:
    // Name: Alice
    // Age: 30
    // Email: alice@example.com
}

func ExampleWriteJSONFile() {
    // Define a struct to hold the JSON data
    type JSONData struct {
        Name  string `json:"name"`
        Age   uint   `json:"age"`
        Email string `json:"email"`
    }

    // Create an instance of JSONData
    data := JSONData{
        Name:  "Alice",
        Age:   30,
        Email: "alice@example.com",
    }

    // Create a temporary file for the example
    tmpFile, err := os.CreateTemp("", "example-*.json")
    if err != nil {
        log.Fatalf("error creating temp file: %v", err)
        return
    }
    defer os.Remove(tmpFile.Name()) // Clean up the file after the example

    // Use the WriteJSONFile function to write the JSON data to the file
    if err := WriteJSONFile(tmpFile.Name(), data); err != nil {
        log.Fatalf("error writing JSON file: %v", err)
        return
    }

    // Open the file to read its contents and confirm the write
    fileContent, err := os.ReadFile(tmpFile.Name())
    if err != nil {
        log.Fatalf("error reading the written file: %v", err)
        return
    }

    // Print the file content
    fmt.Println("Written JSON content:", string(fileContent))

    // Output:
    // Written JSON content: {"name":"Alice","age":30,"email":"alice@example.com"}
}
