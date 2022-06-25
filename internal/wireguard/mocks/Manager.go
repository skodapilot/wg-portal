// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	io "io"

	model "github.com/h44z/wg-portal/internal/model"
	mock "github.com/stretchr/testify/mock"

	netlink "github.com/vishvananda/netlink"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// ApplyDefaultConfigs provides a mock function with given fields: id
func (_m *Manager) ApplyDefaultConfigs(id model.InterfaceIdentifier) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.InterfaceIdentifier) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateInterface provides a mock function with given fields: id
func (_m *Manager) CreateInterface(id model.InterfaceIdentifier) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.InterfaceIdentifier) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteInterface provides a mock function with given fields: id
func (_m *Manager) DeleteInterface(id model.InterfaceIdentifier) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.InterfaceIdentifier) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllUsedIPs provides a mock function with given fields: device
func (_m *Manager) GetAllUsedIPs(device model.InterfaceIdentifier) ([]*netlink.Addr, error) {
	ret := _m.Called(device)

	var r0 []*netlink.Addr
	if rf, ok := ret.Get(0).(func(model.InterfaceIdentifier) []*netlink.Addr); ok {
		r0 = rf(device)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*netlink.Addr)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.InterfaceIdentifier) error); ok {
		r1 = rf(device)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFreshIp provides a mock function with given fields: device, subnetCidr, increment
func (_m *Manager) GetFreshIp(device model.InterfaceIdentifier, subnetCidr string, increment ...bool) (*netlink.Addr, error) {
	_va := make([]interface{}, len(increment))
	for _i := range increment {
		_va[_i] = increment[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, device, subnetCidr)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *netlink.Addr
	if rf, ok := ret.Get(0).(func(model.InterfaceIdentifier, string, ...bool) *netlink.Addr); ok {
		r0 = rf(device, subnetCidr, increment...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*netlink.Addr)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.InterfaceIdentifier, string, ...bool) error); ok {
		r1 = rf(device, subnetCidr, increment...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFreshKeypair provides a mock function with given fields:
func (_m *Manager) GetFreshKeypair() (model.KeyPair, error) {
	ret := _m.Called()

	var r0 model.KeyPair
	if rf, ok := ret.Get(0).(func() model.KeyPair); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.KeyPair)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetImportableInterfaces provides a mock function with given fields:
func (_m *Manager) GetImportableInterfaces() ([]*model.ImportableInterface, error) {
	ret := _m.Called()

	var r0 []*model.ImportableInterface
	if rf, ok := ret.Get(0).(func() []*model.ImportableInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ImportableInterface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInterface provides a mock function with given fields: id
func (_m *Manager) GetInterface(id model.InterfaceIdentifier) (*model.Interface, error) {
	ret := _m.Called(id)

	var r0 *model.Interface
	if rf, ok := ret.Get(0).(func(model.InterfaceIdentifier) *model.Interface); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.InterfaceIdentifier) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInterfaceConfig provides a mock function with given fields: cfg, peers
func (_m *Manager) GetInterfaceConfig(cfg *model.Interface, peers []*model.Peer) (io.Reader, error) {
	ret := _m.Called(cfg, peers)

	var r0 io.Reader
	if rf, ok := ret.Get(0).(func(*model.Interface, []*model.Peer) io.Reader); ok {
		r0 = rf(cfg, peers)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Reader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Interface, []*model.Peer) error); ok {
		r1 = rf(cfg, peers)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInterfaces provides a mock function with given fields:
func (_m *Manager) GetInterfaces() ([]*model.Interface, error) {
	ret := _m.Called()

	var r0 []*model.Interface
	if rf, ok := ret.Get(0).(func() []*model.Interface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPeerConfig provides a mock function with given fields: peer
func (_m *Manager) GetPeerConfig(peer *model.Peer) (io.Reader, error) {
	ret := _m.Called(peer)

	var r0 io.Reader
	if rf, ok := ret.Get(0).(func(*model.Peer) io.Reader); ok {
		r0 = rf(peer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Reader)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.Peer) error); ok {
		r1 = rf(peer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPeers provides a mock function with given fields: device
func (_m *Manager) GetPeers(device model.InterfaceIdentifier) ([]*model.Peer, error) {
	ret := _m.Called(device)

	var r0 []*model.Peer
	if rf, ok := ret.Get(0).(func(model.InterfaceIdentifier) []*model.Peer); ok {
		r0 = rf(device)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Peer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.InterfaceIdentifier) error); ok {
		r1 = rf(device)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPreSharedKey provides a mock function with given fields:
func (_m *Manager) GetPreSharedKey() (model.PreSharedKey, error) {
	ret := _m.Called()

	var r0 model.PreSharedKey
	if rf, ok := ret.Get(0).(func() model.PreSharedKey); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.PreSharedKey)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsedIPs provides a mock function with given fields: device, subnetCidr
func (_m *Manager) GetUsedIPs(device model.InterfaceIdentifier, subnetCidr string) ([]*netlink.Addr, error) {
	ret := _m.Called(device, subnetCidr)

	var r0 []*netlink.Addr
	if rf, ok := ret.Get(0).(func(model.InterfaceIdentifier, string) []*netlink.Addr); ok {
		r0 = rf(device, subnetCidr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*netlink.Addr)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.InterfaceIdentifier, string) error); ok {
		r1 = rf(device, subnetCidr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportInterface provides a mock function with given fields: cfg
func (_m *Manager) ImportInterface(cfg *model.ImportableInterface) error {
	ret := _m.Called(cfg)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.ImportableInterface) error); ok {
		r0 = rf(cfg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemovePeer provides a mock function with given fields: peer
func (_m *Manager) RemovePeer(peer model.PeerIdentifier) error {
	ret := _m.Called(peer)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.PeerIdentifier) error); ok {
		r0 = rf(peer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SavePeers provides a mock function with given fields: peers
func (_m *Manager) SavePeers(peers ...*model.Peer) error {
	_va := make([]interface{}, len(peers))
	for _i := range peers {
		_va[_i] = peers[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...*model.Peer) error); ok {
		r0 = rf(peers...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateInterface provides a mock function with given fields: cfg
func (_m *Manager) UpdateInterface(cfg *model.Interface) error {
	ret := _m.Called(cfg)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Interface) error); ok {
		r0 = rf(cfg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}