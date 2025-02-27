package client

import (
	"io"

	"github.com/netbirdio/netbird/client/system"
	"github.com/netbirdio/netbird/management/proto"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type Client interface {
	io.Closer
	Sync(msgHandler func(msg *proto.SyncResponse) error) error
	GetServerPublicKey() (*wgtypes.Key, error)
	Register(serverKey wgtypes.Key, setupKey string, jwtToken string, sysInfo *system.Info, sshKey []byte) (*proto.LoginResponse, error)
	Login(serverKey wgtypes.Key, sysInfo *system.Info, sshKey []byte) (*proto.LoginResponse, error)
	GetDeviceAuthorizationFlow(serverKey wgtypes.Key) (*proto.DeviceAuthorizationFlow, error)
	GetPKCEAuthorizationFlow(serverKey wgtypes.Key) (*proto.PKCEAuthorizationFlow, error)
	GetNetworkMap() (*proto.NetworkMap, error)
}
