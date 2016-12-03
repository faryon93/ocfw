package iptables
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
    "os/exec"
)


// ----------------------------------------------------------------------------------
//  constants
// ----------------------------------------------------------------------------------

const (
    IPTABLES = "iptables"
)


// ----------------------------------------------------------------------------------
//  private functions
// ----------------------------------------------------------------------------------

func iptables(args ...string) (error) {
    // build the command and execute
    cmd := exec.Command(IPTABLES, args...)
    err := cmd.Start()
    if err != nil {
        return err
    }

    // wait for finish
    return cmd.Wait()
}