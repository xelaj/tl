package somepkg

import (
	"github.com/xelaj/tl"
)

type Updates interface {
	tl.Object
	_Updates()
}

func newUpdatesByCRC(crc uint32) Updates {
	switch crc {
	case (*UpdateShort)(nil).CRC():
		return &UpdateShort{}
	case (*UpdateShortChatMessage)(nil).CRC():
		return &UpdateShortChatMessage{}
	case (*UpdateShortMessage)(nil).CRC():
		return &UpdateShortMessage{}
	case (*UpdateShortSentMessage)(nil).CRC():
		return &UpdateShortSentMessage{}
	case (*UpdatesObj)(nil).CRC():
		return &UpdatesObj{}
	case (*UpdatesCombined)(nil).CRC():
		return &UpdatesCombined{}
	case uint32(UpdatesTooLong),
		uint32(UpdatesTooShort):
		return UpdatesEnum(crc)
	default:
		return nil
	}
}

type UpdateShort struct {
	Update string
	Date   int32
}

func (*UpdateShort) CRC() uint32 { return 0x78d4dec1 }
func (*UpdateShort) _Updates()   {}

type UpdateShortChatMessage struct {
	Out         bool `tl:"flag:1,encoded_in_bitflags"`
	Mentioned   bool `tl:"flag:4,encoded_in_bitflags"`
	MediaUnread bool `tl:"flag:5,encoded_in_bitflags"`
	Silent      bool `tl:"flag:13,encoded_in_bitflags"`
	ID          int32
	FromID      int32
	ChatID      int32
	Message     string
	Pts         int32
	PtsCount    int32
	Date        int32
}

func (*UpdateShortChatMessage) CRC() uint32 { return 0x402d5dbb }
func (*UpdateShortChatMessage) _Updates()   {}

type UpdateShortMessage struct {
	Out         bool `tl:"flag:1,encoded_in_bitflags"`
	Mentioned   bool `tl:"flag:4,encoded_in_bitflags"`
	MediaUnread bool `tl:"flag:5,encoded_in_bitflags"`
	Silent      bool `tl:"flag:13,encoded_in_bitflags"`
	ID          int32
	UserID      int32
	Message     string
	Pts         int32
	PtsCount    int32
	Date        int32
}

func (*UpdateShortMessage) CRC() uint32 { return 0x2296d2c8 }
func (*UpdateShortMessage) _Updates()   {}

type UpdateShortSentMessage struct {
	Out      bool `tl:"flag:1,encoded_in_bitflags"`
	ID       int32
	Pts      int32
	PtsCount int32
	Date     int32
}

func (*UpdateShortSentMessage) CRC() uint32 { return 0x11f1331c }
func (*UpdateShortSentMessage) _Updates()   {}

type UpdatesObj struct {
	Date int32
	Seq  int32
}

func (*UpdatesObj) CRC() uint32 { return 0x74ae4240 }
func (*UpdatesObj) _Updates()   {}

type UpdatesCombined struct {
	Date     int32
	SeqStart int32
	Seq      int32
}

func (*UpdatesCombined) CRC() uint32 { return 0x725b04c3 }
func (*UpdatesCombined) _Updates()   {}

type UpdatesEnum uint32

const (
	UpdatesTooLong  UpdatesEnum = 0xe317af7e
	UpdatesTooShort UpdatesEnum = 0x12345678
)

func (e UpdatesEnum) CRC() uint32 { return uint32(e) }
func (_ UpdatesEnum) _Updates()   {}
