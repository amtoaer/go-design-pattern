package adapter

import "testing"

func TestAdapter(t *testing.T) {
	var (
		oldPlayer OldPlayer = &OldPlayerImpl{}
		newPlayer NewPlayer = &NewPlayerImpl{o: oldPlayer}
	)
	testCases := []struct {
		musicType MusicType
		function  func(string) string
	}{
		{FLAC, oldPlayer.PlayFlac},
		{WAV, oldPlayer.PlayWav},
		{MP3, oldPlayer.PlayMp3},
	}
	for _, testCase := range testCases {
		if newPlayer.Play(testCase.musicType, "filename") != testCase.function("filename") {
			t.Error("适配者模式出现错误")
		}
	}
}
