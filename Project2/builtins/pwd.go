package builtins

import (
    "fmt"
    "io/ioutil"
)

func List(args ...string) error {
    // If no arguments are passed, list all files in the current directory
    if len(args) == 0 {
        args = append(args, ".")
    }

    // Loop through each specified directory and list files
    for _, dir := range args {
        files, err := ioutil.ReadDir(dir)
        if err != nil {
            return fmt.Errorf("ERROR: can't list files in directory '%s': %v", dir, err)
        }

        for _, file := range files {
            fmt.Println(file.Name())
        }
    }

    return nil
}