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


func benchmarkQuerying(query string, b *testing.B) {
    config()
    loadSearchItems()
    for n := 0; n < b.N; n++ {
       SearchQuery(query, 17)
    }
}


func BenchmarkQueryG(b *testing.B) { benchmarkQuerying("G", b) }
func BenchmarkQueryGloves(b *testing.B) { benchmarkQuerying("Gloves", b) }
func BenchmarkQueryGlovesG(b *testing.B) { benchmarkQuerying("Gloves Germiph", b) }
