package main

import (
	"bufio"
	"fmt"
	"os"
    "regexp"
)

func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}


func main() {
    file, err := os.Open("./data.txt")

    start_re, _ := regexp.Compile("one|two|three|four|five|six|seven|eight|nine|[0-9]")
    end_re, _ := regexp.Compile("enin|thgie|neves|xis|evif|ruof|eerht|owt|eno|[0-9]")

    digitMap := map[string]int{
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
        "eno": 1,
        "owt": 2,
        "eerht": 3,
        "ruof": 4,
        "evif": 5,
        "xis": 6,
        "neves": 7,
        "thgie": 8,
        "enin": 9,
    }

    if err != nil {
        fmt.Println(err)
        return
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)

    var firstDigit int;
    var secondDigit int;
    sum := 0;

    for scanner.Scan() {
        code := scanner.Text();

        firstMatch := start_re.FindString(code)
        lastMatch := end_re.FindString(reverseString(code))

        if len(firstMatch) == 1 {
            firstDigit = int(firstMatch[0] - '0');
        } else {
            firstDigit = digitMap[firstMatch];
        }

        if len(lastMatch) == 1 {
            secondDigit = int(lastMatch[0] - '0');
        } else {
            secondDigit = digitMap[lastMatch];
        }

        fmt.Println(code);
        fmt.Println(firstDigit)
        fmt.Println(secondDigit)
        fmt.Println(firstDigit * 10 + secondDigit)
        sum += firstDigit * 10 + secondDigit;
    }

    fmt.Println(sum)

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }

}
