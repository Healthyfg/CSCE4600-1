package builtins

import (
    "fmt"
    "strings"
)

func Echo(args ...string) error {
    if len(args) < 2 {
        return fmt.Errorf("echo: no arguments provided")
    }
    fmt.Println(strings.Join(args[1:], " "))
    return nil
}
