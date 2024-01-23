package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter project name: ")
	projectName, _ := reader.ReadString('\n')
	projectName = projectName[:len(projectName)-1]

	baseDirs := []string{"app", "ui", "handlers", "cmd/server", "internal/auth", "data", "public", "configs"}
	for _, dir := range baseDirs {
		dirPath := filepath.Join(projectName, dir)
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			fmt.Println("Failed to create directory:", err)
		}
	}

	files := map[string][]string{
		"app":           {"app.go", "app.layout.go"},
		"ui":            {"sign-in.go", "sign-out.go"},
		"handlers":      {"app.handler.go", "auth.handler.go"},
		"cmd/server":    {"main.go"},
		"internal/auth": {"auth.service.go"},
		"data":          {"db.go", "user.model.go"},
	}

	for dir, fileList := range files {
		for _, file := range fileList {
			filePath := filepath.Join(projectName, dir, file)
			err := os.WriteFile(filePath, []byte{}, 0644)
			if err != nil {
				fmt.Println("Failed to create file:", err)
				return
			}
		}
	}

	fileCreationData := []struct {
		Path, Content string
	}{
		{filepath.Join(projectName, "go.mod"), "module " + projectName},
		{filepath.Join(projectName, "go.sum"), ""},
		{filepath.Join(projectName, "README.md"), "# " + projectName},
		{filepath.Join(projectName, ".gitignore"), ".DS_Store\n*.swp\nvendor/"},
	}

	for _, data := range fileCreationData {
		if err := os.WriteFile(data.Path, []byte(data.Content), 0644); err != nil {
			fmt.Println("Failed to create file:", data.Path, "Error:", err)
			return
		}
	}
}
