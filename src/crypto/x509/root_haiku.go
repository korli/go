// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

// Possible certificate files; stop after finding one.
var certFiles = []string{
	"/system/data/ssl/CARootCertificates.pem",
}

// Possible directories with certificate files; all will be read.
var certDirectories = []string{
	"/system/data/ssl/certs",
}
