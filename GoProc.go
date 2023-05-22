package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`Please specify the folder and file name as a command line argument, e.g., ./goprocgo.exe 'Child/GoChild.go'`)
		return
	}

	path := os.Args[1]

	fmt.Println("Creating folder...")
	CreateFolder(path)

	fmt.Printf("Executing Program...\n\n")
	cmd := `GoChild.go`
	ExecuteProgram(cmd)
}

func CreateFolder(path string) {
	if !strings.HasSuffix(path, ".go") {
		fmt.Println("Invalid file extension. Only .go files are supported.")
		return
	}

	dir, file := filepath.Split(path)
	dir = strings.TrimRight(dir, "/")
	if dir == "" {
		fmt.Println("Invalid folder name.")
		return
	}

	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		// Folder already exists, delete it
		err := os.RemoveAll(dir)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	err := os.Mkdir(dir, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	filepath := filepath.Join(dir, file)
	f, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	fmt.Fprintln(f, "package main")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "import \"fmt\"")
	fmt.Fprintln(f)
	fmt.Fprintln(f, "func main() {")
	fmt.Fprintln(f, "    fmt.Println(\"Hallo aus Berlin!\")")
	fmt.Fprintln(f, "}")
}

func ExecuteProgram(cmd string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	cmdPath := filepath.Join(dir, "Child", cmd)
	goCmd := exec.Command("go", "run", cmdPath)
	goCmd.Dir = filepath.Join(dir, "Child")

	// Output will run the process
	// terminates and returns the standard
	// output in a byte slice.
	buff, err := goCmd.Output()

	if err != nil {
		panic(err)
	}

	// The output of the child
	// process in the form of a byte slice
	// printed as a string
	fmt.Println(string(buff))
	fmt.Printf("\nExecution completed Successfully!\n")

}
