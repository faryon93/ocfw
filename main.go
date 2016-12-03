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
    "os"

    "github.com/faryon93/ocfw/ocenv"
    "github.com/faryon93/ocfw/iptables"
)


// ----------------------------------------------------------------------------------
//  application entry
// ----------------------------------------------------------------------------------

func main() {
    // some metadata
    clientChain := "VPN_CLIENT_" + strings.ToUpper(ocenv.TunDevice)

    // a connect job
    if ocenv.IsConnect() {
        // create chain for the client
        err := iptables.NewChain(clientChain)
        if err != nil {
            log.Println("failed to create client chain:", err.Error())
            os.Exit(-1)
        }

        // add some allowed hosts for this client
        err = iptables.Chain(clientChain).
                Append().
                Destination("192.168.2.254").
                Accept().
                Apply()
        if err != nil {
            log.Println("failed to populate client chain with custom hosts:", err.Error())
            os.Exit(-1)
        }

        // the client should jump to its own chain
        err = iptables.Chain("FORWARD").
                Prepend().
                SrcIf(ocenv.TunDevice).
                Jump(clientChain).
                Apply()
        if err != nil {
            log.Println("failed to apply jump rule:", err.Error())
            os.Exit(-1)
        }   

        log.Println("successfully set up firewall for", ocenv.Username)

    // this is a disconnect call
    } else if (ocenv.IsDisconnect()) {
        // flush client chain
        err := iptables.FlushChain(clientChain)
        if err != nil {
            log.Println("failed to flush client chain", clientChain + ":", err.Error())
            os.Exit(-1)
        }

        // delete references to the client chain
        err = iptables.Chain("FORWARD").
                Delete().
                SrcIf(ocenv.TunDevice).
                Jump(clientChain).
                Apply()
        if err != nil {
            log.Println("failed to delete jump rule:", err.Error())
            os.Exit(-1)
        }   

        // delete chain
        err = iptables.DeleteChain(clientChain) 
        if err != nil {
            log.Println("failed to delete client chain", clientChain + ":", err.Error())
            os.Exit(-1)
        }

        log.Println("successfully removed firewall rules for", ocenv.Username)
    }
}