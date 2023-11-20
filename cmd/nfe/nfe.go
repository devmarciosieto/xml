package nfe

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"strconv"
)

// NFe representa a estrutura da Nota Fiscal Eletrônica
type NFe struct {
	XMLName xml.Name `xml:"nfeProc"`
	NFe     struct {
		InfNFe struct {
			Total struct {
				ICMSTot struct {
					VNF string `xml:"vNF"`
				} `xml:"ICMSTot"`
			} `xml:"total"`
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
