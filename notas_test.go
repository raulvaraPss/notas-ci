package main

import "testing"

func TestNota(t *testing.T) {
    if nota() != "nota lista" {
        t.Fatal(nota())
    }
}
