package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func main() {
    file, err := os.Open("./data.txt")

    if err != nil {
        fmt.Println(err);
        return
    }

    defer file.Close();

    scanner := bufio.NewScanner(file);
    sum := 0;

    for scanner.Scan() {
        line := scanner.Text();

        gameData := strings.Split(line, ":")
        gameMeta, roundsString := gameData[0], gameData[1]
        gameNumber, _ := strconv.Atoi(strings.Split(strings.TrimSpace(gameMeta), " ")[1])

        fmt.Println(gameNumber)

        rounds := strings.Split(roundsString, ";")
        maxMap := map[string] int {}

        for _, round := range rounds {
            colors := strings.Split(round, ",")
            for _, color := range colors {
                parts := strings.Split(strings.TrimSpace(color), " ")
                number, _ := strconv.Atoi(parts[0]);
                color := parts[1]
                maxMap[color] = max(number, maxMap[color])
            }
        }

        power := 1
        for _, min_ := range maxMap {
            power *= min_;
        }
        sum += power;

    }

    fmt.Println(sum)
}
