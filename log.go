// Copyright 2017 Baliance. All rights reserved.
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package unioffice

import (
	"fmt"
	"log"
	"os"
)

// Log is used to log content from within the library.  The intent is to use
// logging sparingly, preferring to return an error.  At the very least this
// allows redirecting logs to somewhere more appropriate than stdout.
var Log = log.Printf

// DisableLogging sets the Log function to a no-op so that any log messages are
// silently discarded.
func DisableLogging() {
	Log = func(string, ...interface{}) {}
}

// Enable log to write to file
func EnableLogWritetoFile(path string) {
	Log = func(format string, a ...interface{}) {
		f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err.Error())
		}
		if _, err := f.Write([]byte(fmt.Sprintf(format, a))); err != nil {
			log.Fatal(err.Error())
		}
		if err := f.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}
}
