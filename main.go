package main

import ("bufio"; "fmt"; "os")

// go run main.go
// go build -o hello.exe
func main() {
  var name string
  fmt.Printf("Enter name: %v", name)

  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  name = scanner.Text()

  fmt.Printf("hello %v\n", name)
  
  scanner.Scan() // wait to see output
}
