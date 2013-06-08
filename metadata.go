package mpris2

type Metadata map[string]interface{}

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
