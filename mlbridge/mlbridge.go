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
	"st