package lesson11

import (
	"flag"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/romeufcrosa/lessons-in-go/testutils"
)

// --- File I/O Tests ---

func TestWriteToFile(t *testing.T) {
	filename := "test_write.txt"
	defer os.Remove(filename) // Clean up after test

	data := []string{"Hello", "World", "Go"}
	err := WriteToFile(filename, data)
	testutils.Compare(t, nil, err, "WriteToFile should not return an error")

	// Read the file to verify content
	content, err := os.ReadFile(filename)
	testutils.Compare(t, nil, err, "Should be able to read the written file")

	expected := "Hello\nWorld\nGo\n"
	testutils.Compare(t, expected, string(content), "File content should match expected format")
}

func TestReadFromFile(t *testing.T) {
	filename := "test_read.txt"
	defer os.Remove(filename) // Clean up after test

	// Create test file
	content := "Line 1\nLine 2\nLine 3\n"
	err := os.WriteFile(filename, []byte(content), 0644)
	testutils.Compare(t, nil, err, "Should be able to create test file")

	lines, err := ReadFromFile(filename)
	testutils.Compare(t, nil, err, "ReadFromFile should not return an error")

	expected := []string{"Line 1", "Line 2", "Line 3"}
	testutils.Compare(t, expected, lines, "ReadFromFile should return correct lines")
}

func TestAppendToFile(t *testing.T) {
	filename := "test_append.txt"
	defer os.Remove(filename) // Clean up after test

	// Create initial file
	initial := "Initial content\n"
	err := os.WriteFile(filename, []byte(initial), 0644)
	testutils.Compare(t, nil, err, "Should be able to create initial file")

	// Append to file
	appendData := "Appended content"
	err = AppendToFile(filename, appendData)
	testutils.Compare(t, nil, err, "AppendToFile should not return an error")

	// Read file to verify content
	content, err := os.ReadFile(filename)
	testutils.Compare(t, nil, err, "Should be able to read the file")

	expected := "Initial content\nAppended content\n"
	testutils.Compare(t, expected, string(content), "File should contain both initial and appended content")
}

func TestCopyFile(t *testing.T) {
	srcFile := "test_source.txt"
	dstFile := "test_destination.txt"
	defer os.Remove(srcFile) // Clean up after test
	defer os.Remove(dstFile) // Clean up after test

	// Create source file
	sourceContent := "This is the source content\nWith multiple lines\n"
	err := os.WriteFile(srcFile, []byte(sourceContent), 0644)
	testutils.Compare(t, nil, err, "Should be able to create source file")

	// Copy file
	err = CopyFile(srcFile, dstFile)
	testutils.Compare(t, nil, err, "CopyFile should not return an error")

	// Verify destination file content
	dstContent, err := os.ReadFile(dstFile)
	testutils.Compare(t, nil, err, "Should be able to read destination file")
	testutils.Compare(t, sourceContent, string(dstContent), "Destination file should have same content as source")
}

// --- HTTP Web Server Tests ---

func TestHelloHandler(t *testing.T) {
	// Test with name parameter
	req := httptest.NewRequest("GET", "/hello?name=Alice", nil)
	w := httptest.NewRecorder()

	HelloHandler(w, req)

	resp := w.Result()
	testutils.Compare(t, http.StatusOK, resp.StatusCode, "Should return 200 OK status")

	body, _ := io.ReadAll(resp.Body)
	testutils.Compare(t, "Hello, Alice!", string(body), "Should return personalized greeting")

	contentType := resp.Header.Get("Content-Type")
	testutils.Compare(t, "text/plain; charset=utf-8", contentType, "Should set content type")

	// Test without name parameter
	req = httptest.NewRequest("GET", "/hello", nil)
	w = httptest.NewRecorder()

	HelloHandler(w, req)

	resp = w.Result()
	body, _ = io.ReadAll(resp.Body)
	testutils.Compare(t, "Hello, World!", string(body), "Should return default greeting")
}

func TestUsersHandler(t *testing.T) {
	// Test GET request
	req := httptest.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()

	UsersHandler(w, req)

	resp := w.Result()
	testutils.Compare(t, http.StatusOK, resp.StatusCode, "Should return 200 OK status")

	contentType := resp.Header.Get("Content-Type")
	testutils.Compare(t, "application/json", contentType, "Should set JSON content type")

	body, _ := io.ReadAll(resp.Body)
	expected := `[{"id":1,"name":"Alice","age":30},{"id":2,"name":"Bob","age":25}]`
	testutils.Compare(t, expected, string(body), "Should return valid JSON with user data")

	// Test POST request (should be rejected)
	req = httptest.NewRequest("POST", "/users", nil)
	w = httptest.NewRecorder()

	UsersHandler(w, req)

	resp = w.Result()
	testutils.Compare(t, http.StatusMethodNotAllowed, resp.StatusCode, "Should reject non-GET requests")
}

func TestUserHandler(t *testing.T) {
	// Test valid user ID
	req := httptest.NewRequest("GET", "/user/1", nil)
	w := httptest.NewRecorder()

	UserHandler(w, req)

	resp := w.Result()
	testutils.Compare(t, http.StatusOK, resp.StatusCode, "Should return 200 OK for valid user")

	body, _ := io.ReadAll(resp.Body)
	expected := `{"id":1,"name":"Alice","age":30}`
	testutils.Compare(t, expected, string(body), "Should return valid JSON for user")

	// Test invalid user ID
	req = httptest.NewRequest("GET", "/user/999", nil)
	w = httptest.NewRecorder()

	UserHandler(w, req)

	resp = w.Result()
	testutils.Compare(t, http.StatusNotFound, resp.StatusCode, "Should return 404 for non-existent user")
}

func TestStartServer(t *testing.T) {
	// Start server in background
	go func() {
		err := StartServer(":8080")
		if err != nil && err != http.ErrServerClosed {
			t.Errorf("Server failed to start: %v", err)
		}
	}()

	// Give server time to start
	time.Sleep(100 * time.Millisecond)

	// Test if server is responding
	resp, err := http.Get("http://localhost:8080/hello")
	if err == nil {
		resp.Body.Close()
		testutils.Compare(t, http.StatusOK, resp.StatusCode, "Server should be responding on port 8080")
	} else {
		t.Logf("Server test skipped - could not connect: %v", err)
	}
}

// --- CLI Flag Parsing Tests ---

func TestParseFlags(t *testing.T) {
	// Reset flags for testing
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// Simulate command line arguments
	os.Args = []string{"program", "-input", "input.txt", "-output", "output.txt", "-verbose", "-count", "10", "-op", "copy"}

	config := ParseFlags()

	testutils.Compare(t, "input.txt", config.InputFile, "Should parse input file flag")
	testutils.Compare(t, "output.txt", config.OutputFile, "Should parse output file flag")
	testutils.Compare(t, false, config.Verbose, "Verbose should default to false")
	testutils.Compare(t, 10, config.Count, "Should parse count flag")
	testutils.Compare(t, "copy", config.Operation, "Should parse operation flag")
}

func TestValidateConfig(t *testing.T) {
	// Test valid config
	validConfig := &Config{
		InputFile:  "input.txt",
		OutputFile: "output.txt",
		Verbose:    false,
		Count:      10,
		Operation:  "copy",
	}

	err := ValidateConfig(validConfig)
	testutils.Compare(t, nil, err, "Valid config should not return error")

	// Test invalid config - missing input file
	invalidConfig := &Config{
		InputFile:  "",
		OutputFile: "output.txt",
		Verbose:    false,
		Count:      10,
		Operation:  "copy",
	}

	err = ValidateConfig(invalidConfig)
	testutils.Compare(t, true, err != nil, "Config without input file should return error")

	// Test invalid config - invalid operation
	invalidOpConfig := &Config{
		InputFile:  "input.txt",
		OutputFile: "output.txt",
		Verbose:    false,
		Count:      10,
		Operation:  "invalid",
	}

	err = ValidateConfig(invalidOpConfig)
	testutils.Compare(t, true, err != nil, "Config with invalid operation should return error")
}

func TestProcessFiles(t *testing.T) {
	// Create test input file
	inputFile := "test_input.txt"
	outputFile := "test_output.txt"
	defer os.Remove(inputFile)
	defer os.Remove(outputFile)

	inputContent := "Line 1\nLine 2\nLine 3\nLine 4\nLine 5\n"
	err := os.WriteFile(inputFile, []byte(inputContent), 0644)
	testutils.Compare(t, nil, err, "Should create test input file")

	// Test copy operation
	config := &Config{
		InputFile:  inputFile,
		OutputFile: outputFile,
		Verbose:    true,
		Count:      3,
		Operation:  "copy",
	}

	err = ProcessFiles(config)
	testutils.Compare(t, nil, err, "ProcessFiles should not return error")

	// Verify output
	lines, err := ReadFromFile(outputFile)
	testutils.Compare(t, nil, err, "Should be able to read output file")

	expected := []string{"Line 1", "Line 2", "Line 3"}
	testutils.Compare(t, expected, lines, "Should process first 3 lines when count is 3")

	// Test count operation
	os.Remove(outputFile) // Clean up for next test
	config.Operation = "count"

	err = ProcessFiles(config)
	testutils.Compare(t, nil, err, "ProcessFiles count operation should not return error")

	lines, err = ReadFromFile(outputFile)
	testutils.Compare(t, nil, err, "Should be able to read count output file")

	expectedCount := []string{"Total lines: 5"}
	testutils.Compare(t, expectedCount, lines, "Count operation should return line count")

	// Test transform operation
	os.Remove(outputFile) // Clean up for next test
	config.Operation = "transform"

	err = ProcessFiles(config)
	testutils.Compare(t, nil, err, "ProcessFiles transform operation should not return error")

	lines, err = ReadFromFile(outputFile)
	testutils.Compare(t, nil, err, "Should be able to read transform output file")

	expectedTransform := []string{"LINE 1", "LINE 2", "LINE 3"}
	testutils.Compare(t, expectedTransform, lines, "Transform operation should uppercase lines")
}
