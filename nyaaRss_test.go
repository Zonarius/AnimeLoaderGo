package AnimeLoaderGo

import (
    "testing"
)

func TestFetch(t *testing.T) {
    t.Parallel()
    items, err := Fetch("baked")
    if err != nil {
        t.Error(err)
    }
    len := len(items)
    if len == 0 {
        t.Errorf("Expected to get some items, instead got %d", len)
    }
}