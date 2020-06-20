// Copyright (C) 2020 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

package model

import (
	"github.com/syncthing/lib/fs"
)

type OwnershipConfiguration struct {
	MapToFixedGlobal bool   `xml:"mapToFixedGlobal,attr" json:"mapToFixedGlobal"`
	FixGlobalUid     uint32 `xml:"fixGlobalUid,attr" json:"fixGlobalUid"`
	FixGlobalGid     uint32 `xml:"fixGlobalGid,attr" json:"fixGlobalGid"`
	// only POSIX:
	MapToFixedLocal  bool `xml:"mapToFixedLocal,attr" json:"mapToFixedLocal"`
	FixLocalUid      int  `xml:"fixLocalUid,attr" json:"fixLocalUid"`
	FixLocalGid      int  `xml:"fixLocalGid,attr" json:"fixLocalGid"`
	IdentityMapLocal bool `xml:"identityMapLocal,attr" json:"identiyMapLocal"`
}

func (oc OwnershipConfiguration) getOwnerGroup(file fs.FileInfo) (uid uint32, gid uint32) {
	if oc.MapToFixedGlobal {
		uid, gid = oc.FixGlobalUid, oc.FixGlobalGid
		return
	}

	if runtime.GOOS == "windows" {
		// Can't do anything.
		return
	}

	uid = uint32(fi.Owner())
	gid = uint32(fi.Group())
}

func (oc OwnershipConfiguration) SetOwnerhip(path string, fs fs.Filesystem, uid uint32, gid uint32) error {
	if runtime.GOOS == "windows" {
		// Can't do anything => use Ownership of syncthing caller
		return nil
	}

	if oc.MapToFixedLocal {
		return fs.Lchown(path, oc.FixLocalUid, oc.FixLocalGid)
	}

	if oc.IdentityMapLocal {
		return fs.Lchown(path, int(uid), int(gid))
	}

	// Not configured to do anything => use Ownership of syncthing caller
	return nil
}
