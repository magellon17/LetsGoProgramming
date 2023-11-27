package wifi_test

import (
	"errors"
	myWifi "example_mock/internal/wifi"
	"net"
	"testing"

	wifi "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

type wifiInfo struct {
	addr string
	name string
}

type rowTestSysInfo struct {
	wifiInfos   []wifiInfo
	errExpected error
}

var testTable = []rowTestSysInfo{
	{
		wifiInfos: []wifiInfo{
			{addr: "00:11:22:33:44:55", name: "my"},
			{addr: "aa:bb:cc:dd:ee:ff", name: "your"},
		},
	},
	{
		errExpected: errors.New("ExpectedError"),
	},
}

func TestGetAdresses(t *testing.T) {
	for i, row := range testTable {
		mockWifi := NewWiFi(t)
		wifiService := myWifi.Service{WiFi: mockWifi}

		mockWifi.On("Interfaces").Return(mockIfaces(row.wifiInfos), row.errExpected)
		actualAddrs, err := wifiService.GetAddresses()
		expectedMacs := parseMACs(row.wifiInfos)

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, expectedMacs, actualAddrs, "row: %d, expected addrs: %s, actual addrs: %s", i, expectedMacs, actualAddrs)
	}
}

func TestGetNames(t *testing.T) {
	for i, testRow := range testTable {
		mockWifi := NewWiFi(t)
		wifiService := myWifi.Service{WiFi: mockWifi}

		mockWifi.On("Interfaces").Return(mockIfaces(testRow.wifiInfos), testRow.errExpected)
		actualNames, err := wifiService.GetNames()
		expectedNames := parseNames(testRow.wifiInfos)

		if testRow.errExpected != nil {
			require.ErrorIs(t, err, testRow.errExpected, "row: %d, expected error: %w, actual error: %w", i, testRow.errExpected, err)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, expectedNames, actualNames, "row: %d, expected names: %s, actual names: %s", i, expectedNames, actualNames)
	}
}

func mockIfaces(wifiInfos []wifiInfo) []*wifi.Interface {
	interfaces := make([]*wifi.Interface, len(wifiInfos))
	for i, wifiInfo := range wifiInfos {
		hwAddr := parseMAC(wifiInfo.addr)

		if hwAddr == nil {
			continue
		}

		iface := &wifi.Interface{
			Index:        i + 1,
			Name:         wifiInfo.name,
			HardwareAddr: hwAddr,
			PHY:          1,
			Device:       1,
			Type:         wifi.InterfaceTypeAPVLAN,
			Frequency:    0,
		}
		interfaces[i] = iface
	}

	return interfaces
}

func parseNames(wifiInfos []wifiInfo) []string {
	names := make([]string, len(wifiInfos))

	for i, wifi := range wifiInfos {
		names[i] = wifi.name
	}

	return names
}

func parseMACs(wifiInfos []wifiInfo) []net.HardwareAddr {
	addrs := make([]net.HardwareAddr, len(wifiInfos))

	for i, wifiInfo := range wifiInfos {
		addrs[i] = parseMAC(wifiInfo.addr)
	}

	return addrs
}
func parseMAC(macStr string) net.HardwareAddr {
	hwAddr, err := net.ParseMAC(macStr)

	if err != nil {
		return nil
	}

	return hwAddr
}
