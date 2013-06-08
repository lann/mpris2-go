package mpris2

import (
	"fmt"
	"strings"
)

// http://specifications.freedesktop.org/mpris-spec/latest/

const (
	mprisPath          = "/org/mpris/MediaPlayer2"
	mprisPrefix       = "org.mpris.MediaPlayer2"
	rootInterface      = mprisPrefix
	playerInterface    = mprisPrefix + ".Player"
	trackListInterface = mprisPrefix + ".TrackList"
	playListsInterface = mprisPrefix + ".PlayLists"
)

type MediaPlayer struct {
	root   *Object
	player *Object
}

func (conn *Conn) GetMediaPlayer(objectName string) *MediaPlayer {
	return &MediaPlayer{
		root:   conn.getObject(objectName, mprisPath, rootInterface),
		player: conn.getObject(objectName, mprisPath, playerInterface),
	}
}

func (conn *Conn) ListMediaPlayers() (names []string, err error) {
	allNames, err := conn.ListNames()
	if err != nil {
		return
	}

	for _, name := range allNames {
		if strings.HasPrefix(name, mprisPrefix) {
			names = append(names, name)
		}
	}
	return
}

func (conn *Conn) GetFirstMediaPlayer() (*MediaPlayer, error) {
	names, err := conn.ListMediaPlayers()
	if err != nil {
		return nil, err
	}

	if len(names) == 0 {
		return nil, fmt.Errorf("No media players found")
	}
	
	return conn.GetMediaPlayer(names[0]), nil
}

// root methods

func (mp *MediaPlayer) Raise() error {
	return mp.root.callVoid("Raise")
}

func (mp *MediaPlayer) Quit() error {
	return mp.root.callVoid("Quit")
}

func (mp *MediaPlayer) CanRaise() (bool, error) {
	return mp.root.getBool("CanRaise")
}

func (mp *MediaPlayer) CanQuit() (bool, error) {
	return mp.root.getBool("CanQuit")
}

func (mp *MediaPlayer) Identity() (string, error) {
	return mp.root.getString("Identity")
}

func (mp *MediaPlayer) SupportedUriSchemes() ([]string, error) {
	return mp.root.getStringArray("SupportedUriSchemes")
}

func (mp *MediaPlayer) SupportedMimeTypes() ([]string, error) {
	return mp.root.getStringArray("SupportedMimeTypes")
}

// player methods

func (mp *MediaPlayer) Next() error {
	return mp.player.callVoid("Next")
}

func (mp *MediaPlayer) Previous() error {
	return mp.player.callVoid("Previous")
}

func (mp *MediaPlayer) Pause() error {
	return mp.player.callVoid("Pause")
}

func (mp *MediaPlayer) PlayPause() error {
	return mp.player.callVoid("PlayPause")
}

func (mp *MediaPlayer) Stop() error {
	return mp.player.callVoid("Stop")
}

func (mp *MediaPlayer) Play() error {
	return mp.player.callVoid("Play")
}

func (mp *MediaPlayer) OpenUri(uri string) error {
	return mp.player.callVoid("OpenUri", uri)
}

func (mp *MediaPlayer) PlaybackStatus() (string, error) {
	return mp.player.getString("PlaybackStatus")
}

func (mp *MediaPlayer) CanGoNext() (bool, error) {
	return mp.player.getBool("CanGoNext")
}

func (mp *MediaPlayer) CanGoPrevious() (bool, error) {
	return mp.player.getBool("CanGoPrevious")
}

func (mp *MediaPlayer) CanPlay() (bool, error) {
	return mp.player.getBool("CanPlay")
}

func (mp *MediaPlayer) CanPause() (bool, error) {
	return mp.player.getBool("CanPause")
}

func (mp *MediaPlayer) CanControl() (bool, error) {
	return mp.player.getBool("CanControl")
}

type Metadata map[string]interface{}

func (mp *MediaPlayer) Metadata() (Metadata, error) {
	return mp.player.getStringMap("Metadata")
}

func (data Metadata) ArtUrl() string {
	val, _ := data["mpris:artUrl"].(string)
	return val
}

func (data Metadata) Length() uint64 {
	val, _ := data["mpris:length"].(uint64)
	return val
}

func (data Metadata) TrackId() string {
	val, _ := data["mpris:trackid"].(string)
	return val
}

func (data Metadata) Album() string {
	val, _ := data["xesam:album"].(string)
	return val
}

func (data Metadata) Artists() []string {
	val, _ := data["xesam:artist"].([]string)
	return val
}

func (data Metadata) DiscNumber() int32 {
	val, _ := data["xesam:discNumber"].(int32)
	return val
}

func (data Metadata) Title() string {
	val, _ := data["xesam:title"].(string)
	return val
}

func (data Metadata) TrackNumber() int32 {
	val, _ := data["xesam:trackNumber"].(int32)
	return val
}

func (data Metadata) Url() string {
	val, _ := data["xesam:url"].(string)
	return val
}
