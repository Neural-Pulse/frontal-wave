package emailsender

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gopkg.in/gomail.v2"
)

// Produto representa um produto no banco de dados.
type Produto struct {
	ID        int
	Nome      string
	Preco     float64
	UsuarioID int
}

// Usuario representa um usuário no banco de dados.
type Usuario struct {
	ID    int
	Nome  string
	Email string
}

// Função principal para enviar e-mails com detalhes do produto para os usuários
func EnviarEmails() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env:", err)
	}

	// Configuração da conexão com o banco de dados
	db, err := sql.Open("postgres", "user=username password=password dbname=database sslmode=disable")
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	// Consulta à view para obter os dados dos produtos e dos usuários
	rows, err := db.Query("SELECT produto_id, produto_nome, produto_preco, usuario_id, usuario_nome, usuario_email FROM produtos_usuarios_view")
	if err != nil {
		log.Fatal("Erro ao executar a consulta:", err)
	}
	defer rows.Close()

	// Processamento dos resultados da consulta
	var produtos []Produto
	var usuarios map[int]Usuario = make(map[int]Usuario)
	for rows.Next() {
		var produto Produto
		var usuario Usuario
		err := rows.Scan(&produto.ID, &produto.Nome, &produto.Preco, &usuario.ID, &usuario.Nome, &usuario.Email)
		if err != nil {
			log.Fatal("Erro ao ler o resultado da consulta:", err)
		}
		produtos = append(produtos, produto)
		usuarios[usuario.ID] = usuario
	}

	// Envio de e-mails com os produtos para os usuários
	for _, produto := range produtos {
		usuario := usuarios[produto.UsuarioID]
		emailBody := fmt.Sprintf("Olá %s,\n\nAqui estão os detalhes do produto:\n\nID: %d\nNome: %s\nPreço: %.2f\n\nAtenciosamente,\nSua loja", usuario.Nome, produto.ID, produto.Nome, produto.Preco)
		err := sendEmail(usuario.Email, "Detalhes do Produto", emailBody)
		if err != nil {
			log.Printf("Erro ao enviar e-mail para %s: %s", usuario.Email, err)
		} else {
			log.Printf("E-mail enviado com sucesso para %s", usuario.Email)
		}
	}
}

// Função auxiliar para enviar e-mails
func sendEmail(to, subject, body string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPortStr := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Fatal("Erro ao converter porta SMTP para inteiro:", err)
	}

	// Configuração do cliente SMTP
	dialer := gomail.NewDialer(smtpHost, smtpPort, smtpUsername, smtpPassword)

	// Criação da mensagem
	message := gomail.NewMessage()
	message.SetHeader("From", "your-email@example.com")
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)

	// Envio do e-mail
	err = dialer.DialAndSend(message)
	if err != nil {
		return err
	}
	return nil
}
