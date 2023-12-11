package builtins

import (
    "fmt"
    "os"
)

func Touch(args ...string) error {
    if len(args) < 2 {
        return fmt.Errorf("touch: missing file operand")
    }
    file, err := os.Create(args[1])
    if err != nil {
        return fmt.Errorf("touch: %s", err)
    }
    defer file.Close()
    return nil
}
