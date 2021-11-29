package game

import (
	"bytes"
	"fmt"
	"image"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Width      int
	Height     int
	NoiseImage *image.RGBA

	AudioBuffer  *bytes.Buffer
	AudioContext *audio.Context
	Player       *audio.Player
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.ReplacePixels(g.NoiseImage.Pix)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f", ebiten.CurrentTPS(), ebiten.CurrentFPS()))
}

func (g *Game) Update() error {
	s := g.Width * g.Height
	for i := 0; i < s; i++ {
		x := rand.Uint32()
		g.NoiseImage.Pix[4*i] = uint8(x >> 24)
		g.NoiseImage.Pix[4*i+1] = uint8(x >> 24)
		g.NoiseImage.Pix[4*i+2] = uint8(x >> 24)
		g.NoiseImage.Pix[4*i+3] = 0xff
	}
	g.AudioBuffer.Write(g.NoiseImage.Pix)

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func New(width, height int) (*Game, error) {
	buf := bytes.NewBuffer(make([]byte, 0, 48000))
	ctx := audio.NewContext(48000)
	oggS, err := vorbis.DecodeWithSampleRate(48000, buf)
	if err != nil {
		return nil, fmt.Errorf("failed to decode sample: %w", err)
	}

	s := audio.NewInfiniteLoop(oggS, 48000)
	player, err := ctx.NewPlayer(s)

	if err != nil {
		return nil, err
	}

	go player.Play()
	return &Game{
		Width:      width,
		Height:     height,
		NoiseImage: image.NewRGBA(image.Rect(0, 0, width, height)),

		AudioBuffer:  buf,
		Player:       player,
		AudioContext: ctx,
	}, nil
}
