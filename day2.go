package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// https://adventofcode.com/2023/day/2
//--- Day 2: Cube Conundrum ---
func Day2(){
    redCount:= 12
    blueCount:=14
    greenCount:=13
    
    gameIdSum := 0
      

    fp, err := os.Open("day2Input.txt")
    if err != nil{
        panic(err.Error())
    }

    fileScanner := bufio.NewScanner(fp)
    for fileScanner.Scan(){
        parts := strings.Split(fileScanner.Text(), ": ")
        gameId, err:= strconv.Atoi(strings.Split(parts[0],"Game ")[1])

        if err != nil{
            panic(err.Error())
        }
          

        sanitized := parts[1]
        start := -1
        end := -1
        i := 0 
        reds:=0
        blues:=0
        greens:=0

        for i < len(sanitized){
            char := sanitized[i]
            if string(char) == ";"{
                reds=0
                blues=0
                greens=0
                start = -1
                end = -1
            }

            if unicode.IsDigit(rune(char)){
                if start == -1{
                    start = i
                }

            }else if start != -1{
                end = i
            }

            if start != -1 && end != -1{
                digits,err:= strconv.Atoi(sanitized[start:end])
                if err != nil{
                    panic(err.Error())
                }

                if sanitized[end+1:end+4] == "red"{
                    reds+=digits
                    if reds>redCount{
                        gameIdSum-=gameId
                        break
                    }

                } else if sanitized[end+1:end+5] == "blue"{
                    blues+=digits
                    if blues>blueCount{
                        gameIdSum-=gameId
                        break
                    }
                } else if sanitized[end+1:end+6] == "green"{
                    greens+=digits
                    if greens>greenCount{
                        gameIdSum-=gameId
                        break
                    }
                }

                start = -1
                end = -1
            }

            i+=1
        }

    }

    fmt.Println("game sum",gameIdSum)
}

func Day2P2(){
    
    powersSum := 0
      

    fp, err := os.Open("day2Input.txt")
    if err != nil{
        panic(err.Error())
    }

    fileScanner := bufio.NewScanner(fp)
    for fileScanner.Scan(){
        parts := strings.Split(fileScanner.Text(), ": ")

        if err != nil{
            panic(err.Error())
        }
          

        sanitized := parts[1]
        start := -1
        end := -1
        i := 0 
        reds:=0
        blues:=0
        greens:=0

        for i < len(sanitized){
            char := sanitized[i]
            if string(char) == ";"{
                start = -1
                end = -1
            }

            if unicode.IsDigit(rune(char)){
                if start == -1{
                    start = i
                }

            }else if start != -1{
                end = i
            }

            if start != -1 && end != -1{
                digits,err:= strconv.Atoi(sanitized[start:end])
                if err != nil{
                    panic(err.Error())
                }

                if sanitized[end+1:end+4] == "red"{
                    reds = int(math.Max(float64(reds),float64(digits)))
                } else if sanitized[end+1:end+5] == "blue"{
                    blues = int(math.Max(float64(blues),float64(digits)))
                } else if sanitized[end+1:end+6] == "green"{
                    greens = int(math.Max(float64(greens),float64(digits)))
                }

                start = -1
                end = -1
            }

            i+=1
        }
        powersSum += greens*reds*blues

    }

    fmt.Println("powers sum",powersSum)
}

