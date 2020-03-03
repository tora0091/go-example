package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	Hash        []byte
	Transaction []byte
	PrevHash    []byte
	TimeStamp   []byte
}

type Blockchain struct {
	blockchain []*Block
}

func CreateBlock(transaction string, prevHash []byte) *Block {
	timeStamp := time.Now().String()
	info := bytes.Join([][]byte{[]byte(transaction), []byte(timeStamp), prevHash}, []byte{})
	hash := sha256.Sum256(info)
	return &Block{
		Hash:        hash[:],
		Transaction: []byte(transaction),
		PrevHash:    prevHash,
		TimeStamp:   []byte(timeStamp),
	}
}

func (chain *Blockchain) AddBlock(transaction string) {
	prevHash := chain.blockchain[len(chain.blockchain)-1].Hash
	block := CreateBlock(transaction, prevHash)
	chain.blockchain = append(chain.blockchain, block)
}

func InitBlockChain() *Blockchain {
	block := CreateBlock("Genesis", []byte{})
	return &Blockchain{
		blockchain: []*Block{block},
	}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("first of all")
	chain.AddBlock("100 cent from Mike to Array")
	chain.AddBlock("150 cent from Teddy to Nick")
	chain.AddBlock("50 cent from Allen to Fred")

	for _, block := range chain.blockchain {
		fmt.Printf("%x\n", block.PrevHash)
		fmt.Printf("%s\n", block.Transaction)
		fmt.Printf("%s\n", block.TimeStamp)
		fmt.Printf("%x\n", block.Hash)
		fmt.Printf("-------------------------------\n")
	}
}
