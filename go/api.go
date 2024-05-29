package capyskygo

import (
	"github.com/davidforest123/goutil/net/grpcs"
)

type (
	CapyskyClient struct {
		network   string
		addr      string
		accessKey string
		rpc       *grpcs.Client
	}
)

func New(network, addr, accessKey string) (*CapyskyClient, error) {
	cli := &CapyskyClient{
		network:   network,
		addr:      addr,
		accessKey: accessKey,
	}

	err := error(nil)
	cli.rpc, err = grpcs.Dial(grpcs.RpcTypeJSON, network, addr, NewRPCChecker())
	if err != nil {
		return nil, err
	}

	return cli, nil
}

// Close connection to server.
func (c *CapyskyClient) Close() error {
	return c.rpc.Close()
}
