package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Block struct {
	PrevHash  string
	CurrHash  string
	Data      string
	Nonce     uint64
	Timestamp int64
	Index     uint64
	Validate  bool
}
type Valid struct {
	Status bool
	Index  uint64
}
type dem struct {
	Str string
}
type TransInfo struct {
	Sender  string
	Reciver string
	Amount  uint64
	Miner   string
}
type CryptoBlock struct {
	PrevHash  string
	CurrHash  string
	Nonce     uint64
	Timestamp int64
	Index     uint64
	Validate  bool
	Data      TransInfo
}

type Wallet struct {
	Name    string
	Amount  uint64
	Address string
}

var bchain []Block
var cryptochain []CryptoBlock
var wallets []Wallet

var diff uint64 = 3
var cryptodiff uint64 = 3

func main() {
	port := os.Getenv("PORT")
	port = "2222"
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	Opp := []dem{}
	Opp = append(Opp, dem{"A"})
	router := gin.New()
	//router.Use(cors.Default())
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.html")

	router.Static("/templates", "./templates")

	info := TransInfo{"0x000000", "0x111111", 50, "0x333333"}
	cryptochain = CryptoAddBlock(info, "0", cryptochain)

	info = TransInfo{"0x000000", "0x222222", 50, "0x111111"}
	cryptochain = CryptoAddBlock(info, cryptochain[0].CurrHash, cryptochain)

	info = TransInfo{"0x000000", "0x333333", 50, "0x222222"}
	cryptochain = CryptoAddBlock(info, cryptochain[1].CurrHash, cryptochain)

	steve := Wallet{"Steve", 0, "0x111111"}
	mark := Wallet{"Mark", 0, "0x222222"}
	bill := Wallet{"Bill", 0, "0x333333"}

	wallets = append(wallets, steve)
	wallets = append(wallets, mark)
	wallets = append(wallets, bill)

	router.GET("/", func(c *gin.Context) {

		fmt.Println(bchain)
		fmt.Println(Opp)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/reset", func(c *gin.Context) {
		bchain = []Block{}
		c.JSON(200, gin.H{"blockchain": bchain})
	})
	router.GET("/addblock", func(c *gin.Context) {

		if len(bchain) == 0 {
			bchain = AddBlock("0", "0", bchain)

		} else {
			fmt.Println(c.Query("ids"))
			bchain = AddBlock(bchain[len(bchain)-1].CurrHash, c.Query("data"), bchain)

		}

		c.JSON(200, gin.H{"blockchain": bchain})
	})
	router.GET("/setblock", func(c *gin.Context) {

		i, err := strconv.Atoi(c.Query("index"))
		fmt.Println(i, err)

		bchain = SetBlock(bchain, uint64(i), c.Query("data"))

		c.JSON(200, gin.H{"blockchain": bchain})
	})

	router.GET("/validate", func(c *gin.Context) {
		ReMine(bchain)
		bchain = ValidateChain(bchain)
		c.JSON(200, gin.H{"blockchain": bchain})
	})

	router.GET("/demo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "a.html", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/cryptodemo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "crypto.html", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/getchain", func(c *gin.Context) {

		c.JSON(200, gin.H{"blockchain": bchain})

	})

	router.GET("/getcchain", func(c *gin.Context) {

		c.JSON(200, gin.H{"blockchain": cryptochain})

	})
	router.GET("/changediff", func(c *gin.Context) {

		i, err := strconv.Atoi(c.Query("diff"))

		if err != nil {
			if uint64(i) > 0 && uint64(i) < 6 {
				diff = uint64(i)

			}

		} else {

		}

		c.JSON(200, gin.H{"blockchain": bchain})

	})

	router.GET("/cryptoreset", func(c *gin.Context) {
		cryptochain = []CryptoBlock{}
		info := TransInfo{"0x000000", "0x111111", 50, "0x222222"}
		cryptochain = CryptoAddBlock(info, "0", cryptochain)

		info = TransInfo{"0x000000", "0x222222", 50, "0x111111"}
		cryptochain = CryptoAddBlock(info, cryptochain[0].CurrHash, cryptochain)

		info = TransInfo{"0x000000", "0x333333", 50, "0x222222"}
		cryptochain = CryptoAddBlock(info, cryptochain[1].CurrHash, cryptochain)

		c.JSON(200, gin.H{"blockchain": cryptochain})
	})

	router.GET("/setdata", func(c *gin.Context) {

		i, err := strconv.Atoi(c.Query("index"))

		j := c.Query("data")

		fmt.Println(err)
		SetBlock(bchain, uint64(i), j)

		c.JSON(200, gin.H{"blockchain": bchain})

	})

	router.GET("/crypto", func(c *gin.Context) {

		wallets = CryptoCalculate(cryptochain)
		c.JSON(200, gin.H{"blockchain": cryptochain, "wallets": wallets})

	})

	router.GET("/cryptosetdata", func(c *gin.Context) {

		i, err := strconv.Atoi(c.Query("index"))
		amnt, erri := strconv.Atoi(c.Query("amount"))
		fmt.Println(i, err, erri)
		info := TransInfo{c.Query("sender"), c.Query("reciver"), uint64(amnt), c.Query("miner")}
		cryptochain = CryptoSetBlock(cryptochain, uint64(i), info)

		c.JSON(200, gin.H{"blockchain": cryptochain})

	})

	router.GET("/cryptoadd", func(c *gin.Context) {

		amnt, erri := strconv.Atoi(c.Query("amount"))
		fmt.Println(erri)
		info := TransInfo{c.Query("sender"), c.Query("reciver"), uint64(amnt), c.Query("miner")}

		cryptochain = CryptoAddBlock(info, c.Query("prevhash"), cryptochain)
		wallets = CryptoCalculate(cryptochain)

		c.JSON(200, gin.H{"blockchain": cryptochain, "wallets": wallets})

	})

	router.GET("/cryptovalidate", func(c *gin.Context) {
		CryptoReMine(cryptochain)
		cryptochain = CryptoValidateChain(cryptochain)
		c.JSON(200, gin.H{"blockchain": cryptochain})
	})

	router.GET("/cryptochangediff", func(c *gin.Context) {

		i, err := strconv.Atoi(c.Query("diff"))
		cryptodiff = uint64(i)
		fmt.Println(err)
		c.JSON(200, gin.H{"blockchain": cryptochain})

	})

	router.GET("/wallets", func(c *gin.Context) {

		wallets = CryptoCalculate(cryptochain)
		c.JSON(200, gin.H{"wallet": wallets})

	})

	router.Run(":" + port)
}

//TO GENERATE MD5 HASH
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

//PROOF OF WORK FOR MINIG FINDING Nonce
func ProofOfWork(str string) Block {

	b := Block{}
	var non bytes.Buffer
	for i := 0; i < int(diff); i++ {
		non.WriteString("0")
	}
	i := 0
	for i < 1000000 {
		o := strconv.Itoa(i)
		var buffer bytes.Buffer
		buffer.WriteString(str)
		buffer.WriteString(o)
		hash := GetMD5Hash(buffer.String())
		substring := hash[0:diff]
		if substring == non.String() {

			//fmt.Println(hash)
			//fmt.Println(substring)
			//fmt.Println(i)
			b.CurrHash = hash
			b.Nonce = uint64(i)

			break
		}
		i++
	}

	return b
}

//ADD BLOCK TO BLOCKCHAIN

func AddBlock(PrevHash string, Data string, blockchain []Block) []Block {

	var buffer bytes.Buffer
	buffer.WriteString(PrevHash)
	buffer.WriteString(Data)
	b := ProofOfWork(buffer.String())
	b.PrevHash = PrevHash
	b.Data = Data
	msec := time.Now().UnixNano() / 1000000
	b.Timestamp = msec
	b.Validate = true
	blockchain = append(blockchain, b)

	return blockchain
}

func SetBlock(blockchain []Block, pos uint64, Data string) []Block {

	blockchain[pos].Data = Data
	return blockchain
}

func ReMine(blockchain []Block) {
	for i := 0; i < len(blockchain); i++ {
		var buffer bytes.Buffer
		buffer.WriteString(blockchain[i].PrevHash)
		buffer.WriteString(blockchain[i].Data)
		b := ProofOfWork(buffer.String())
		b.Data = blockchain[i].Data
		b.PrevHash = blockchain[i].PrevHash
		b.Timestamp = blockchain[i].Timestamp
		blockchain[i] = b
	}

}

func ValidateChain(blockchain []Block) []Block {
	Vali := Valid{}
	Vali.Status = true
	for i := 0; i < len(blockchain)-1; i++ {

		if blockchain[i].CurrHash == blockchain[i+1].PrevHash {
			blockchain[i].Validate = true
		} else {

			for j := i; j < len(blockchain); j++ {

				blockchain[j].Validate = false
			}
			Vali.Index = uint64(i)
			Vali.Status = false
			break
		}
	}

	return blockchain
}

///ADD BLOCK TO CRYPTOCHAIN

func CryptoAddBlock(Data TransInfo, Prevhash string, blockchain []CryptoBlock) []CryptoBlock {

	var buffer bytes.Buffer

	buffer.WriteString(Data.Sender)
	buffer.WriteString(Data.Reciver)
	buffer.WriteString(string(Data.Amount))

	msec := time.Now().UnixNano() / 1000000
	b := ProofOfWork2(buffer.String())
	b.PrevHash = Prevhash
	b.Data = Data
	b.Timestamp = msec
	b.Validate = true
	blockchain = append(blockchain, b)

	return blockchain
}

func ProofOfWork2(str string) CryptoBlock {

	b := CryptoBlock{}
	var non bytes.Buffer
	for i := 0; i < int(cryptodiff); i++ {
		non.WriteString("0")
	}
	i := 0
	for i < 1000000 {
		o := strconv.Itoa(i)
		var buffer bytes.Buffer
		buffer.WriteString(str)
		buffer.WriteString(o)
		hash := GetMD5Hash(buffer.String())
		substring := hash[0:diff]
		if substring == non.String() {

			//fmt.Println(hash)
			//fmt.Println(substring)
			//fmt.Println(i)
			b.CurrHash = hash
			b.Nonce = uint64(i)

			break
		}
		i++
	}

	return b
}

func CryptoValidateChain(blockchain []CryptoBlock) []CryptoBlock {
	Vali := Valid{}
	Vali.Status = true
	for i := 0; i < len(blockchain)-1; i++ {

		if blockchain[i].CurrHash == blockchain[i+1].PrevHash {
			blockchain[i].Validate = true
		} else {

			for j := i; j < len(blockchain); j++ {

				blockchain[j].Validate = false
			}
			Vali.Index = uint64(i)
			Vali.Status = false
			break
		}
	}

	return blockchain
}

/////REMINE

func CryptoReMine(blockchain []CryptoBlock) {
	for i := 0; i < len(blockchain); i++ {
		var buffer bytes.Buffer

		buffer.WriteString(blockchain[i].Data.Sender)
		buffer.WriteString(blockchain[i].Data.Reciver)
		buffer.WriteString(string(blockchain[i].Data.Amount))

		b := ProofOfWork2(buffer.String())
		b.Data = blockchain[i].Data
		b.PrevHash = blockchain[i].PrevHash
		b.Timestamp = blockchain[i].Timestamp
		blockchain[i] = b
	}

}

func CryptoSetBlock(blockchain []CryptoBlock, pos uint64, Data TransInfo) []CryptoBlock {

	blockchain[pos].Data = Data
	return blockchain
}

func CryptoCalculate(blockchain []CryptoBlock) []Wallet {

	var a uint64 = 0
	var b uint64 = 0
	var c uint64 = 0

	for i := 0; i < len(blockchain); i++ {

		if blockchain[i].Data.Reciver == "0x111111" {
			a = a + blockchain[i].Data.Amount
			wallets[0].Amount = a

		}
		if blockchain[i].Data.Reciver == "0x222222" {
			b = b + blockchain[i].Data.Amount
			wallets[1].Amount = b

		}
		if blockchain[i].Data.Reciver == "0x333333" {
			c = c + blockchain[i].Data.Amount
			wallets[2].Amount = c

		}

		////////////////////////////////////

		if blockchain[i].Data.Sender == "0x111111" {
			a = a - blockchain[i].Data.Amount
			wallets[0].Amount = a

		}

		if blockchain[i].Data.Sender == "0x222222" {

			b = b - blockchain[i].Data.Amount
			wallets[1].Amount = b

		}
		if blockchain[i].Data.Sender == "0x333333" {
			c = c - blockchain[i].Data.Amount
			wallets[2].Amount = c

		}

		//////////////////////////////////////

		if blockchain[i].Data.Miner == "0x111111" {
			a++
			wallets[0].Amount = a

		}

		if blockchain[i].Data.Miner == "0x222222" {
			b++
			wallets[1].Amount = b

		}
		if blockchain[i].Data.Miner == "0x333333" {
			c++
			wallets[2].Amount = c

		}

	}
	return wallets

}
