package adapter

import "fmt"

type OldPlayer interface {
	PlayMp3(string) string
	PlayFlac(string) string
	PlayWav(string) string
}

type OldPlayerImpl struct{}

func (o *OldPlayerImpl) PlayMp3(filename string) string {
	return fmt.Sprintf("PLAYING%sMP3", filename)
}

func (o *OldPlayerImpl) PlayFlac(filename string) string {
	return fmt.Sprintf("PLAYING%sFLAC", filename)
}

func (o *OldPlayerImpl) PlayWav(filename string) string {
	return fmt.Sprintf("PLAYING%sWAV", filename)
}

type MusicType uint8

const (
	MP3 MusicType = iota
	FLAC
	WAV
)

type NewPlayer interface {
	Play(MusicType, string) string
}

type NewPlayerImpl struct {
	o OldPlayer
}

func (n *NewPlayerImpl) Play(musicType MusicType, filename string) string {
	switch musicType {
	case MP3:
		return n.o.PlayMp3(filename)
	case FLAC:
		return n.o.PlayFlac(filename)
	case WAV:
		return n.o.PlayWav(filename)
	default:
		return ""
	}
}
