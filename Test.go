package main

import (
	"fmt"
)



func main() {
	blockchain := []Block{}
	blockchain = AddBlock("0", "0", blockchain)
	blockchain = AddBlock(blockchain[0].currHash, "1", blockchain)
	blockchain = AddBlock(blockchain[1].currHash, "2", blockchain)

	//fmt.Println(blockchain)
	blockchain=SetBlock(blockchain,0,"322");
	//fmt.Println(blockchain)
	ReMine(blockchain);
	fmt.Print(ValidateChain(blockchain));
}

