func reverse(x string) string {
    p := ""
    for i := len(x) - 1; i >= 0; i-- {
        p += string(x[i]) //convert char to string
    }
    return p
}

func isPalindrome(x int) bool {
    s := strconv.Itoa(x) // convert int to string

    if reverse(s) == s {
        return true
    } else {
        return false
    }
}
