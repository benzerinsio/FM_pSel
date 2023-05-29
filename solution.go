package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
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

type Products struct {
	XMLName xml.Name `xml:"produtos"`
	Items   []Item   `xml:"item"`
}

func main() {
	// Read the content of the XML file
	content, err := ioutil.ReadFile("psel.xml")
	if err != nil {
		fmt.Println("Error reading the file: ", err)
		return
	}

	var products Products

	// Decode the XML content
	err = xml.Unmarshal(content, &products)
	if err != nil {
		fmt.Println("Error decoding the XML: ", err)
		return
	}

	// Creating a Map to group de undesired ids for further remove them
	removedIDs := map[int]struct{}{
		403921: {},
		595044: {},
	}

	//New Array for storing the updated informations
	var newItems []Item

	// Getting only the value item for the iteration
	for _, item := range products.Items {
		// Checking if the value is present in the undesired ids (don't need the value itself)
		if _, ok := removedIDs[item.ID]; !ok {
			// Extract the product name from the link to compare with the image
			productName := strings.TrimPrefix(item.Link, "www.loja.com.br/p/")
			productName = strings.Trim(productName, "\"")
			productName += ".jpg"

			// Extract the image name from the image link to compare with the product
			imageName := strings.TrimPrefix(item.ImageLink, "www.loja.com.br/imagens/")
			imageName = strings.Trim(imageName, "\"")

			// Check if the product name and image name are different and correcting if the case
			if productName != imageName {
				item.ImageLink = "\"www.loja.com.br/imagens/" + productName + "\""
			}

			// Replace "BRL" with "R$" in the price field
			item.Price = strings.ReplaceAll(item.Price, "BRL", "R$")

			// Appending the updated values (item) into the newItems array
			newItems = append(newItems, item)
		}
	}

	// Updating the changes to the struct
	products.Items = newItems

	// Printing the new feed
	for index, item := range products.Items {
		fmt.Println("Product ", index+1)
		fmt.Println("ID: ", item.ID)
		fmt.Println("Title: ", item.Title)
		fmt.Println("Price: ", item.Price)
		fmt.Println("Link: ", item.Link)
		fmt.Println("Image Link: ", item.ImageLink)
		fmt.Println("Product Type: ", item.ProductType)
		fmt.Println("--------------------------------------------------")
	}

	// Encode the updated structure back to XML
	xmlContent, err := xml.MarshalIndent(products, "", "    ")
	if err != nil {
		fmt.Println("Error encoding the XML:", err)
		return
	}

	// Write the updated content to a new file
	err = ioutil.WriteFile("file.xml", xmlContent, 0644)
	if err != nil {
		fmt.Println("Error writing the file:", err)
		return
	}
}
