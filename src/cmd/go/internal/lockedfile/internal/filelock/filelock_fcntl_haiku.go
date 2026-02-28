// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build haiku

package filelock

type inode = int64 // type of syscall.Stat_t.Ino

