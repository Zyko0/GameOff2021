package assets

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

const (
	defaultSFXVolume       = 0.2
	defaultGameMusicVolume = 0.4
	defaultMainMenuVolume  = 0.8
)

var (
	ctx = audio.NewContext(44100)

	//go:embed sound/mainmenu.wav
	mainmenuMusicBytes  []byte
	mainmenuAudioPlayer *audio.Player
	//go:embed sound/gameover.wav
	gameoverMusicBytes  []byte
	gameoverAudioPlayer *audio.Player
	//go:embed sound/ingame.wav
	ingameMusicBytes  []byte
	ingameAudioPlayer *audio.Player
	//go:embed sound/hit.wav
	hitSoundBytes  []byte
	hitAudioPlayer *audio.Player
	//go:embed sound/heart.wav
	heartSoundBytes  []byte
	heartAudioPlayer *audio.Player

	defaultSFXManager = &sfxManagerImpl{}
	noopSFXManager    = &noopSFXManagerImpl{}
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

	reader, err = wav.Decode(ctx, bytes.NewReader(heartSoundBytes))
	if err != nil {
		log.Fatal(err)
	}
	heartAudioPlayer, err = ctx.NewPlayer(reader)
	if err != nil {
		log.Fatal(err)
	}
	heartAudioPlayer.SetVolume(defaultSFXVolume)

	reader, err = wav.Decode(ctx, bytes.NewReader(ingameMusicBytes))
	if err != nil {
		log.Fatal(err)
	}
	infiniteReader := audio.NewInfiniteLoop(reader, reader.Length())
	ingameAudioPlayer, err = ctx.NewPlayer(infiniteReader)
	if err != nil {
		log.Fatal(err)
	}
	ingameAudioPlayer.SetVolume(defaultGameMusicVolume)

	reader, err = wav.Decode(ctx, bytes.NewReader(mainmenuMusicBytes))
	if err != nil {
		log.Fatal(err)
	}
	infiniteReader = audio.NewInfiniteLoopWithIntro(reader, reader.Length()/5, reader.Length())
	mainmenuAudioPlayer, err = ctx.NewPlayer(infiniteReader)
	if err != nil {
		log.Fatal(err)
	}
	mainmenuAudioPlayer.SetVolume(defaultMainMenuVolume)

	reader, err = wav.Decode(ctx, bytes.NewReader(gameoverMusicBytes))
	if err != nil {
		log.Fatal(err)
	}
	infiniteReader = audio.NewInfiniteLoop(reader, reader.Length())
	gameoverAudioPlayer, err = ctx.NewPlayer(infiniteReader)
	if err != nil {
		log.Fatal(err)
	}
	gameoverAudioPlayer.SetVolume(defaultGameMusicVolume)
}

type SFXManager interface {
	PlayHitSound()
	PlayHeartSound()
}

type noopSFXManagerImpl struct{}

func (n *noopSFXManagerImpl) PlayHitSound() {}

func (n *noopSFXManagerImpl) PlayHeartSound() {}

func NoopSFXManager() SFXManager {
	return noopSFXManager
}

type sfxManagerImpl struct{}

func (s *sfxManagerImpl) PlayHitSound() {
	hitAudioPlayer.Rewind()
	hitAudioPlayer.Play()
}

func (s *sfxManagerImpl) PlayHeartSound() {
	heartAudioPlayer.Rewind()
	heartAudioPlayer.Play()
}

func DefaultSFXManager() SFXManager {
	return defaultSFXManager
}

// Global music options

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

func PlayGameoverMusic() {
	if !gameoverAudioPlayer.IsPlaying() {
		gameoverAudioPlayer.Rewind()
		gameoverAudioPlayer.Play()
	}
}

func StopGameoverMusic() {
	if gameoverAudioPlayer.IsPlaying() {
		gameoverAudioPlayer.Pause()
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
