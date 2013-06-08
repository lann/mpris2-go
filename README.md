mpris2-go
=========

[MPRIS2](http://specifications.freedesktop.org/mpris-spec/latest/) client package for Go

Docs: http://godoc.org/github.com/lann/mpris2-go

Example
-------

```golang
import "github.com/lann/mpris2"

conn, err := mpris2.Connect() // Connect to DBus

mp, err := conn.GetAnyMediaPlayer()

err = mp.Play()

meta, err := mp.Metadata()
fmt.Println("Title: ", meta.Title())
```

Also, check out [flipperdinger](https://github.com/lann/flipperdinger), a CLI remote.
