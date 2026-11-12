package main

import (
    "fmt"

    "github.com/google/uuid"
)

func nota() string {
    return "nota lista"
}

func main() {
    fmt.Println(nota())
    fmt.Println(uuid.NewString())
}
