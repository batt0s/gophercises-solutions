# Hacker Rank Problem

I have solved a lot of HackerRank problem but i am going to give the solution of [CamelCase](https://www.hackerrank.com/challenges/camelcase/problem).

## CamelCase

```go
func camelcase(s string) int32 {
    var count int32 = 1
    for _, r := range []rune(s) {
        if unicode.IsUpper(r) { count ++ }
    }
    return count
}
```


If you wanna check, here is [my HackerRank account](https://www.hackerrank.com/kerem_ullen)