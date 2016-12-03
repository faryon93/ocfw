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
    "strings"
)


// ----------------------------------------------------------------------------------
//  typedef
// ----------------------------------------------------------------------------------

type Rule string


// ----------------------------------------------------------------------------------
//  public functions
// ----------------------------------------------------------------------------------

func Chain(name string) (Rule) {
    return Rule(name)
}

func (r Rule) Prepend() (Rule) {
    if strings.HasPrefix(string(r), "-I") ||
       strings.HasPrefix(string(r), "-A") ||
       strings.HasPrefix(string(r), "-D"){
        return r
    }

    return "-I " + r
}

func (r Rule) Append() (Rule) {
    if strings.HasPrefix(string(r), "-I") ||
       strings.HasPrefix(string(r), "-A") ||
       strings.HasPrefix(string(r), "-D"){
        return r
    }

    return "-A " + r
}

func (r Rule) Delete() (Rule) {
    if strings.HasPrefix(string(r), "-I") ||
       strings.HasPrefix(string(r), "-A") ||
       strings.HasPrefix(string(r), "-D"){
        return r
    }

    return "-D " + r
}


func (r Rule) SrcIf(iface string) (Rule) {
    if strings.Contains(string(r), "-i",) {
        return r
    }

    return r + " -i " + Rule(iface)
}

func (r Rule) Destination(dst string) (Rule) {
    if strings.Contains(string(r), "-d",) {
        return r
    }

    return r + " -d " + Rule(dst)
}

func (r Rule) Accept() (Rule) {
    if strings.Contains(string(r), "-j") {
        return r
    }

    return r + " -j ACCEPT"
}

func (r Rule) Drop() (Rule) {
    if strings.Contains(string(r), "-j") {
        return r
    }

    return r + " -j DROP"
}

func (r Rule) Jump(chain string) (Rule) {
    if strings.Contains(string(r), "-j") {
        return r
    }

    return r + " -j " + Rule(chain)
}

func (r Rule) Apply() (error) {
    cmd := strings.Split(string(r), " ")
    return iptables(cmd...)
}
