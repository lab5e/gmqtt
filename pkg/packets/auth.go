package packets

import (
	"bytes"
	"fmt"
	"io"

	"github.com/lab5e/lmqtt/pkg/codes"
)

// Auth is the AUTH packet (https://docs.oasis-open.org/mqtt/mqtt/v5.0/os/mqtt-v5.0-os.html#_Toc3901217)
type Auth struct {
	FixHeader  *FixHeader
	Code       byte
	Properties *Properties
}

// String is the stringer method for the Auth type
func (a *Auth) String() string {
	return fmt.Sprintf("Auth, Code: %v, Properties: %s", a.Code, a.Properties)
}

// Pack encodes the Auth type to an io.Writer
func (a *Auth) Pack(w io.Writer) error {
	a.FixHeader = &FixHeader{PacketType: AUTH, Flags: FlagReserved}
	bufw := &bytes.Buffer{}
	if a.Code != codes.Success || a.Properties != nil {
		bufw.WriteByte(a.Code)
		a.Properties.Pack(bufw, AUTH)
	}
	a.FixHeader.RemainLength = bufw.Len()
	err := a.FixHeader.Pack(w)
	if err != nil {
		return err
	}
	_, err = bufw.WriteTo(w)
	return err
}

// Unpack decodes an Auth type from an io.Reader
func (a *Auth) Unpack(r io.Reader) error {
	if a.FixHeader.RemainLength == 0 {
		a.Code = codes.Success
		return nil
	}
	restBuffer := make([]byte, a.FixHeader.RemainLength)
	_, err := io.ReadFull(r, restBuffer)
	if err != nil {
		return codes.ErrMalformed
	}
	bufr := bytes.NewBuffer(restBuffer)
	a.Code, err = bufr.ReadByte()
	if err != nil {
		return codes.ErrMalformed
	}
	if !ValidateCode(AUTH, a.Code) {
		return codes.ErrProtocol
	}
	a.Properties = &Properties{}
	return a.Properties.Unpack(bufr, AUTH)
}

// NewAuthPacket creates a new auth packet
func NewAuthPacket(fh *FixHeader, r io.Reader) (*Auth, error) {
	p := &Auth{FixHeader: fh}
	//Determine whether the flags flags are legal [MQTT-2.2.2-2]
	if fh.Flags != FlagReserved {
		return nil, codes.ErrMalformed
	}
	err := p.Unpack(r)
	if err != nil {
		return nil, err
	}
	return p, err
}
