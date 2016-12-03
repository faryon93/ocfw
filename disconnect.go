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

func disconnect(conf *config.Config) (int) {
    // some metadata
    clientChain := "VPN_CLIENT_" + strings.ToUpper(ocenv.TunDevice)

    // delete references to the client chain
    err := iptables.Chain("FORWARD").
            Delete().
            SrcIf(ocenv.TunDevice).
            Jump(clientChain).
            Apply()
    if err != nil {
        log.Println("failed to delete jump rule:", err.Error())
    }   

    // delete chain
    err = iptables.DeleteChain(clientChain) 
    if err != nil {
        log.Println("failed to delete client chain", clientChain + ":", err.Error())
    }

    if err == nil {
        log.Println("successfully removed firewall rules for", ocenv.Username)
    }
    return 0
}