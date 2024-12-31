package validator

import (
    "path/filepath"
	"fmt"
	"os"
)

// ValidateFilePath checks if the given file path is valid, accessible, and points to a JSON file.
// If the file does not exist and creation is allowed, it will not return an error.
// Returns an error if the path is invalid, inaccessible, not a JSON file, or if it's a directory.
func ValidateFilePath(filePath string, allowCreate bool) error {
    // Check if the file exists and is accessible
    info, err := os.Stat(filePath)
    if err != nil {
        if os.IsNotExist(err) && allowCreate {
            // If the file doesn't exist and we allow creation, skip further checks
            return nil
        }
        return fmt.Errorf("could not access file: %v", err)
    }
    
    // Ensure the path is not a directory
    if info.IsDir() {
        return fmt.Errorf("provided path is a directory, not a file")
    }

    // Check if the file has a '.json' extension
    if filepath.Ext(filePath) != ".json" {
        return fmt.Errorf("file is not a valid JSON file (must have .json extension)")
    }

    return nil

}
