package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{
	color.Black,                        // fundo preto
	color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // vermelho
	color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // verde
	color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // azul
	color.RGBA{0xFF, 0xFF, 0x00, 0xFF}, // amarelo
	color.RGBA{0xFF, 0x00, 0xFF, 0xFF}, // magenta
	color.RGBA{0x00, 0xFF, 0xFF, 0xFF}, // ciano
}

const (
	backgroundIndex = 0 // first color in palette (preto)
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	if len(os.Args) > 1 && os.Args[1] == "file" {
		gf, err := os.Create("out.gif")
		if err != nil {
			log.Fatalf("Erro ao criar arquivo: %v", err)
		}
		defer gf.Close()

		lissajous(gf)
		log.Println("GIF gerado com sucesso: out.gif")
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // número de revoluções completas do oscilador x
		res     = 0.001 // resolução angular
		size    = 100   // tamanho da imagem
		nframes = 64    // número de frames da animação
		delay   = 8     // atraso entre frames em unidades de 10ms
	)
	freq := rand.Float64() * 3 // frequência relativa do oscilador y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // diferença de fase
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		// Escolhe uma cor diferente para cada frame
		colorIndex := uint8(i%len(palette)) + 1 // ignora o fundo preto

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		log.Fatalf("Erro ao codificar GIF: %v", err)
	}
}
