package main

import (
	"fmt"
	"github.com/devmarciosieto/xml/cmd/nfe"
	"github.com/devmarciosieto/xml/cmd/util"
)

func main() {

	total, err := nfe.SomaValoresNotas("/home/marcio/apps/back/golang/xml/cmd/xml")
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println(util.FormatarValor(total))
}
