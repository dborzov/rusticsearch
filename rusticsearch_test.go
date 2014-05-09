package main

import (
  "testing"
  )


func TestQuerying(t *testing.T) {
    config()
    loadSearchItems()
    am := SearchQuery("gloves", 12)
    if len(am) == 0 {
      t.Error("No search results for query \"gloves\"", am)
    }
}
