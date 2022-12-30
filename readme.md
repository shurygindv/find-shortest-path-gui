## Preview

TODO

## Installation

1. Install https://go.dev/doc/install
2. git clone https://github.com/shurygindv/shortest-path-between-cities
3. cd shortest-path-between-cities
3. go run .

## Description

There is a file with a description of the map of the city. Each road has third weights
1. A maximum bandwidth of the road
2. current load level in percentage 
3. and distance between cities

The possible file structure:

N - number of nodes (cities)
X1 Y1 // nodes coordinates
X2 Y2
...
XN YN
M - number of connections.
N1 K1 B1 L1 D1 //which nodes are connected (N1, K1), bandwidth (B1), load level (L1), distance (D1)
N2 K2 B2 L2 D2 
...
NM KM BM LM DM


## Solution
https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm

## Stack

Go 1.19 +fyne (interface)

Go: https://go.dev/
Fyne: https://fyne.io/

