package config
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
    "github.com/BurntSushi/toml"
)


// ----------------------------------------------------------------------------------
//  types
// ----------------------------------------------------------------------------------

type Config struct {
    Common struct {
        LockFile string `toml:"lock_file"`
    }
    Users map[string]User `toml:"user"`
    Groups map[string]Group `toml:"group"`
}

type User struct {
    Allow []string `toml:"allow"`
    Groups []string `toml:"groups"`
}

type Group struct {
    Allow []string `tomls:"groups"`
}


// ----------------------------------------------------------------------------------
//  functions
// ----------------------------------------------------------------------------------

func Load(path string) (*Config, error) {
    // decode the config file to struct
    var config Config
    if _, err := toml.DecodeFile(path, &config); err != nil {
        return nil, err
    }

    return &config, nil
}
