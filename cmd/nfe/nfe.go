package nfe

import (
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

type NFe struct {
	XMLName xml.Name `xml:"nfeProc"`
	NFe     struct {
		InfNFe struct {
			Total struct {
				ICMSTot struct {
					VNF string `xml:"vNF"`
				} `xml:"ICMSTot"`
			} `xml:"total"`

			Dest struct {
				CNPJ string `xml:"CNPJ"`
				CPF  string `xml:"CPF"`
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
				return err
			}

			//for _, item := range nfe.NFe.InfNFe.Det {
			//	if strings.Contains(strings.ToLower(item.Prod.XProd), "trelicada") {
			//		fmt.Println("Produto com trelicada encontrado: " + item.Prod.XProd)
			//
			//		nomeArquivoCopia := "xml_encontrada_" + filepath.Base(path)
			//		caminhoCopia := filepath.Join(diretorio, nomeArquivoCopia)
			//		err := copiarArquivo(path, caminhoCopia)
			//		if err != nil {
			//			fmt.Println("Erro ao copiar arquivo:", err)
			//		} else {
			//			fmt.Println("Arquivo copiado com sucesso:", nomeArquivoCopia)
			//		}
			//		break
			//	}
			//}

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
