package main

import (
	"flag"
	"fmt"
	"os/exec"
	"runtime"
)

// execute command with path and commands parameters
// path will be will the directory and first to be executed
// commands are set of commands separated by ";"
func execute(path, commands string) {
	// NOTE: the path is first executed that is, cd <selected-path>
	// and the commands are separated by ";"
	out, err := exec.Command("/bin/sh", "-c", fmt.Sprintf("cd %s; %s", path, commands)).Output()

	// if there is an error with our execution
	// handle it here
	if err != nil {
		fmt.Println("Command execution failed. Logs:")
		panic(err)
	}

	// as the out variable defined above is of type []byte we need to convert
	// this to a string or else we will see garbage printed out in our console
	// this is how we convert it to a string
	fmt.Println("Command successfully executed. Logs:")
	output := string(out[:])
	fmt.Println(output)
}

func main() {
	// region flags
	path := flag.String("path", "/var/www", "path of project commands will be executed")
	commands := flag.String("commands", "whoami; uptime; pwd;", "one line set of commands")

	flag.Parse()
	// endregion flags

	// TODO: support for windows
	if runtime.GOOS == "windows" {
		fmt.Println("Can't execute this on a windows machine")
	} else {
		execute(*path, *commands)
	}
}
