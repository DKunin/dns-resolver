package main

import (
    "fmt"
    "github.com/nilslice/email"
)

func main() {
    msg := email.Message{
        To: "info@dkunin.ru", // do not add < > or name in quotes
        From: "me@server.name", // do not add < > or name in quotes
        Subject: "A simple email",
        Body: "Plain text email body. HTML not yet supported, but send a PR!",
    }

    err := msg.Send()
    if err != nil {
        fmt.Println(err)
    }
}
