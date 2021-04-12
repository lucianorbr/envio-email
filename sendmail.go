package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Criamos um slice do tipo string do tamanho máximo 1 para receber nosso email de destino
	recipients := make([]string, 1)
	recipients[0] = "email" //email de destino armazendo na posicao '0'

	err := smtp.SendMail(
		/* endereço do servidor SMTP */ "pop.gmail.com:995",
		/* mecanismo de autenticação */ smtp.PlainAuth("", "email", "senha", "pop.gmail.com"),
		/* email de origem */ "email",
		/* mensagem do RFC */ recipients,
		/* corpo do email */ []byte("Subject:Olá!\n\n Olá Luciano. Tudo de bom com Go, teste enviado!"))
	if err != nil {
		log.Fatal(err)
	}
}
