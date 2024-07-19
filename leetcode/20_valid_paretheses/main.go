package main

func isValid(s string) bool {
    if len(s) == 0 || len(s) % 2 == 1 {
        return false
    }
    brackets := map[rune]rune {
        '(': ')',
        '[': ']',
        '{': '}',
    }
    stack := []rune{}
    // [][](){]()
    for _, r := range s{
        if _, ok := brackets[r]; ok{
            stack = append(stack, r)
        } else if len(stack) == 0 || brackets[stack[len(stack)-1]] != r{
            return false
        } else {
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}