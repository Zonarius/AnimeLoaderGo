package AnimeLoaderGo

import (
    "testing"
)

func TestReadConfig(t *testing.T) {
    cfg, err := ReadConfig()
    if err != nil {
        t.Error(err)
    }
    if cfg == nil {
        t.Error("Config is nil")
    }
}