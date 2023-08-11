package message

type VgDw200In struct {
	CmdHeader [2]byte
	Cmd       byte
	Flag      byte
	DataLen   [2]byte
	Data      []byte
	Checksum  byte
}

// Parser 解析应答报文
func Parser(b []byte) *VgDw200In {
	lastIndex := len(b) - 1

	return &VgDw200In{
		CmdHeader: *(*[2]byte)(b[0:2]),
		Cmd:       b[2],
		Flag:      b[3],
		DataLen:   *(*[2]byte)(b[4:6]),
		Data:      b[6:lastIndex],
		Checksum:  b[lastIndex],
	}
}
