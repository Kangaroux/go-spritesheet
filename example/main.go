package main

import (
	"fmt"

	"github.com/Kangaroux/go-spritesheet"
)

func main() {
	sheet, err := spritesheet.OpenAndRead("spritesheet.yml")

	if err != nil {
		panic(err)
	}

	sprites := sheet.Sprites()

	fmt.Printf("%12s: %s\n", "image", sheet.Image)
	fmt.Printf("%12s: %dpx\n", "sprite size", sheet.Size)
	fmt.Printf("%12s: %d\n", "# of sprites", len(sprites))
	fmt.Printf("%12s: %d\n", "rows", sheet.Rows)
	fmt.Printf("%12s: %d\n", "cols", sheet.Cols)
	fmt.Println("\nsprites:")

	for _, sprite := range sprites {
		rect := sprite.Rect()
		fmt.Printf("%12s: %v to %v\n", sprite.Name, rect.Min, rect.Max)
	}
}
