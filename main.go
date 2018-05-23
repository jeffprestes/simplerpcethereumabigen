package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jeffprestes/simplerpcethereumabigen/livrovisitas"
)

var client *ethclient.Client
var contract *livrovisitas.LivroVisitas
var logFile *os.File
var mw io.Writer
var ctx context.Context

func main() {

	var err error
	ctx = context.Background()

	logFile, err = os.OpenFile(os.Getenv("LOGFILE"), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalln(err.Error(), " - error creating log file")
		return
	}
	mw = io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	log.Println("Conectando-se ao nó...")
	// Create an IPC based RPC connection to a remote node
	//client, err = ethclient.Dial("/Users/jeffprestes/Library/Ethereum/geth.ipc")
	client, err = ethclient.Dial(os.Getenv("ETH_PROTOCOL") + "://" + os.Getenv("ETH_IP") + ":" + os.Getenv("ETH_PORT"))
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	log.Println("Criando uma instancia do contrato...")
	// Instantiate the contract and display its name
	contract, err = livrovisitas.NewLivroVisitas(common.HexToAddress("0x2092f46e073fff25a316437ec3f986ccabe6b135"), client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}

	log.Println("Iniciando o servidor web...")
	http.HandleFunc("/", executeContractMethod)
	http.ListenAndServe(os.Getenv("WEB_IP")+":"+os.Getenv("WEB_PORT"), nil)
	log.Println("Servidor web iniciado...")
}

func executeContractMethod(w http.ResponseWriter, r *http.Request) {
	log.SetOutput(mw)
	log.Println(" Execute Contract Method called...")
	fmt.Fprintln(w, "<html><head><title>Calling smart contract via web</title><head><body>")

	addr, err := contract.WhoIsTheOwner(nil)
	if err != nil {
		log.Printf("Failed to get Owner contract: %v", err)
		fmt.Fprintln(w, "<h5>Não foi possivel obter o owner </h5>")
		footer(w)
		return
	}
	fmt.Fprintln(w, "<h5>Dono do contrato: ", addr.String(), "</h5>")

	ultimaVisita, err := contract.GetMessageOfVisit(nil, common.HexToAddress("0x7144b1f4454a7056f7dd0089ef507722b9964cb9"))
	if err != nil {
		log.Printf("Failed to get last visit: %v", err)
		fmt.Fprintln(w, "<h5>Não foi possivel obter a ultima visita </h5>")
		footer(w)
		return
	}
	fmt.Fprintln(w, "<h3>Ultima visita: ", ultimaVisita, "</h3>")

	var dadoARegistrar = "Teste de visita " + time.Now().Format("15:04:05")
	fmt.Fprintln(w, "<h3>Agora vou registrar uma nova visita para voce verificar na proxima esse valor: ", dadoARegistrar)

	//file := "/Users/jeffprestes/ethereum7masters/keystore/UTC--2018-05-15T18-24-40.890152706Z--3096dc394a540d8f856c6840f84fe1d511c05d30"
	passphrase := os.Getenv("ETH_SENHA")
	keyjson, err := os.Open(os.Getenv("ETH_ACCOUNT"))
	if err != nil {
		log.Printf("Failed to load the json key file: %v", err)
		fmt.Fprintln(w, "<h5>Não foi obter a chave da conta principal </h5>")
		footer(w)
		return
	}

	auth, err := bind.NewTransactor(keyjson, passphrase)
	if err != nil {
		log.Printf("Failed to load the key: %v", err)
		fmt.Fprintln(w, "<h5>Não foi obter a chave da conta principal </h5>")
		footer(w)
		return
	}
	auth.GasLimit = 360450
	auth.GasPrice = big.NewInt(8000200000)

	tx, err := contract.RecordVisit(auth, common.HexToAddress("0x7144b1f4454a7056f7dd0089ef507722b9964cb9"), dadoARegistrar)
	if err != nil {
		log.Printf("Failed to save last visit: %v", err)
		fmt.Fprintln(w, "<h5>Não foi possivel salvar a ultima visita </h5>")
		footer(w)
		return
	}
	log.Printf("Transaction: %+v - Context: %+v - Dado: %s\n", tx.Hash().String(), ctx, dadoARegistrar)
	fmt.Fprintf(w, "<h5>Detalhes da transação do registro ultuma visita: %+v</h5>\n", tx)

	var txStatus bool
	if ctx != nil && len(tx.Hash()) > 10 {
		//recibo, err := client.TransactionReceipt(ctx, tx.Hash())
		tx, txStatus, err = client.TransactionByHash(ctx, tx.Hash())
		if err != nil {
			log.Printf("Failed to get transaction receipt: %v", err)
			fmt.Fprintln(w, "<h5>Não foi obter o status da ultima transacao </h5>")
			footer(w)
			return
		}
		fmt.Fprintf(w, "<h5>Status da transação do registro ultuma visita: %+v</h5><h5>Detalhes: %+v</h5>\n", txStatus, tx)
	}

	/*
		log.Println("Verificando logs...")
		opt := &bind.FilterOpts{}
		events, err := contract.LivroVisitasFilterer.FilterNewVisitor(opt)
		if err != nil {
			log.Printf("Failed to get the events of visit: %v", err)
			fmt.Fprintln(w, "<h5>Não foi possivel os eventos da transacao da visita </h5>")
			footer(w)
			return
		}
		for events.Next() {
			if events.Error() == nil {
				fmt.Fprintf(w, "<h5>Há um novo visitante: %+v</h5>\n", events.Event.Visitor.String())
			} else if events.Error() != nil {
				log.Printf("Failed during visit registering: %v", err)
				fmt.Fprintln(w, "<h5>Erro no registro da visita </h5>")
			}
		}
	*/

	log.Println("Abrindo watcher...")
	go watcher(contract)
	footer(w)
}

func footer(w http.ResponseWriter) {
	fmt.Fprintln(w, "</body></html>")
	log.Println(" Execute Contract Method finished")
}

func watcher(contract *livrovisitas.LivroVisitas) {
	var blockNumber uint64 = 5692
	ch := make(chan *livrovisitas.LivroVisitasNewVisitor)
	opts := &bind.WatchOpts{}
	opts.Start = &blockNumber
	s, err := contract.WatchNewVisitor(opts, ch)
	if err != nil {
		log.Printf("Failed WatchNewVisitor: %v", err)
		return
	}
	select {
	case err := <-s.Err():
		log.Printf("Failed to WatchNewVisitor in select: %+v", err)
	case newEvent := <-ch:
		log.Println("Ha um novo visitante: ", newEvent.Visitor.String())
		recibo, err := client.TransactionReceipt(ctx, newEvent.Raw.TxHash)
		if err == nil {
			json, _ := recibo.MarshalJSON()
			log.Printf("Watch Recibo: %+s\n", string(json))
		} else {
			log.Printf("Erro no Recibo: %s\n", err.Error())
		}
	}
}
