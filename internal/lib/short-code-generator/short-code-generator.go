package shortcodegenerator

import (
	"fmt"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

const generatorAlphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Generator struct {
	length   int
	alphabet string
}

func New(length int) *Generator {
	return &Generator{
		length:   length,
		alphabet: generatorAlphabet,
	}
}

func (g *Generator) Generate() (string, error) {
	const op = "shortcodegenerator.Generate"

	shortCode, err := gonanoid.Generate(g.alphabet, g.length)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return shortCode, nil
}
