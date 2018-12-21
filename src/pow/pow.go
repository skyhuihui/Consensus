package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//0000 0000 0000 0000 1001 0001 0000 .... 0001

// 256位Hash里面前面至少要有16个零
const targetBit = 16

// hash nil
//256位
// 32
// 8
// 二进制表示 0000 0000 0000 0000 0000 0000 0000 0000
// 2 的 32 - 8 次方

type ProofOfWork struct {
	Block *Block // 当前要验证的区块
	// 0000 0001 0000 0000 0000 0000 0000 0000
	target *big.Int // 大数据存储 2^24
}

// 前面必须有8个零
// 0000 0000 1111 1111 1111 1111 1111 1111

// 0000 0001
// 0001 0000
// 0000 1111

// 0001 0000
// 0000 1111

// 0000 0000 0000 0000 1111 111

// 数据拼接，返回字节数组
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBit)),
			IntToHex(int64(nonce)),
			IntToHex(int64(pow.Block.Height)),
		},
		[]byte{},
	)

	return data
}

func (proofOfWork *ProofOfWork) IsValid() bool {

	//1.proofOfWork.Block.Hash
	//2.proofOfWork.Target

	var hashInt big.Int
	// []byte 转 Int
	hashInt.SetBytes(proofOfWork.Block.Hash)

	// Cmp compares x and y and returns:
	//
	//   -1 if x <  y
	//    0 if x == y
	//   +1 if x >  y
	if proofOfWork.target.Cmp(&hashInt) == 1 {
		return true
	}

	return false
}

func (proofOfWork *ProofOfWork) Run() ([]byte, int64) {

	//1. 将Block的属性拼接成字节数组

	//2. 生成hash

	//3. 判断hash有效性，如果满足条件，跳出循环

	nonce := 0

	var hashInt big.Int // 存储我们新生成的hash
	var hash [32]byte

	for {
		//准备数据
		dataBytes := proofOfWork.prepareData(nonce)

		// 生成hash
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x", hash)

		// 将hash存储到hashInt
		hashInt.SetBytes(hash[:])

		//判断hashInt是否小于Block里面的target
		// Cmp compares x and y and returns:
		//
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		if proofOfWork.target.Cmp(&hashInt) == 1 {
			break
		}

		nonce = nonce + 1
	}

	return hash[:], int64(nonce)
}

// 创建新的工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {

	//1.big.Int对象 1
	// 2
	//0000 0001
	// 8 - 2 = 6
	// 0100 0000  64
	// 0010 0000
	// 0000 0000 0000 0001 0000 0000 0000 0000 0000 0000 .... 0000

	//1. 创建一个初始值为1的target

	// int Int
	target := big.NewInt(1) // 1

	//2. 左移256 - targetBit

	target = target.Lsh(target, 256-targetBit)

	return &ProofOfWork{block, target}
}
