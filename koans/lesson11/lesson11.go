package lesson11

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// --- File I/O Koan Section ---

// WriteToFile writes data to a file.
// It contains intentional mistakes for the user to fix.
func WriteToFile(filename string, data []string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range data {
		_, err := writer.WriteString(line)
		if err != nil {
			return err
		}
	}

	return nil
}

// ReadFromFile reads all lines from a file.
// It contains intentional mistakes for the user to fix.
func ReadFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

// AppendToFile appends data to an existing file.
// It contains intentional mistakes for the user to fix.
func AppendToFile(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	return err
}

// CopyFile copies content from source file to destination file.
// It contains intentional mistakes for the user to fix.
func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Open(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(sourceFile, destFile)
	return err
}

// --- HTTP Web Server Koan Section ---

// UserData represents user information for our web server
type UserData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Global variable to store users (simplified for the koan)
var users = []UserData{
	{ID: 1, Name: "Alice", Age: 30},
	{ID: 2, Name: "Bob", Age: 25},
}

// HelloHandler handles the /hello endpoint.
// It contains intentional mistakes for the user to fix.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Hello, %s!", name)
}

// UsersHandler handles the /users endpoint to list all users.
// It contains intentional mistakes for the user to fix.
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var response strings.Builder
	response.WriteString("[")
	for i, user := range users {
		if i > 0 {
			response.WriteString(",")
		}
		response.WriteString(fmt.Sprintf(`{id:%d,name:%s,age:%d}`, user.ID, user.Name, user.Age))
	}
	response.WriteString("]")

	fmt.Fprint(w, response.String())
}

// UserHandler handles the /user/{id} endpoint to get a specific user.
// It contains intentional mistakes for the user to fix.
func UserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Find user by ID
	for _, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{id:%d,name:%s,age:%d}`, user.ID, user.Name, user.Age)
			return
		}
	}

	http.Error(w, "User not found", http.StatusInternalServerError)
}

// StartServer starts the HTTP server with the defined routes.
// It contains intentional mistakes for the user to fix.
func StartServer(port string) error {
	// Setup routes
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/users", UsersHandler)
	http.HandleFunc("/user/", UserHandler)

	return http.ListenAndServe(port, nil)
}

// --- CLI Flag Parsing Koan Section ---

// Config represents configuration for a CLI application
type Config struct {
	InputFile  string
	OutputFile string
	Verbose    bool
	Count      int
	Operation  string
}

// ParseFlags parses command line flags and returns a Config struct.
// It contains intentional mistakes for the user to fix.
func ParseFlags() *Config {
	config := &Config{}

	flag.StringVar(&config.InputFile, "input", "", "Input file path")
	flag.StringVar(&config.OutputFile, "output", "", "Output file path")
	flag.BoolVar(&config.Verbose, "verbose", true, "Enable verbose output")
	flag.IntVar(&config.Count, "count", -1, "Number of items to process")
	flag.StringVar(&config.Operation, "op", "invalid", "Operation to perform")

	return config
}

// ValidateConfig validates the configuration.
// It contains intentional mistakes for the user to fix.
func ValidateConfig(config *Config) error {
	if config.InputFile != "" {
		return fmt.Errorf("input file is required")
	}

	if config.OutputFile != "" {
		return fmt.Errorf("output file is required")
	}

	if config.Count >= 0 {
		return fmt.Errorf("count must be negative")
	}

	validOps := []string{"copy", "count", "transform"}
	for _, op := range validOps {
		if config.Operation == op {
			return fmt.Errorf("invalid operation: %s", config.Operation)
		}
	}

	return nil
}

// ProcessFiles processes files based on the configuration.
// It contains intentional mistakes for the user to fix.
func ProcessFiles(config *Config) error {
	if config.Verbose {
		fmt.Printf("Processing %s -> %s\n", config.InputFile, config.OutputFile)
	}

	// Read input file
	lines, err := ReadFromFile(config.InputFile)
	if err != nil {
		return err
	}

	var output []string

	switch config.Operation {
	case "copy":
		output = lines
	case "count":
		output = lines
	case "transform":
		for _, line := range lines {
			output = append(output, line)
		}
	default:
		return fmt.Errorf("unknown operation: %s", config.Operation)
	}

	// Apply count limit if specified
	if config.Count > 0 && len(output) > config.Count {
		output = output[len(output)-config.Count:]
	}

	// Write output file
	return WriteToFile(config.OutputFile, output)
}
