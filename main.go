package main
// ocfw - open connect server firewall script
// Copyright (C) 2016 Maximilian Pachl

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.


// ----------------------------------------------------------------------------------
//  imports
// ----------------------------------------------------------------------------------

import (
    "fmt"
    "log"
    "os"

    "github.com/faryon93/ocfw/ocenv"
    "github.com/faryon93/ocfw/config"

    "github.com/alexflint/go-filemutex"
)


// ----------------------------------------------------------------------------------
//  constants
// ----------------------------------------------------------------------------------

const (
    CONFIG_FILE = "/etc/ocserv/ocfw.conf"
)


// ----------------------------------------------------------------------------------
//  application entry
// ----------------------------------------------------------------------------------

func main() {
    retval := 0
    defer os.Exit(retval)

    // load the config file
    conf, err := config.Load(CONFIG_FILE)
    if err != nil {
        fmt.Println("failed to load configuration file:", err.Error())
        retval = -1
        return
    }

    // aquire the lockfile
    mutex, err := filemutex.New(conf.Common.LockFile)
    if err != nil {
        log.Println("failed to aquire lock file:", err.Error())
        retval = -1
        return
    }
    mutex.Lock()
    defer mutex.Unlock()

    // print some info
    log.Println("user", ocenv.Username, ocenv.Reason, "from", ocenv.RealIp)

    // a connect job
    if ocenv.IsConnect() {
        retval = connect()

    // this is a disconnect call
    } else if (ocenv.IsDisconnect()) {
        retval = disconnect()

    // invalid REASON
    } else {
        log.Println("failure: invalid openconnect reason:", "\"" + ocenv.Reason + "\"")
        retval = -1
    }
}