package core

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/elcarta/evilginx2/log"
)

const (
	BLACKLIST_MODE_FULL   = 0
	BLACKLIST_MODE_UNAUTH = 1
	BLACKLIST_MODE_OFF    = 2
)

type BlockIP struct {
	ipv4 net.IP
	mask *net.IPNet
}

type Blacklist struct {
	ips        map[string]*BlockIP
	masks      []*BlockIP
	configPath string
	mode       int
}

func NewBlacklist(path string) (*Blacklist, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	bl := &Blacklist{
		ips:        make(map[string]*BlockIP),
		configPath: path,
		mode:       BLACKLIST_MODE_OFF,
	}

	fs := bufio.NewScanner(f)
	fs.Split(bufio.ScanLines)
