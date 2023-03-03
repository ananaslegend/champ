package main

import fastchamp "champ/fast"

func main() {
	var ch = fastchamp.Server{}

	ch.ListenAndServe("0.0.0.0:5000")
}
