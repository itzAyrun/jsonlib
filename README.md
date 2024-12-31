<p align="center">
    <img src="./assets/jsonlib.png" alt="jsonlib logo"/>
</p>
<p align="center"><i>A <a href="https://go.dev">Go</a> library for effortless JSON handling and manipulation</i></p>

<!-- Insert shields badges here -->

## Overview

`jsonlib` is a Go package that provides functions for reading from and writing to JSON files. It simplifies handling JSON data by offering easy-to-use methods to load and save JSON content. Whether you're working with configuration files or data serialization, `jsonlib` aims to make JSON handling in Go straightforward.

## Features

- **Read JSON Files**: Easily read and parse JSON files into Go structures.
- **Write JSON Files**: Serialize Go structures into JSON format and save them to files.
- **Pretty-Printing Support (Coming Soon)**: Write JSON data with proper formatting for better readability.
- **Error Handling**: Includes error handling for common file and JSON-related issues

## Installation

To use `jsonlib` in your Go project, simply run:

```bash
go get github.com/itzAyrun/jsonlib
```

## Usage

### Reading JSON Files

Suppose we have a JSON file `config.json`:

```json
{
    "name": "jsonlib",
    "version": "alpha"
}
```

We can parse the data of this json file in our go code like this:

```go
package main

import (
    "fmt"
    "log"

    "github.com/itzAyrun/jsonlib"
)

type Config struct {
    Name    string `json:"name"`
    Version string `json:"version"`
}

func main() {
    var config Config
    err := jsonlib.ReadJSONFile("config.json", &config)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Config: %+v\n", config)
}
```

### Writing JSON Files

Let's continue with our example from above:

```go
package main

import (
    "log"

    "github.com/itzAyrun/jsonlib"
)

type Config struct {
    Name    string `json:"name"`
    Version string `json:"version"`
}

func main() {
    config := Config{
        Name:    "jsonlib",
        Version: "alpha",
    }

    err := jsonlib.WriteJSONFile("config.json", config)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Config saved successfully!")
}
```

## License

`jsonlib` is realised under the MIT License. See the [LICENSE](./LICENSE) file for more details.

## Contributing

Contributions are welcome! If you'd like to contribute to jsonlib, feel free to fork the repository and submit a pull request.
