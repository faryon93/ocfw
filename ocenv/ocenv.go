package ocenv
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
    "os"
    "strings"
)


// ----------------------------------------------------------------------------------
//  global variables
// ----------------------------------------------------------------------------------

var (
    // user information
    Username string
    Group string
    Reason string

    // IPv4 addresses
    RealIp string
    RealLocalIp string
    LocalIp string
    RemoteIp string

    // IPv6 addresses
    LocalIpV6 string
    RemoteIpV6 string
    PrefixIpV6 string

    // routing
    TunDevice string
    Routes []string
    NoRoutes string
    Dns []string

    // statistics (only on disconnect)
    BytesIn string  // bytes
    BytesOut string // bytes
    Duration string // seconds
)


// ----------------------------------------------------------------------------------
//  initializer
// ----------------------------------------------------------------------------------

func init() {
    Username = os.Getenv("USERNAME")
    Group = os.Getenv("GROUP")
    TunDevice = os.Getenv("DEVICE")
    Reason = os.Getenv("REASON")
    RealIp = os.Getenv("IP_REAL")
    RealLocalIp = os.Getenv("IP_REAL_LOCAL")
    LocalIp = os.Getenv("IP_LOCAL")
    RemoteIp = os.Getenv("IP_REMOTE")
    LocalIpV6 = os.Getenv("IPV6_LOCAL")
    RemoteIpV6 = os.Getenv("IPV6_REMOTE")
    PrefixIpV6 = os.Getenv("IPV6_PREFIX")
    Routes = strings.Split(os.Getenv("OCSERV_ROUTES"), " ")
    NoRoutes = os.Getenv("OCSERV_NO_ROUTES")
    Dns = strings.Split(os.Getenv("OCSERV_DNS"), " ")
    BytesIn = os.Getenv("STATS_BYTES_IN")
    BytesOut = os.Getenv("STATS_BYTES_OUT")
    Duration = os.Getenv("STATS_DURATION")
}


// ----------------------------------------------------------------------------------
//  public functions
// ----------------------------------------------------------------------------------

func IsConnect() (bool) {
    return Reason == "connect"
}

func IsDisconnect() (bool) {
    return Reason == "disconnect"
}