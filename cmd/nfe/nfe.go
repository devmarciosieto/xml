package nfe

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type NFe struct {
	XMLName xml.Name `xml:"nfeProc"`
	NFe     struct {
		InfNFe struct {
			Versao string `xml:"versao,attr"`
			Total  struct {
				ICMSTot struct {
					VNF string `xml:"vNF"`
				} `xml:"ICMSTot"`
			} `xml:"total"`

			Dest struct {
				CNPJ  string `xml:"CNPJ"`
				CPF   string `xml:"CPF"`
				XNome string `xml:"xNome"` // Campo para o nome do cliente ou razão social
			} `xml:"dest"`

			Det []struct {
				Prod struct {
					CProd  string `xml:"cProd"`
					XProd  string `xml:"xProd"`
					QCom   string `xml:"qCom"`
					VUnCom string `xml:"vUnCom"`
				} `xml:"prod"`
			} `xml:"det"`
		} `xml:"infNFe"`
	} `xml:"NFe"`
}

func SomaValoresNotas(diretorio string) (float64, error) {
	var somaTotal float64

	err := filepath.Walk(diretorio, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".xml" {
			dados, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			var nfe NFe
			err = xml.Unmarshal(dados, &nfe)
			if err != nil {
				// Ignorar arquivos que não podem ser deserializados como NFe
				return nil
			}

			// Verificar se o arquivo é um padrão NFe 4.0
			if !isNFe40(nfe) {
				return nil
			}

			for _, item := range nfe.NFe.InfNFe.Det {

				//nomeCliente := nfe.NFe.InfNFe.Dest.XNome
				cnpj_find := nfe.NFe.InfNFe.Dest.CNPJ

				//if strings.Contains(item.Prod.XProd, "LAJE TRELICADA TIPO FORRO BETA 11 REFORCADO") {
				//if strings.Contains(nomeCliente, "TRANSMARTINS TRANSPORTE DE CARGAS E LOGISTICA LTDA") {
				if strings.Contains(cnpj_find, "01231592000115") {
					fmt.Println("Produto com trelicada encontrado: " + item.Prod.XProd)

					nomeArquivoCopia := "xml_encontrada_" + filepath.Base(path)
					caminhoCopia := filepath.Join(diretorio, nomeArquivoCopia)
					err := copiarArquivo(path, caminhoCopia)
					if err != nil {
						fmt.Println("Erro ao copiar arquivo:", err)
					} else {
						fmt.Println("Arquivo copiado com sucesso:", nomeArquivoCopia)
					}
					break
				}
			}

			valor, err := strconv.ParseFloat(nfe.NFe.InfNFe.Total.ICMSTot.VNF, 64)
			if err != nil {
				return err
			}

			somaTotal += valor
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return somaTotal, nil
}

func copiarArquivo(src, dst string) error {
	entrada, err := os.Open(src)
	if err != nil {
		return err
	}
	defer entrada.Close()

	saida, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer saida.Close()

	_, err = io.Copy(saida, entrada)
	return err
}

func isNFe40(nfe NFe) bool {
	// Verifique se a estrutura 'nfe' contém campos específicos da NFe 4.0
	// Exemplo: Verificar a versão do layout da NFe
	versaoLayout := nfe.NFe.InfNFe.Versao
	if versaoLayout == "4.00" {
		return true
	}

	// Adicione aqui outras verificações específicas necessárias para o padrão NFe 4.0

	return false
}
