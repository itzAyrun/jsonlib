package jsonlib

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/itzAyrun/jsonlib/internal/validator"
)

// ReadJSONFile reads a JSON file from the specified file path and unmarshals its contents into the provided object.
// It validates the file path to ensure it points to an accessible JSON file and handles errors during file reading and unmarshaling.
func ReadJSONFile(filePath string, obj any) error {
	if err := validator.ValidateFilePath(filePath, false); err != nil {
		return fmt.Errorf("invalid filepath: %v", err)
	}

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	if err := json.Unmarshal(fileData, obj); err != nil {
		return fmt.Errorf("error unmarshaling json: %v", err)
	}

	return nil
}

// WriteJSONFile marshals the provided object into a JSON string and writes it to the specified file path.
// It validates the file path, allowing the creation of a new file if it doesn't already exist, and handles errors during marshaling and writing the JSON data.
func WriteJSONFile(filePath string, obj any) error {
	if err := validator.ValidateFilePath(filePath, true); err != nil {
		return fmt.Errorf("invalid filepath: %v", err)
	}

	jsonData, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("error marshaling json: %v", err)
	}

	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

// ParseJSONFile reads a JSON file from the specified file path, unmarshals it's content into an object and returns it.
// The return type of the object depends on the format of the JSON file, if it's just a plain JSON object, the return type will be inferred as map[string]interface{}
// Just like ReadJSONFile, this also validates the filepath to ensure it points to an accessible JSON file and handles errors during file reading and unmarshaling
func ParseJSONFile(filePath string) (obj any, err error) {
	if err := validator.ValidateFilePath(filePath, false); err != nil {
		return nil, fmt.Errorf("invalid filepath: %v", err)
	}

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	if err := json.Unmarshal(fileData, &obj); err != nil {
		return nil, fmt.Errorf("error unmarshaling json: %v", err)
	}

	return obj, nil
}
