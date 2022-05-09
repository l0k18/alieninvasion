# alieninvasion
A simple CLI command to run a random driven simulation over a sparse grid

To generate a random world map file:

    go run ./cmd/worldgen/. <h> <v> <seed> <filename>

To simulate an invasion on the world map file:

    go run ./cmd/war/. <aliencount> <filename>

A one liner example:

    go run ./cmd/worldgen/. 20 20 20 /tmp/test;go run ./cmd/war/. 20 /tmp/test

City name list was generated with `go generate ./...` from 
[namegen/names.csv](namegen/names.csv) using field 2 of this file.