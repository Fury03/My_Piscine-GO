package student

func ToLower(s string) string {
    runes := []rune(s)
    for i, ch := range runes {
        if ch >= 'A' && ch <= 'Z' {
            runes[i] = ch + ('a' - 'A') // convert to lowercase could have used ch+32 but to make the code easier to read
        }
    }
    return string(runes)
}
