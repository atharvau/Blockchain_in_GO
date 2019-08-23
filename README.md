<p align="center">
<img height=100px src="https://i.imgur.com/ZoSs6PW.pnghttps://i.imgur.com/ZoSs6PW.png" />  
<h1 align="center"> Blockchain Demonstartion in GOLANG </h1>
<h4 align="center">The Aim of The Project is provide a Visual Simualtion of Blockchain World for those who are new to this Technology  </h4>
<h4 align="center">Live Demo  <a href="https://blockchaingo.herokuapp.com/">https://blockchaingo.herokuapp.com/ </a>  </h4>

</p>


![alt text](https://img.shields.io/badge/GOLANG-1.12.9-brightgreen)  ![alt text](https://img.shields.io/badge/Gin%20Web%20Framework-1.4.0-blue) ![alt text](https://img.shields.io/badge/HTML-5-red)  ![alt text](https://img.shields.io/badge/Jquery%20-3.1.0-yellow) ![alt text](https://img.shields.io/badge/Heroku-Free%20Dyno-lightgrey)  ![alt text](https://img.shields.io/badge/license-MIT-green)  ![alt text](https://img.shields.io/badge/build-passing-brightgreen)
 



# Blockchain In GOLANG

## Live Demo
Checkout the live demo [here](https://blockchaingo.herokuapp.com/)

* The Webapp may take 20 -30 sec to Load
* Prefer Desktop to access the website

>https://blockchaingo.herokuapp.com/


![alt text](https://i.imgur.com/8Aj3Eq5.gif) 


![alt text](https://i.imgur.com/08wIfdG.gif) 

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) version 1.12 or newer.

```sh
$ git clone https://github.com/atharvau/Blockchain_in_GO
$ cd Blockchain_in_GO
$ go get -u github.com/gin-gonic/gin
$ go build Main.go 


```

Your app should now be running on [localhost:2222](http://localhost:5000/).


## Introduction


Blockchain in GoLang is a Golang Project which consists simualtion of two blockchains

* Simple Blockchain
* Crypto Blockchain 

It gives visual feel of the blockchain data structure to easly understand the concepts.
Features of the blockchain are


| Blockchain    | Consensus Algorithm| Stored Data  |
| ------------- |:-------------:| -----:|
| Simple        | Proof-of-Work | String Data |
|  Crypto       | Proof-of-Work   |   Transaction Data |




> An awesome project.

## Technical Details 

### Technology 

1. **BackEnd** : GOLANG-Gin
2. **FrontEnd** : HTML CSS Jquery
3. **Deployement** :Heroku 

### Directory Structure 

 ```
project
│   README.md
│   main.go   
│   Procfile
│   app.json
│   heroku.yml             
└───folder1
	a.html  
	index.html
	crypto.html
	index.css
 ```


The main.go is the Main file which has implementation of Blockchain and a server to serve It.
The Template Folder Consists of All Front End Html Css And JS files required in Front End 
The Remaining files in root directory are autogenrated deployement files required for hosting on Heroku

### Code Structure 
The main.go can be categorized into  1. blockchain implementaion 2.Server Code
The Server Code consits typical GET Request and Responses for handling incoming requests from FrontEnd
The Blockchain Implementation consists 2 Chains Structure
1. Simple


* Structure

```
type Block struct {
	PrevHash  string
	CurrHash  string
	Data      string
	Nonce     uint64
	Timestamp int64
	Index     uint64
	Validate  bool
}

```

* Methods

```

func AddBlock(PrevHash string, Data string, blockchain []Block) []Block
func SetBlock(blockchain []Block, pos uint64, Data string) []Block 
func ReMine(blockchain []Block)
func ValidateChain(blockchain []Block) []Block
func ProofOfWork
func GetMD5Hash(text string) string


```

2. Crypto


* Structure


```
type CryptoBlock struct {
	PrevHash  string
	CurrHash  string
	Nonce     uint64
	Timestamp int64
	Index     uint64
	Validate  bool
	Data      TransInfo
}
type TransInfo struct {
	Sender  string
	Reciver string
	Amount  uint64
	Miner   string
}
type Wallet struct {
	Name    string
	Amount  uint64
	Address string
}


```

* Methods

```
func CryptoAddBlock(Data TransInfo, Prevhash string, blockchain []CryptoBlock) []CryptoBlock 
func ProofOfWork2(str string) CryptoBlock 
func CryptoValidateChain(blockchain []CryptoBlock) []CryptoBlock 
func CryptoReMine(blockchain []CryptoBlock) 
func CryptoSetBlock(blockchain []CryptoBlock, pos uint64, Data TransInfo) []CryptoBlock 
func CryptoCalculate(blockchain []CryptoBlock) []Wallet

```

## How To Use : Simple Blockchain

### Add New Block

New Block Can be added by inputing a string and click add new data


![alt text](https://i.imgur.com/hGmU5ml.gif)


### Change Difficulty

Difficulty is a measure of how difficult it is to find a hash below a given target


![alt text](https://i.imgur.com/HobOXjK.gif)

### Reset Blockchain

The Blockchain can be cleared easily cleared by clicking


![alt text](https://i.imgur.com/z7OJFXa.gif)


### Corrupting Blockchain 

You can Change the contents of blockchain by inputing data and clicking Change Data Button 
It Corrupts the blockchain and display corrupted Chain in Red Color


![alt text](https://i.imgur.com/GljX13A.gif)

## How To Use :  Crypto Blockchain

### Wallets 

The Crypto chain is provided with 3 default Wallets 

| Name    | Account| Default CryptoCoin |
| ------------- |:-------------:| -----:|
| Steve         | 0x111111 | 51 |
|  Mark       | 0x222222  |  51 |
|  Bill       | 0x333333  |  51 |

> The Mining Reward is 1 CryptoCoin

### Create New Transaction

1. Select The Sender
2. Select The Reciver
3. Select The Miner 
4. Select the Amount to Transfer
5. Select the Transfer button and wait till it mines new Block

![alt text](https://i.imgur.com/08wIfdG.gif) 


## Contributing

* We're are open to enhancements & bug-fixes.
* Feel free to add issues and submit patches.







