package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type intPair struct {
    first  int
    second int
}

func isByteDigit(b byte) bool {
    return b >= '0' && b <= '9'
}

func genContour(start int, finish int) <- chan intPair {
    ch := make(chan intPair)

    go func() {
        for i := start - 1; i <= finish; i++ {
            ch <- intPair{0, i}
            ch <- intPair{2, i}
        }
        ch <- intPair{1, start - 1}
        ch <- intPair{1, finish}

        close(ch)
    }()

    return ch
}

func main() {
    file, err := os.Open("./data2.txt")

    if err != nil {
        fmt.Println(err);
        return
    }

    defer file.Close();

    scanner := bufio.NewScanner(file);

    scanner.Scan()


    sum := 0
    var lines [3] string;

    lines[2] = scanner.Text();

    fake := strings.Repeat(".", len(lines[2]));
    lines[1] = fake;


    for scanner.Scan() {
        lines[0] = lines[1];
        lines[1] = lines[2];

        lines[2] = scanner.Text();

        i := 0;
        var start int;


        fmt.Println("new")
        fmt.Println(lines[0])
        fmt.Println(lines[1])
        fmt.Println(lines[2])

        for i < len(lines[1]) {
            if isByteDigit(lines[1][i]) {
                start = i;
                for i < len(lines[1]) && isByteDigit(lines[1][i]) {
                    i++
                }

                is_part := true;
                for pair := range genContour(start, i) {
                    if lines[pair.first][pair.second] != '.' {
                        is_part = false;
                        fmt.Println(lines[pair.first][pair.second] )
                        break;
                    }
                }

                if is_part {
                    rez, _ := strconv.Atoi(lines[1][start:i])
                    fmt.Println(strconv.Atoi(lines[1][start:i]))
                    fmt.Println(sum)
                    sum += rez
                }

                continue
            }
            i++
        }

    }

    fmt.Println(sum)

}
