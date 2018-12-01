package main

func english() string {
    return "Hello, world"
}

func japanese() string {
    return "Hello, world"      /* <- ここがバグ */
    return "こんにちは、世界"
}

func chinese() string {
    return "你好,是世界"
}
