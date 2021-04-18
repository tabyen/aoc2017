package main



import (
    //"bufio"
    "fmt"
    //"io"
    "io/ioutil"
    "log"
    //"os"
    "strconv"
    "strings"
)

func GetInput() string {
  dat, err := ioutil.ReadFile("input")
  if err != nil {
    log.Fatal(err)
  }
  return string(dat)
}

func getNext(myincr int,  mycurr int,  mysize int) int {
  if (mycurr + myincr) > mysize {
    return ((mycurr + myincr) - mysize) -1
  }
  return mycurr + myincr
}

func getCaptcha(data []string, partNo int) int {
  skip := 1
  fmt.Sprint(skip)
  if (partNo == 2) {
    skip = (len(data)/2)
  }
  total := 0
  var next int
  inputLen := len(data) - 1
  for i := 0; i <= inputLen; i++ {
    curr, err := strconv.Atoi(data[i])
    if err != nil {
      log.Fatal(err)
    }
    next, err = strconv.Atoi(data[getNext(skip, i, inputLen)])

    if (curr == next) {
      total += curr
    }
  }
  return total

}

func main() {
  dat := GetInput()
  inputArray := strings.Split(strings.Trim(dat, "\n"), "")

  fmt.Print("Part 1:\n")
  fmt.Printf("CAPTCHA: %d\n", getCaptcha(inputArray, 1))
  fmt.Printf("CAPTCHA: %d\n", getCaptcha(inputArray, 2))
}
