package builtins

import (
    "fmt"
    "os"
)

func Mkdir(args ...string) error {
    if len(args) < 2 {
        return fmt.Errorf("mkdir: missing operand")
    }
    err := os.Mkdir(args[1], 0755)
    if err != nil {
        return fmt.Errorf("mkdir: %s", err)
    }
    return nil
}
