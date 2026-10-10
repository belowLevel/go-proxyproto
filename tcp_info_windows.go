package proxyproto

import "errors"

var (
	ErrNotSupportOnWindows = errors.New("notSupportOnWindows")
)

type TCPFullInfo struct {
}

func (m *Manager) TCPFullInfoByID(id string) (*TCPFullInfo, error) {
	return nil, ErrNotSupportOnWindows
}
