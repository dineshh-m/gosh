package shell

import (
	"bufio"
	"errors"
	"fmt"
	"gosh/term"
	"os"
	"os/exec"
	"strings"
)

var cwd string

func executeCommand(input string) error {
  input = strings.TrimSuffix(input, "\n")
  t := strings.Split(input, " ")
  command, args := t[0], t[1:]
  if command == "" {
    return nil
  }

  switch command {
  case "cd":
    if len(args) < 1 {
      return errors.New("path required")
    } 
    return os.Chdir(args[0])
  case "exit":
    os.Exit(0)
  } 

  cmd := exec.Command(command, args...)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  
  return cmd.Run()
}

func Gosh() {
  reader := bufio.NewReader(os.Stdin)
  for {
    dir, err := os.Getwd()
    if err != nil {
      fmt.Fprintln(os.Stderr, err)
      os.Exit(1)
    }
    splitDir := strings.Split(dir, "/")
    cwd = splitDir[len(splitDir)-1]
    fmt.Print(term.Term(cwd))
    input, err := reader.ReadString('\n') 
    if err != nil {
      fmt.Fprintf(os.Stderr, "Error occured while reading command")
      os.Exit(1)
    }

    if err := executeCommand(input); err != nil {
      fmt.Fprintln(os.Stderr, err)
    }
  }
}
