package libspf2

import (
	"net"
	"testing"
)

func TestIPv4(t *testing.T) {
	var err error

	s := NewServer()
	defer s.Free()

	req := NewRequest(s)
	defer req.Free()

	err = req.SetIPv4Addr("173.194.39.150")
	if err != nil {
		t.Fatal(err)
	}

	err = req.SetEnvFrom("gmail.com")
	if err != nil {
		t.Fatal(err)
	}
	resp, err := req.Query()
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Free()

	res := resp.Result()
	if res != SPFResultPASS {
		t.Fatal(res)
	}
	if s := res.String(); s != "pass" {
		t.Fatal(s)
	}
}

func TestIPv6(t *testing.T) {
	var err error
	var ip net.IP

	s := NewServer()
	defer s.Free()

	req := NewRequest(s)
	defer req.Free()

	ip = net.ParseIP("2404:6800:4003:803::1006")
	err = req.SetIPAddr(ip)
	if err != nil {
		t.Fatal(err)
	}

	err = req.SetEnvFrom("gmail.com")
	if err != nil {
		t.Fatal(err)
	}
	resp, err := req.Query()
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Free()

	res := resp.Result()
	if res != SPFResultPASS {
		t.Fatal(res)
	}
	if s := res.String(); s != "pass" {
		t.Fatal(s)
	}
}
