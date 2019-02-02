
package main

import (
	"chain/core"
)

func main()  {
	bc := core.NewBlock()
	bc.SendData("my name is lizan222")
	bc.SendData("my name is lizanaa")
	bc.Print()
}