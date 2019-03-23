package main

//
//import (
//	"crypto/sha256"
//	"encoding/hex"
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//// 声明一个块结构体
//type Block struct {
//	Index     int  //区块高度
//	Timestamp string //时间戳，也就是当前时间转换成字符串
//	BPM       int   //要保存的上链数据
//	Hash      string  //当前区块hash值
//	PrevHash  string  //父区块hash值
//	Validator string  //当前区块出块人的账户地址
//}
//
////持币人信息
//type Node struct {
//	tokens int
//	address string
//}
//
////声明一个区块链，类型是切片。也就是将区块放在这个切片里面
//var Blockchain []Block
//
////出块函数，也就是当随机抽中哪个账户后，该账户就会作为该区块的参数，生成一个最新区块
//func generateBlock(oldBlock Block, BPM int, address string) (Block, error) {
//
//	var newBlock Block
//
//	t := time.Now()
//
//	//生成区块过程就是给区块的字段赋值
//	newBlock.Index = oldBlock.Index + 1
//	newBlock.Timestamp = t.String()
//	newBlock.BPM = BPM
//	newBlock.PrevHash = oldBlock.Hash
//	newBlock.Hash = calculateBlockHash(newBlock)
//	newBlock.Validator = address
//
//	//返回最新区块
//	return newBlock, nil
//}
//
//// SHA256算法计算当前区块的hash值
//func calculateHash(s string) string {
//	h := sha256.New()
//	h.Write([]byte(s))
//	hashed := h.Sum(nil)
//	return hex.EncodeToString(hashed)
//}
//
////将区块字段拼接，为生成hash函数做准备
//func calculateBlockHash(block Block) string {
//	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
//	return calculateHash(record)
//}
//
////现在参与挖矿的币总数为15个，我们认为现在是平等机会的15个持币1个的矿工。
////声明一个数组保存15个账户地址
////两个节点参与挖矿
//var N[15] string
//var p [2] Node
//
////进入主函数
//func main() {
//
//	//给两个挖矿节点赋值
//	p[0]=Node{10,"abc"}
//	p[1]=Node{5,"bcd"}
//
//	//通过下边的for循环，按照两个节点他们的持币数，分别将他们的地址赋值到之前定义好的账户地址数组中。
//	//可以看出，前10个账户都是p[0]的；后边5个是p[1]的
//	var cnt = 0
//	for i:=0;i<2;i++ {
//		for j:=0;j<p[i].tokens;j++{
//
//			N[cnt] = p[i].address
//			cnt++
//		}
//	}
//
//	//设置一个随机数种子
//	rand.Seed(time.Now().Unix())
//	var firstBlock Block
//
//	//根据随机数种子产生的随机数，找到矿工的地址，然后出块
//	var b,_ = generateBlock(firstBlock,10,N[rand.Intn(cnt)])
//	//将新的区块放到区块链上。
//	Blockchain = append(Blockchain,b)
//	var b1,_ = generateBlock(Blockchain[len(Blockchain)-1],10,N[rand.Intn(cnt)])
//	//将新的区块放到区块链上。
//	Blockchain = append(Blockchain,b1)
//	fmt.Println(Blockchain)
//}
