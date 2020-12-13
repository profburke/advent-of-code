package main

import (
       "bufio"
       "fmt"
       "os"
       "strconv"
)

func main() {
     scanner := bufio.NewScanner(os.Stdin)
     freq := 0

     for scanner.Scan() {
       line := scanner.Text()
       i, err := strconv.Atoi(line)

       if err != nil {
          fmt.Fprintln(os.Stderr, "converting string to int: ", err)
          os.Exit(1)
       }

       freq += i
     }

     if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input: ", err)
     }

     fmt.Println("result = ", freq)
}
