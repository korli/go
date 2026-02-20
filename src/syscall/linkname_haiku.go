// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

import _ "unsafe"

// used by os
//go:linkname Closedir
//go:linkname Readdir_r

// used by internal/poll
//go:linkname Fdopendir
