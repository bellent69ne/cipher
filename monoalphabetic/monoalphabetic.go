package monoalphabetic

import (
    "math/rand"
    "time"
    "fmt"
    "errors"
)

func keyIsValid(key []rune) bool {
    for _, rn := range key {
        if rn >= 'A' && rn <= 'Z' ||
            rn >= 'a' && rn <= 'z' {
                continue
            } else {
                return false
            }
    }
    return true
}

func makeRotor(key []rune) (map[rune]rune, error) {
    rotor := make(map[rune]rune)

    if keyIsValid(key) {
        i := 0

        for upper, lower := 'A', 'a'; upper <= 'Z';
            upper, lower = upper + 1, lower + 1 {
            rotor[upper] = key[i]
            rotor[lower] = key[i] + 32
            i++
        }
    } else {
        return nil, errors.New("invalid key... Contains runes other than english letters")
    }

    return rotor, nil
}

func Encrypt(plain string, key []rune) string {

    rotor, err := makeRotor(key)
    if err != nil {
        fmt.Printf("Error occured: %v\n", err)
    }

    cipher := []rune(plain)

    for i, rn := range cipher {
        elem, ok := rotor[rn]
        if ok {
            cipher[i] = elem
        }
    }

    return string(cipher)
}

func reverseRotor(key []rune) (map[rune]rune, error) {
    rotor := make(map[rune]rune)

    if keyIsValid(key) {
        i := 0

        for upper, lower := 'A', 'a'; upper <= 'Z';
            upper, lower = upper + 1, lower + 1 {
            rotor[key[i]] = upper
            rotor[key[i] + 32] = lower
            i++
        }
    } else {
        return nil, errors.New("invalid key... Contains runes other than english letters")
    }

    return rotor, nil
}

func Decrypt(cipher string, key []rune) string {
    rotor, err := reverseRotor(key)
    if err != nil {
        fmt.Printf("Error occured: %v\n", err)
    }

    plain:= []rune(cipher)

    for i, rn := range plain {
        elem, ok := rotor[rn]
        if ok {
            plain[i] = elem
        }
    }

    return string(plain)
}

func KeyGen() []rune {
    seed := rand.NewSource(time.Now().UnixNano())
    random := rand.New(seed)

    chars := make([]rune, 26)

    rn := 'A'
    for i := range chars {
        chars[i] = rn
        rn++
    }

    key := make([]rune, 26)
    for i := range key {
        index := random.Intn(len(chars))
        key[i] = chars[index]

        chars = remove(chars, index)
    }

    return key
}

func remove(slice []rune, i int) []rune {
    slice[i] = slice[len(slice)-1]
    return slice[:len(slice)-1]
}
