package mpris2

import (
	"fmt"
	"strings"
)

// http://specifications.freedesktop.org/mpris-spec/latest/

const (
	mprisPath          = "/org/mpris/MediaPlayer2"
	mprisPrefix        = "org.mpris.MediaPlayer2"
	rootInterface      = mprisPrefix
	playerInterface    = mprisPrefix + ".Player"
	trackListInterface = mprisPrefix + ".TrackList"
	playListsInterface = mprisPrefix + ".PlayLists"
)

type MediaPlayer struct {
	root   *object
	player *object
}

func (conn *Conn) GetMediaPlayer(objectName string) *MediaPlayer {
	if !strings.ContainsRune(objectName, '.') {
		objectName = fmt.Sprint(mprisPrefix, ".", objectName)
	}

	return &MediaPlayer{
		root:   conn.getObject(objectName, mprisPath, rootInterface),
		player: conn.getObject(objectName, mprisPath, playerInterface),
	}
}

func (conn *Conn) ListMediaPlayers() (names []string, err error) {
	allNames, err := conn.listNames()
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

func (conn *Conn) GetAnyMediaPlayer() (*MediaPlayer, error) {
	names, err := conn.ListMediaPlayers()
	if err != nil {
		return nil, err
	}

	if len(names) == 0 {
		return nil, fmt.Errorf("No media players found")
	}

	return conn.GetMediaPlayer(names[0]), nil
}

// root

func (mp *MediaPlayer) Raise() error                  { return mp.root.callVoid("Raise") }
func (mp *MediaPlayer) Quit() error                   { return mp.root.callVoid("Quit") }
func (mp *MediaPlayer) CanRaise() (bool, error)       { return mp.root.getBool("CanRaise") }
func (mp *MediaPlayer) CanQuit() (bool, error)        { return mp.root.getBool("CanQuit") }
func (mp *MediaPlayer) Identity() (string, error)     { return mp.root.getString("Identity") }
func (mp *MediaPlayer) DesktopEntry() (string, error) { return mp.root.getString("DesktopEntry") }

func (mp *MediaPlayer) SupportedUriSchemes() ([]string, error) {
	return mp.root.getStringArray("SupportedUriSchemes")
}

func (mp *MediaPlayer) SupportedMimeTypes() ([]string, error) {
	return mp.root.getStringArray("SupportedMimeTypes")
}

// player

func (mp *MediaPlayer) Next() error                     { return mp.player.callVoid("Next") }
func (mp *MediaPlayer) Previous() error                 { return mp.player.callVoid("Previous") }
func (mp *MediaPlayer) Pause() error                    { return mp.player.callVoid("Pause") }
func (mp *MediaPlayer) PlayPause() error                { return mp.player.callVoid("PlayPause") }
func (mp *MediaPlayer) Stop() error                     { return mp.player.callVoid("Stop") }
func (mp *MediaPlayer) Play() error                     { return mp.player.callVoid("Play") }
func (mp *MediaPlayer) Seek(offset int64) error         { return mp.player.callVoid("Seek", offset) }
func (mp *MediaPlayer) OpenUri(uri string) error        { return mp.player.callVoid("OpenUri", uri) }
func (mp *MediaPlayer) PlaybackStatus() (string, error) { return mp.player.getString("PlaybackStatus") }
func (mp *MediaPlayer) Metadata() (Metadata, error)     { return mp.player.getStringMap("Metadata") }
func (mp *MediaPlayer) Position() (int64, error)        { return mp.player.getInt64("Position") }
func (mp *MediaPlayer) CanGoNext() (bool, error)        { return mp.player.getBool("CanGoNext") }
func (mp *MediaPlayer) CanGoPrevious() (bool, error)    { return mp.player.getBool("CanGoPrevious") }
func (mp *MediaPlayer) CanPlay() (bool, error)          { return mp.player.getBool("CanPlay") }
func (mp *MediaPlayer) CanPause() (bool, error)         { return mp.player.getBool("CanPause") }
func (mp *MediaPlayer) CanControl() (bool, error)       { return mp.player.getBool("CanControl") }
