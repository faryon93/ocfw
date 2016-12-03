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
    "strings"

    "github.com/faryon93/ocfw/ocenv"
    "github.com/faryon93/ocfw/iptables"
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

    // norma ocserver connection handler
    if len(os.Args) < 2 {
        retval = connection(conf)

    // firewall setup mode
    } else {
        retval = setup(conf)
    }
}


// ----------------------------------------------------------------------------------
//  application modi
// ----------------------------------------------------------------------------------

func connection(conf *config.Config) (int) {
    // print some info aber the request
    log.Println("user", ocenv.Username, ocenv.Reason, "from", ocenv.RealIp)

    // a connect job
    if ocenv.IsConnect() {
        return connect(conf)

    // this is a disconnect call
    } else if (ocenv.IsDisconnect()) {
        return disconnect(conf)
    }

    // invalid REASON was supplied
    log.Println("failure: invalid openconnect reason:", "\"" + ocenv.Reason + "\"")
    return -1
}

func setup(conf *config.Config) (int) {
    log.Println("initializing iptables firewall")

    // flush the forward chain
    err := iptables.FlushChain("FORWARD") 
    if err != nil {
        log.Println("failed to flush the FORWARD chain:", err.Error())
        return -1
    }

    // allow established and related connections
    err = iptables.Chain("FORWARD").Append().State("ESTABLISHED", "RELATED").Accept().Apply()
    if err != nil {
        log.Println("failed to allow ESTABLISHED and RELATED connections:", err.Error())
        return -1
    }

    // drop anything else
    err = iptables.Chain("FORWARD").Append().Drop().Apply()
    if err != nil {
        log.Println("failed to drop anything else:", err.Error())
        return -1
    }

    // setup all groups as individual chains
    for name, group := range conf.Groups {
        // some metadata
        groupChain := "VPN_GROUP_" + strings.ToUpper(name)

        // create the new group specific chains
        err = iptables.NewChain(groupChain)
        if err != nil {
            log.Println("failed to create group chain", groupChain + ":", err.Error())
            return -1
        }

        // add all allows to the group specific chains
        for _, destination := range group.Allow {
            err = iptables.Chain(groupChain).Append().Destination(destination).Accept().Apply()
            if err != nil {
                log.Println("failed to add allowed destination", destination + ":", err.Error())
            }
        }
        
        log.Println("created chain", groupChain)
    }

    return 0
}