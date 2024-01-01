package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    file, err := os.Open("./data2.txt")

    if err != nil {
        fmt.Println(err);
        return
    }

    defer file.Close();

    scanner := bufio.NewScanner(file);
    rulesMap := map[string] int {
        "red": 12,
        "green": 13,
        "blue": 14,
    }
    sum := 0;

    for scanner.Scan() {
        line := scanner.Text();

        gameData := strings.Split(line, ":")
        gameMeta, roundsString := gameData[0], gameData[1]
        gameNumber, _ := strconv.Atoi(strings.Split(strings.TrimSpace(gameMeta), " ")[1])

        fmt.Println(gameNumber)

        is_valid := true;
        rounds := strings.Split(roundsString, ";")

        for _, round := range rounds {
            colors := strings.Split(round, ",")
            for _, color := range colors {
                parts := strings.Split(strings.TrimSpace(color), " ")
                number, _ := strconv.Atoi(parts[0]);
                color := parts[1]

                if rulesMap[color] < number {
                    is_valid = false;
                    break
                }
            }
            if ! is_valid {
                break
            }
        }


        if is_valid {
            sum += gameNumber;
        }

    }

    fmt.Println(sum)
}
