package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
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
	Items []Item `xml:"item"`
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

	// Remover o item com o ID desejado
	idDesejados := [2]int{403921, 595044}
	var novosItens []Item
	for _, item := range produtos.Items {
		for _, id := range idDesejados {
			if item.ID != id {
				novosItens = append(novosItens, item)
			}
		}
	}

	// Atualizar os itens na struct Produtos
	produtos.Items = novosItens

	// Codificar a estrutura de volta para o XML
	xmlContent, err := xml.MarshalIndent(produtos.Items, "", "    ")
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

	for _, item := range produtos.Items {
		fmt.Println("ID:", item.ID)
		fmt.Println("Title:", item.Title)
		fmt.Println("Price:", item.Price)
		fmt.Println("Link:", item.Link)
		fmt.Println("Image Link:", item.ImageLink)
		fmt.Println("Product Type:", item.ProductType)
		fmt.Println()
	}

}
