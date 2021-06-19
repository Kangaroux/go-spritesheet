# go-spritesheet

Use YAML to describe your sprite sheet and give your sprites names. `go-spritesheet` will take care of the math for each sprite location, as well as provide a map for easy sprite lookups.

## Using With Ebiten

`go-spritesheet` pairs nicely with a library like [ebiten](https://github.com/hajimehoshi/ebiten).

```go
// Load the config
sheet, err := spritesheet.OpenAndRead("spritesheet.yml")

if err != nil {
    panic(err)
}

// Load the image
img, _, err := ebitenutil.NewImageFromFile(sheet.Image)

if err != nil {
    panic(err)
}

sprites := sheet.Sprites()

// Get the sprite
s := img.SubImage(sprites["mySprite"].Rect())
```

## YAML Example

```yaml
image: hero.png

rows: 3
cols: 4
size: 16

sprites: [
    idle_1, idle_2, idle_3, idle_4,
    run_1, run_2, run_3, run_4,
    atk_1, atk_2, atk_3, atk_4
]
```

## Config Format

- `image`: The path to your sprite sheet image.
- `rows`: The number of rows in the sprite sheet.
- `cols`: The number of columns in the sprite sheet.
- `size`: The size of each sprite, in pixels.
- `sprites`: A list of sprite names.
    - Names must be unique.
    - Must contain `0` to `n` entries, where `n` is `rows*cols`
    - Using an underscore `_` skips the sprite. Useful if you have "holes" in your sprite sheet.
