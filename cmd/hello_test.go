package main

import (
    "testing"
    "."
)

func TestGreeting(t *testing.T) {
    if japanese() != "こんにちは、世界" {
        t.Errorf("Excepted %s, But %s.", "こんにちは、世界", japanese())
    }
    if chinese() != "你好,是世界" {
        t.Errorf("Excepted %s, But %s.", "你好,是世界", chinese())
    }
    if english() != "Hello, world" {
        t.Errorf("Excepted %s, But %s.", "Hello, world", english())
    }
}
