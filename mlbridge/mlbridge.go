// Package mlbridge implements a plugin
package mlbridge

import (
	"context"
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
	"strings"
	"net/http"
	"strconv"
	"github.com/coredns/coredns/plugin"
	"github.com/coredns/