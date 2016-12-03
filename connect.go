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
    "log"
    "strings"

    "github.com/faryon93/ocfw/ocenv"
    "github.com/faryon93/ocfw/iptables"
    "github.com/faryon93/ocfw/config"
)


// ----------------------------------------------------------------------------------
//  functions
// ----------------------------------------------------------------------------------

func connect(conf *config.Config) (int) {
    // some metadata
    clientChain := "VPN_CLIENT_" + strings.ToUpper(ocenv.TunDevice)

    // check if there is some configuration for the user
    user, valid := conf.Users[ocenv.Username]
    if !valid {
        log.Println("no rules configured for user", ocenv.Username)
        return 0
    }

    // create chain for the client
    err := iptables.NewChain(clientChain)
    if err != nil {
        log.Println("failed to create client chain:", err.Error())
        return -1
    }

    // add the allowed hosts to the client chain
    for _, allowed := range user.Allow {
        // add some allowed hosts for this client
        err = iptables.Chain(clientChain).Append().Destination(allowed).Accept().Apply()
        if err != nil {
            log.Println("failed to populate client chain with custom hosts:", err.Error())
            return -1
        }    
    }

    // add a jump to all group chains the user is assigned to
    for _,group := range user.Groups {
        // metadata
        groupChain := "VPN_GROUP_" + strings.ToUpper(group)

        // check if the assigned group is a valid group
        if _, valid := conf.Groups[group]; !valid {
            log.Println("user", ocenv.Username, "is assigned to invalid group", group)
            continue
        }

        // add the jump
        err = iptables.Chain(clientChain).Prepend().Jump(groupChain).Apply()
        if err != nil {
            log.Println("failed to add group jump rule", groupChain + ":", err.Error())
            return -1
        }
    }

    // make the client use its own chain
    err = iptables.Chain("FORWARD").Prepend().SrcIf(ocenv.TunDevice).Jump(clientChain).Apply()
    if err != nil {
        log.Println("failed to apply jump rule:", err.Error())
        return -1
    }   

    log.Println("successfully set up firewall for", ocenv.Username)
    return 0
}