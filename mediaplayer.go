package mpris2

import (
	"fmt"
	"strings"
)

const (
	mprisPath          = "/org/mpris/MediaPlayer2"
	mprisPrefix       = "org.mpris.MediaPlayer2"
	rootInterface      = mprisPrefix
	playerInterface    = mprisPrefix + ".Player"
	trackListInterface = mprisPrefix + ".TrackList"
	playListsInterface = mprisPrefix + ".PlayLists"
)

type MediaPlayer struct {
	conn   *Connection
	root   *Interface
	player *Interface
}

func (conn *Connection) GetMediaPlayer(objectName string) (mp *MediaPlayer, err error) {
	obj, err := conn.getObject(objectName, mprisPath)
	if err != nil {
		return
	}

	root, err := obj.getInterface(rootInterface)
	if err != nil {
		return
	}
	
	player, err := obj.getInterface(playerInterface)
	if err != nil {
		return
	}
	
	mp = &MediaPlayer{conn: conn, root: root, player: player}
	return
}

func (conn *Connection) ListMediaPlayers() (names []string, err error) {
	obj, err := conn.getObject(dbusObject, dbusPath)
	if err != nil {
		return
	}

	iface, err := obj.getInterface(dbusInterface)
	if err != nil {
		return
	}

	out, err := iface.call(listMethod)
	if err != nil {
		return
	}

	items, ok := out[0].([]interface{})
	if !ok {
		err = fmt.Errorf("unexpected return value %v", items)
		return
	}

	for _, item := range items {
		name, ok := item.(string)
		if !ok {
			err = fmt.Errorf("unexpected return value %v", item)
			return
		}
		
		if strings.HasPrefix(name, mprisPrefix) {
			names = append(names, name)
		}
	}
	return
}
func (conn *Connection) GetFirstMediaPlayer() (*MediaPlayer, error) {
	names, err := conn.ListMediaPlayers()
	if err != nil {
		return nil, err
	}

	if len(names) == 0 {
		return nil, fmt.Errorf("No media players found")
	}
	
	return conn.GetMediaPlayer(names[0])
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

func (mp *MediaPlayer) Metadata() ([]interface{}, error) {
	return mp.player.getProp("Metadata")
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
