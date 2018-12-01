package hello

func English() string {
    return "Hello, world"
}

func Japanese() string {
//    return "Hello, world"      /* <- ここがバグ */
    return "こんにちは、世界"
}

func Chinese() string {
    return "你好,是世界"
}
