package random

import (
    "math/rand"
    "time"
)

func NewRandomString(size int) string {
    rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

    chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

    result := make([]rune, size)
    for i := 0; i < size; i++ {
        result[i] = chars[rnd.Intn(len(chars))]
    }

    return string(result)
}
