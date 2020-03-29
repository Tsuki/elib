package utils

import "encoding/xml"

type (
	Options struct {
		Debug      bool
		InputPath  string
		BackupPath string
		Worker     uint
	}
	Movie struct {
		XMLName  xml.Name `xml:"movie"`
		Text     string   `xml:",chardata"`
		Title    string   `xml:"title"`
		Set      string   `xml:"set"`
		Studio   string   `xml:"studio"`
		Year     string   `xml:"year"`
		Outline  string   `xml:"outline"`
		Plot     string   `xml:"plot"`
		Runtime  string   `xml:"runtime"`
		Director string   `xml:"director"`
		Poster   string   `xml:"poster"`
		Thumb    string   `xml:"thumb"`
		Fanart   string   `xml:"fanart"`
		Actor    []Actor  `xml:"actor"`
		Maker    string   `xml:"maker"`
		Label    string   `xml:"label"`
		Tag      []string `xml:"tag"`
		Genre    []string `xml:"genre"`
		Num      string   `xml:"num"`
		Release  string   `xml:"release"`
		Cover    string   `xml:"cover"`
		Website  string   `xml:"website"`
	}
	Actor struct {
		Text  string `xml:",chardata"`
		Name  string `xml:"name"`
		Thumb string `xml:"thumb"`
	}
)

var (
	OptionsDefault = map[string]interface{}{
		"debug": true,
	}
	Option = Options{}
)
