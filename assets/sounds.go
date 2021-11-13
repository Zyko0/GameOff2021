package assets

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const (
	defaultSFXVolume   = 0.2
	defaultMusicVolume = 0.4
)

var (
	ctx = audio.NewContext(44100)

	//go:embed sound/mainmenu.wav
	mainmenuMusicBytes  []byte
	mainmenuAudioPlayer *audio.Player
	//go:embed sound/ingame.wav
	ingameMusicBytes  []byte
	ingameAudioPlayer *audio.Player
	//go:embed sound/hit.wav
	hitSoundBytes  []byte
	hitAudioPlayer *audio.Player
)

func init() {
	var err error

	reader, err := wav.Decode(ctx, bytes.NewReader(hitSoundBytes))
	if err != nil {
		log.Fatal(err)
	}
	hitAudioPlayer, err = ctx.NewPlayer(reader)
	if err != nil {
		log.Fatal(err)
	}
	hitAudioPlayer.SetVolume(defaultSFXVolume)

	reader, err = wav.Decode(ctx, bytes.NewReader(ingameMusicBytes))
	if err != nil {
		log.Fatal(err)
	}
	infiniteReader := audio.NewInfiniteLoop(reader, reader.Length())
	ingameAudioPlayer, err = ctx.NewPlayer(infiniteReader)
	if err != nil {
		log.Fatal(err)
	}
	ingameAudioPlayer.SetVolume(defaultMusicVolume)

	reader, err = wav.Decode(ctx, bytes.NewReader(mainmenuMusicBytes))
	if err != nil {
		log.Fatal(err)
	}
	infiniteReader = audio.NewInfiniteLoop(reader, reader.Length())
	mainmenuAudioPlayer, err = ctx.NewPlayer(infiniteReader)
	if err != nil {
		log.Fatal(err)
	}
	mainmenuAudioPlayer.SetVolume(defaultMusicVolume)
}

func ReplayInGameMusic() {
	ingameAudioPlayer.Rewind()
	if !ingameAudioPlayer.IsPlaying() {
		ingameAudioPlayer.Play()
	}
}

func ResumeInGameMusic() {
	if !ingameAudioPlayer.IsPlaying() {
		ingameAudioPlayer.Play()
	}
}

func StopInGameMusic() {
	if ingameAudioPlayer.IsPlaying() {
		ingameAudioPlayer.Pause()
	}
}

func PlayMainmenuMusic() {
	if !mainmenuAudioPlayer.IsPlaying() {
		mainmenuAudioPlayer.Rewind()
		mainmenuAudioPlayer.Play()
	}
}

func StopMainmenuMusic() {
	if mainmenuAudioPlayer.IsPlaying() {
		mainmenuAudioPlayer.Pause()
	}
}

func PlayHitSound() {
	hitAudioPlayer.Rewind()
	hitAudioPlayer.Play()
}
