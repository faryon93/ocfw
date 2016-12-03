package config

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
    Group []string `toml:"groups"`
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

