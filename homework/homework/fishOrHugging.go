package main
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Date(2019,1, 1, 0, 0, 0, 0,time.Local).Unix())
}