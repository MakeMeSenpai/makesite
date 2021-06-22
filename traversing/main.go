package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    directory := "/Users/aspin25/" // install.sh"
    files, err := ioutil.ReadDir(directory)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        fmt.Println(file.Name())
		fmt.Println(file.IsDir())
    }
}
