package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Msg struct {
	line string
	err  error
}

// Read from the standard input.
func read() Msg {
	msg := Msg{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("> ")
	line, err := reader.ReadString('\n')

	msg.err = err
	msg.line = line

	return msg
}

func askProjectName() Msg {
	fmt.Println("Hello astronaut ! How shall we name your project today ?")
	msg := read()

	for msg.err != nil || !isProjectNameValid(msg.line) {
		fmt.Println("Please chose a valid name.")
		fmt.Print("\n")
		// TODO: add command to display help message for valid project name.
		msg = read()
	}

	fmt.Printf("You entered: %s\n", msg.line)

	return Msg{line: msg.line, err: nil}
}

// Check if the project name is valid.
func isProjectNameValid(name string) bool {
	// The use entered an empty name if we only got back the return character
	// as a string.
	if name == "\n" {
		return false
	}
	return true
}

func createProjectDirectory(name string) Msg {
	fmt.Println("Creating a folder for your project...")
	err := os.Mkdir(strings.TrimSpace(name), 0755)
	if err != nil {
		fmt.Println("Uh-oh, a problem occured while creating the folder.")
		log.Fatal(err)
	}

	fmt.Println("Folder created successfully !")

	return Msg{line: strings.TrimSpace(name), err: nil}
}

// Create a 'main.go' and a '.gitignore' files.
func createProjectFiles(name string) Msg {
	err := createMainFile(name)
	if err != nil {
		fmt.Println("'main.go' file creation skipped.")
	}
	err = createGitIgnoreFile(name)
	if err != nil {
		fmt.Println("'.gitignore' file creation skipped.")
	}

	return Msg{line: name, err: nil}
}

func createMainFile(name string) error {
	mainFile, err := os.Create(fmt.Sprintf("%s/main.go", name))
	if err != nil {
		fmt.Println("Could not create 'main.go' file.")
		log.Println(err)
		return err
	} else {
		mainFile.WriteString("package main\n")
		mainFile.WriteString("\n")
		mainFile.WriteString("import \"fmt\"\n")
		mainFile.WriteString("func main() {\n")
		mainFile.WriteString("\tfmt.Println(\"Hello, World!\")\n")
		mainFile.WriteString("}")
	}

	return nil
}

func createGitIgnoreFile(name string) error {
	_, err := os.Create(fmt.Sprintf("%s/.gitignore", name))
	if err != nil {
		fmt.Println("Could not create '.gitignore' file.")
		log.Println(err)
	}

	return nil
}

func main() {
	msg := askProjectName()
	msg = createProjectDirectory(msg.line)
	createProjectFiles(msg.line)
}
