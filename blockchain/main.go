package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
	"time"
)

const Difficulty = 16

type Block struct {
	Hash        []byte
	Transaction []byte
	PrevHash    []byte
	TimeStamp   []byte
	Nonce       int
	Difficulty  int
}

type Blockchain struct {
	blockchain []*Block
}

func CreateBlock(transaction string, prevHash []byte) *Block {
	timeStamp := time.Now().String()
	nonce, hash := ProofOfWork(transaction, prevHash, timeStamp)

	return &Block{
		Hash:        hash[:],
		Transaction: []byte(transaction),
		PrevHash:    prevHash,
		TimeStamp:   []byte(timeStamp),
		Nonce:       nonce,
		Difficulty:  Difficulty,
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
		fmt.Printf("prev hash   : %x\n", block.PrevHash)
		fmt.Printf("transaction : %s\n", block.Transaction)
		fmt.Printf("timestamp   : %s\n", block.TimeStamp)
		fmt.Printf("hash        : %x\n", block.Hash)
		fmt.Printf("verify      : %t\n", block.Verify())
		fmt.Printf("-------------------------------\n")
	}
}

func ProofOfWork(transaction string, prevHash []byte, timeStamp string) (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	nonce := 0
	for nonce < math.MaxInt64 {
		data := bytes.Join([][]byte{
			[]byte(transaction),
			[]byte(timeStamp),
			prevHash,
			ToHex(int64(nonce)),
		}, []byte{})

		hash = sha256.Sum256(data)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(target) == -1 {
			break
		} else {
			nonce++
		}
	}
	return nonce, hash[:]
}

func (b *Block) Verify() bool {
	var intHash big.Int
	var hash [32]byte

	target := big.NewInt(1)
	target.Lsh(target, uint(256-b.Difficulty))

	data := bytes.Join([][]byte{
		[]byte(b.Transaction),
		[]byte(b.TimeStamp),
		b.PrevHash,
		ToHex(int64(b.Nonce)),
	}, []byte{})

	hash = sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(target) == -1
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
