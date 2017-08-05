package fs

import "github.com/spf13/afero"

var OsFs afero.Fs = afero.NewOsFs()
var OsAfero = &afero.Afero{Fs: OsFs}
