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

// GetInput gets the input data and returns it as a string.
func GetInput() string {
  dat, err := ioutil.ReadFile("input")
  if err != nil {
    log.Fatal(err)
  }
  return string(dat)
}

// getNext returns the index for the value we need to compare with the
// current value.  Simple imcrementing wont work since we will get
// array index out of bounds errors.  getNext handles the index for the next
// value when skip is 1 and we are already at the end of the array,
// as well as the case where skip is half the size of the array.
func GetNext(myincr int,  mycurr int,  mysize int) int {
  if (mycurr + myincr) > mysize {
    return ((mycurr + myincr) - mysize) -1
  }
  return mycurr + myincr
}

// getCatpcha calculates the captcha for either part of day1
func GetCaptcha(data []string, partNo int) int {
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
    next, err = strconv.Atoi(data[GetNext(skip, i, inputLen)])

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
  fmt.Printf("CAPTCHA: %d\n", GetCaptcha(inputArray, 1))
  fmt.Printf("CAPTCHA: %d\n", GetCaptcha(inputArray, 2))
}
