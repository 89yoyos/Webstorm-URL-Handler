package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

var execLoc = ""

// Redux Devtools URL scheme:
// [method]://[hostname]/?url=file://[path]&line=[line]&column=[column]
// http://open/?url=file://C:/This/is/an/example.tsx&line=33&column=27
// Using this tool, port=80 and host=open
func main() {
	var executable, host, port string
	flag.StringVar(&port, "port", "80", "port to listen on")
	flag.StringVar(&host, "host", "open", "host name in url")
	flag.StringVar(&executable, "executable", "webstorm.cmd", "Webstorm executable")
	flag.Parse()
	fmt.Println("\n\n\nRaw Arguments:")
	fmt.Println(os.Args)
	fmt.Println("\nParsed Arguments:")
	fmt.Println("host: " + host)
	fmt.Println("port: " + port)
	fmt.Println("executable: " + executable)
	executableDirectory, executableFile := getDirectoryAndExecutable(executable)
	fmt.Println("Executable File: " + executableFile)
	fmt.Println("Executable Directory" + executableDirectory)

	execLoc = executable

	http.HandleFunc("/", openWebstorm)

	err := http.ListenAndServe(host+":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
func openWebstorm(w http.ResponseWriter, req *http.Request) {
	directory, _ := getDirectoryAndExecutable(execLoc)
	file := req.URL.Query().Get("url")
	if file != "" {
		file = string([]rune(file)[7:])
		line := req.URL.Query().Get("line")
		// Redux Devtools supports specifying the column, but almost nothing uses that
		// column := req.URL.Query().Get("column")
		cmd := exec.Command(execLoc)
		cmd.Dir = directory
		cmd.Args = append(cmd.Args, "--line", line, file)
		fmt.Println(cmd)
		var out bytes.Buffer
		var stderr bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &stderr
		err := cmd.Start()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		} else {
			fmt.Println("Result: " + out.String())
		}
		_, _ = fmt.Fprintf(w, "<script>window.close()</script>\n")
	}
}

func getDirectoryAndExecutable(path string) (string, string) {
	fullPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println("Cannot find file: \"" + path + "\"\n")
		log.Fatal(err)
	}
	file, err := os.Stat(fullPath)
	if err != nil {
		fmt.Println("Cannot find file: \"" + path + "\"\n")
		log.Fatal(err)
	}
	if file.IsDir() {
		log.Fatal("Argument must be an executable file: \"" + path + "\"")
	}
	executable := filepath.Base(fullPath)
	directory := filepath.Dir(fullPath)
	return directory, executable
}
