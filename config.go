package AnimeLoaderGo

import (
    "os"
    "encoding/json"
    "io/ioutil"
)

const (
    filename = "config.json"
)

type Config struct {
    TmpDir string
}

func defaultConfig() *Config {
    return &Config{"/var/tmp/"}
}

func ReadConfig() (cfg *Config, err error) {
    file, err := os.Open(filename)
    if os.IsNotExist(err) {
        cfg := defaultConfig()
        WriteConfig(cfg)
        return cfg, nil
    } else if err != nil {
        return nil, err
    }
    dec := json.NewDecoder(file)
    ret := Config{}
    err = dec.Decode(&ret)
    return &ret, err
}

func WriteConfig(cfg *Config) (err error) {
    b, err := json.MarshalIndent(cfg, "", "  ")
    if err != nil {
        return err
    }
    err = ioutil.WriteFile(filename, b, 0664)
    return err
}