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
	Validate bool
}
type Valid struct {
	Status bool
	Index  uint64
}
type dem struct {
	Str string
}
type TransInfo struct {
	Sender string
	Reciver string
	Amount uint64
	Miner string

}
type CryptoBlock struct {
	PrevHash  string
	CurrHash  string
	Nonce     uint64
	Timestamp int64
	Index     uint64
	Validate bool
	Data TransInfo

}

var bchain [] Block
var cryptochain [] CryptoBlock

var diff uint64=3;

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
	router.LoadHTMLGlob("templates/*")


	router.GET("/", func(c *gin.Context) {

		fmt.Println(bchain)
		fmt.Println(Opp)

		c.JSON(200, gin.H{"blockchain": bchain})
	})
	router.GET("/reset", func(c *gin.Context) {
		bchain = [] Block{};


		c.JSON(200, gin.H{"blockchain": bchain})
	})
	router.GET("/addblock", func(c *gin.Context) {


		if(len(bchain)==0){
			bchain = AddBlock("0", "0", bchain)

		}else{
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
		bchain= ValidateChain(bchain)
		c.JSON(200, gin.H{"blockchain": bchain})
	})

	router.GET("/demo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "a.html", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/changediff", func(c *gin.Context) {

		i, err := strconv.Atoi(c.Query("diff"))
		diff= uint64(i);
		fmt.Println(err);

	c.JSON(200, gin.H{"blockchain": bchain})

	})




	router.GET("/setdata", func(c *gin.Context) {

		i, err := strconv.Atoi(c.Query("index"))

		j:=c.Query("data")

		fmt.Println(err);
		SetBlock(bchain, uint64(i),j)

		c.JSON(200, gin.H{"blockchain": bchain})

	})

	router.GET("/crypto", func(c *gin.Context) {

		c.JSON(200, gin.H{"blockchain": cryptochain})
	})

	router.GET("/cryptosetdata", func(c *gin.Context) {

		cryptochain=CryptoAddBlock(TransInfo{"A","B",111,"C"},"@",cryptochain);
		c.JSON(200, gin.H{"blockchain": cryptochain})
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
	b.Validate=true;
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
			blockchain[i].Validate=true;

		} else {

			for j:=i;j<len(blockchain);j++{

				blockchain[j].Validate=false;
			}
			Vali.Index = uint64(i)
			Vali.Status = false
			break
		}
	}

	return blockchain;
}


///ADD BLOCK TO CRYPTOCHAIN

func CryptoAddBlock(Data TransInfo,Prevhash string, blockchain []CryptoBlock) []CryptoBlock {

	var buffer bytes.Buffer

	buffer.WriteString(Data.Sender)
	buffer.WriteString(Data.Reciver)
	buffer.WriteString(string(Data.Amount))

	msec := time.Now().UnixNano() / 1000000
	b := ProofOfWork2(buffer.String())
	b.PrevHash = Prevhash
	b.Data = Data
	b.Timestamp = msec
	b.Validate=true;
	blockchain = append(blockchain, b)

	return blockchain
}




func ProofOfWork2(str string) CryptoBlock {

	b := CryptoBlock{}
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

	return b;
}

///CHANGE CRYPTOCHAIN

func CryptoSetBlock(blockchain []CryptoBlock, pos uint64, Data TransInfo) []CryptoBlock {

	blockchain[pos].Data = Data
	return blockchain
}


func CryptoValidateChain(blockchain []CryptoBlock) []CryptoBlock {
	Vali := Valid{}
	Vali.Status = true
	for i := 0; i < len(blockchain)-1; i++ {

		if blockchain[i].CurrHash == blockchain[i+1].PrevHash {
			blockchain[i].Validate=true;

		} else {

			for j:=i;j<len(blockchain);j++{

				blockchain[j].Validate=false;
			}
			Vali.Index = uint64(i)
			Vali.Status = false
			break
		}
	}

	return blockchain;
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


