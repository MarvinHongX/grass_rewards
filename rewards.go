// rewards.go
//
// Author:    Hong
// Created:   6/11/2024
// Modified:  6/11/2024
// Notes:

package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

// readRewards reads a file and returns a map of epoch to accumulated rewards.
func readRewards(filename string) (map[string]float64, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    rewards := make(map[string]float64)
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Fields(line)

        if len(parts) != 2 {
            continue // skip malformed lines
        }

        epoch := parts[0]
        valueStr := parts[1]
        var value float64

        if strings.HasSuffix(valueStr, "K") {
            baseValue, err := strconv.ParseFloat(strings.TrimSuffix(valueStr, "K"), 64)
            if err != nil {
                return nil, err
            }
            value = baseValue * 1000
        } else {
            value, err = strconv.ParseFloat(valueStr, 64)
            if err != nil {
                return nil, err
            }
        }

        rewards[epoch] += value
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return rewards, nil
}

func main() {
    files := []string{"rewards01.txt", "rewards02.txt"}
    totalRewards := make(map[string]float64)

    // Read and accumulate rewards from files
    for _, file := range files {
        rewards, err := readRewards(file)
        if err != nil {
            fmt.Printf("Error reading %s: %v\n", file, err)
            return
        }

        for epoch, reward := range rewards {
            totalRewards[epoch] += reward
        }
    }

    // Get sorted keys
    var epochs []string
    for epoch := range totalRewards {
        epochs = append(epochs, epoch)
    }
    sort.Strings(epochs) // Sort keys

    grandTotal := 0.0
    fmt.Println("Total rewards by epoch:")
    for _, epoch := range epochs {
        total := totalRewards[epoch]
        fmt.Printf("%s: %.2f\n", epoch, total)
        grandTotal += total
    }
    
    fmt.Printf("Grand Total: %.2f\n", grandTotal)
}

