package monoalphabetic

import (
    "math/rand"
    "time"
    "fmt"
    "errors"
)

func lowerCaseKey(key []rune) []rune {
    lowerCase := make([]rune, len(key))

    for i := range lowercase {
        lowercase[i] = key[i] + rune(32)
    }

    return lowerCase
}

func makeRotor(key []rune) (map[rune]rune, error) {
    var begin, end rune

    fillRotor := func() map[rune]rune {
        rot := make(map[rune]rune)
        i := 0
        for rn := begin; rn <= end; rn++ {
            rot[rn] = key[i]
            i++
        }
        return rot
    }
    var rotor map[rune]rune
    if key[0] >= 'A' && key[0] <= 'Z' {
        //i := 0
        //for rn := 'A'; rn <= 'Z'; rn++ {
        //    rotor[rn] = key[i]
        //    i++
        //}
        begin, end = 'A', 'B'
        rotor = fillRotor()
    } else if key[0] >= 'a' && key[0] <= 'z' {
    //    for rn := 'a'; rn < 'z'; rn++ {
    //        rotor[rn] = key[i]
    //        i++
    //    }
        begin, end = 'a', 'b'
        rotor = fillRotor()
    } else {
        return (nil, errors.New("invalid key...Contains runes other than english letters"))
    }

    return (rotor, nil)
}

func Encrypt(plain, key []rune) string {
    //rotor := make(map[rune]rune)
    //i := 0
    //for rn := 'A'; rn <= 'Z'; rn++ {
    //    rotor[rn] = key[i]
    //    i++
    //}

    //for rn := 'a'; rn <= 'z'; rn++ {
    //    rotor[rn] = key[i]
    //}
    rotor, err := makeRotor(key)
    if err != nil {
        fmt.Printf("Error occured: %v\n", err)
    }

    fmt.Println(rotor)

    return ""
}
//
//func Decrypt(cipher, key *string) string {
//
//}

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
    fmt.Println()
    fmt.Println(key)
    fmt.Println(string(key))

    return ""
}

func remove(slice []rune, i int) []rune {
    slice[i] = slice[len(slice)-1]
    return slice[:len(slice)-1]
}
