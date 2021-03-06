// Copyright 2018 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main contains a plugin which provides an sqlite based datastore.
package main

import (
	log "github.com/golang/glog"

	"github.com/google/fleetspeak/fleetspeak/src/server/db"
	"github.com/google/fleetspeak/fleetspeak/src/server/sqlite"
)

// Factory is a plugins.DatabaseFactory which returns an sqlite based
// database. Its configuration string should be a filename to open or create as
// an sqlite3 database.
func Factory(file string) (db.Store, error) {
	return sqlite.MakeDatastore(file)
}

func main() {
	log.Exitf("unimplemented")
}
