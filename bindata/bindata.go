package bindata

import (
	"github.com/jteeuwen/go-bindata"
)

// go-bindata is vendored for the build process, but dep requires
// that it be imported somewhere.
func dummy() {
	_ = bindata.NewConfig()
}
