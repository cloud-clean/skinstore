package message

import "encoding/xml"

const(
	MSG_TEXT = "text"
	MSG_IMAGE = "image"
	MSG_LOCATION = "location"
)

type WxMsg struct {
	XMLName 	xml.Name `xml:"xml"`
	ToUserName string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime uint64 `xml:"CreateTime"`
	MsgType string `xml:"MsgType"`
	MsgId uint64 `xml:"MsgId"`
	Content string `xml:"Content,omitempty"`
	PicUrl string `xml:"PicUrl,omitempty"`
	MediaId string `xml:"MediaId,omitempty"`
	Format string `xml:"Format,omitempty"`
	Recognition string `xml:"Recognition,omitempty"`
	ThumbMediaId string `xml:"ThumbMediaId,omitempty"`
	Location_X string `xml:"Location_X,omitempty"`
	Location_Y string `xml:"Location_Y,omitempty"`
	Scale	string 	  `xml:"Scale,omitempty"`
	Label	string 	  `xml:"Label,omitempty"`
}
