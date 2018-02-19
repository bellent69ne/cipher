package caesar

func Encrypt(plain *string, key int8) string {
    cipher := []rune(*plain)

    for i, rn := range cipher {
        switch {
        case rn >= 'A' && rn <= 'Z':
            cipher[i] = 'A' + (rn - 'A' + rune(key)) % 26
        case rn >= 'a' && rn <= 'z':
            cipher[i] = 'a' + (rn - 'a' + rune(key)) % 26
        }
    }
    return string(cipher)
}

func Decrypt(cipher *string, key int8) string {
    plain := []rune(*cipher)

    for i, rn := range plain {
        switch {
        case rn >= 'A' && rn <= 'Z':
            plain[i] = 'Z' - ('Z' - rn + rune(key)) % 26
        case rn >= 'a' && rn <= 'z':
            plain[i] = 'z' - ('z' - rn + rune(key)) % 26
        }
    }

    return string(plain)
}
