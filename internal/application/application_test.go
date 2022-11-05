package application

import (
	"strings"
	"testing"
)

var (
	configText = `
# config.yaml

logger:
  level: "debug"
rest-port: "80"
grpc-port: "443"
`

	cfgExpected = Config{
		Logger: Logger{
			Level: "debug",
		},
		RestPort: "80",
		GrpcPorg: "443",
	}
)

func TestReadConfig(t *testing.T) {

	cfg := readConfigFile(strings.NewReader(configText))

	if cfgExpected != *cfg {
		t.Errorf("error reading config: expected %v, got %v", cfgExpected, *cfg)
	}
}
