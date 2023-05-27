package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Item struct {
	ID          int    `xml:"id"`
	Title       string `xml:"title"`
	Price       string `xml:"price"`
	Link        string `xml:"link"`
	ImageLink   string `xml:"image_link"`
	ProductType string `xml:"product_type"`
}

type Produtos struct {
	XMLName xml.Name `xml:"produtos"`
	Items   []Item   `xml:"item"`
}

func main() {
	// Ler o conteúdo do arquivo XML
	content, err := ioutil.ReadFile("psel.xml")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	// Criar uma instância da struct que representa o XML
	var produtos Produtos

	// Decodificar o XML
	err = xml.Unmarshal(content, &produtos)
	if err != nil {
		fmt.Println("Erro ao decodificar o XML:", err)
		return
	}

	// Tratamentos desejados
	idDesejados := [2]int{403921, 595044}
	var novosItens []Item
	for _, item := range produtos.Items {
		adicionar := true
		for _, id := range idDesejados {
			if item.ID == id {
				adicionar = false
				break
			}
		}
		if adicionar {
			nomeProduto := strings.ReplaceAll(item.Link, "www.loja.com.br/p/", "")
			nomeProduto = strings.ReplaceAll(nomeProduto, "\"", "")
			nomeProduto += ".jpg"
			nomeImage := strings.ReplaceAll(item.ImageLink, "www.loja.com.br/imagens/", "")
			nomeImage = strings.ReplaceAll(nomeImage, "\"", "")
			if nomeProduto != nomeImage {
				item.ImageLink = "\"www.loja.com.br/imagens/" + nomeProduto + "\""

				fmt.Println("Corrigido: " + strconv.Itoa(item.ID))
			}
			item.Price = strings.ReplaceAll(item.Price, "BRL", "R$")
			novosItens = append(novosItens, item)
		}
	}

	// Atualizar os itens na struct Produtos
	produtos.Items = novosItens

	// Codificar a estrutura completa de volta para o XML
	xmlContent, err := xml.MarshalIndent(produtos, "", "    ")
	if err != nil {
		fmt.Println("Erro ao codificar o XML:", err)
		return
	}

	// Escrever o conteúdo atualizado no arquivo
	err = ioutil.WriteFile("arquivo.xml", xmlContent, 0644)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
		return
	}
}
