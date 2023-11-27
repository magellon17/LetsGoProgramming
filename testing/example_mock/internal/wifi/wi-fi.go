package wifi

import (
	"net"

	"github.com/mdlayher/wifi"
	"github.com/pkg/errors"
)

type WiFi interface {
	Interfaces() ([]*wifi.Interface, error)
}

type Service struct {
	WiFi WiFi
}

func New(wifi WiFi) Service {
	return Service{WiFi: wifi}
}

func (service Service) GetAddresses() ([]net.HardwareAddr, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, errors.Wrap(err, err.Error())
	}

	addrs := make([]net.HardwareAddr, len(interfaces))

	for i, iface := range interfaces {
		addrs[i] = iface.HardwareAddr
	}

	return addrs, nil
}

func (service Service) GetNames() ([]string, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, errors.Wrap(err, err.Error())
	}

	names := make([]string, len(interfaces))

	for i, iface := range interfaces {
		names[i] = iface.Name
	}

	return names, nil
}
