package main

import "github.com/meshiest/go-dungeon/dungeon"

func main() {
  dungeon := dungeon.NewDungeon(100, 200)
  dungeon.Print()
}