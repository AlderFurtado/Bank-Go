package main

import (
	"AlderFurtado/BankGo.git/infra/db"
	messagebroker "AlderFurtado/BankGo.git/infra/message_broker"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type Message struct {
	Cpf   string `json:"cpf"`
	Value int    `json:"value"`
}

func WorkProcessor(messages chan string) {
	for msg := range messages {
		fmt.Println("Processando mensagem:", msg)
		var m Message

		// Converte string JSON para struct
		err := json.Unmarshal([]byte(msg), &m)
		if err != nil {
			fmt.Print(err)
			return
		}
		value, err := GetUserBalanceValueByCPF(db.GetPostgresDb(), m.Cpf)
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Println("Saldo atual:", value)
		intValue, err := strconv.Atoi(value)
		if err != nil {
			fmt.Print(err)
			return
		}
		m.Value += intValue

		fmt.Println("Novo saldo:", m.Value)
		err = UpdateBalanceValueByCPF(db.GetPostgresDb(), m.Cpf, fmt.Sprintf("%d", m.Value))
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Println("Mensagem processada", msg)
	}
}

func GetUserBalanceValueByCPF(db *sql.DB, cpf string) (string, error) {
	var value string
	err := db.QueryRow(`
		SELECT b.value
		FROM balance b
		JOIN usuario u ON b.id = u.balance_id
		WHERE u.cpf = $1
	`, cpf).Scan(&value)
	if err != nil {
		return "", err
	}
	return value, nil
}

func UpdateBalanceValueByCPF(db *sql.DB, cpf string, newValue string) error {
	_, err := db.Exec(`
        UPDATE balance
        SET value = $1
        WHERE id = (
            SELECT balance_id FROM usuario WHERE cpf = $2
        )
    `, newValue, cpf)
	return err
}

func main() {
	// Canal que vai receber mensagens do Kafka
	messages := make(chan string, 100)

	// Contexto para cancelamento
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Captura Ctrl+C para encerrar
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigchan
		fmt.Println("\nEncerrando...")
		cancel()
	}()

	// Goroutine para consumir do Kafka
	go func() {
		for {
			msg, err := messagebroker.Consumer()
			if err != nil {
				if ctx.Err() != nil {
					break // Contexto cancelado
				}
				log.Println("Erro lendo mensagem:", err)
				continue
			}
			// Envia mensagem para o canal
			messages <- string(msg)
		}
	}()

	// Goroutine para processar mensagens do canal
	WorkProcessor(messages)

	// Mantém o main vivo
	<-ctx.Done()
	close(messages)
	fmt.Println("Programa finalizado")
}
