## Go Dungeon

This is a simple dungeon generator built in go

To use it, simply `import "github.com/meshiest/go-dungeon/dungeon"` and `go get`

To generate a dungeon, use: `dungeon := dungeon.NewDungeon(size int, numRooms int)` to generate a square dungeon with `size` side length and `numRooms` rooms

To preview a dungeon, you can `dungeon.Print()`, which will demo the dungeon using '#' to denote floor and ' ' to denote wall

To access the grid of a generated dungeon, you can use the `[][]int dungeon.Grid` where `1` denotes floor and `0` denotes wall

Check out *main.go* for an example! (`go run main.go`)

Here is a 100x100 with an attempted 200 rooms:

![](http://i.imgur.com/5aFnOGH.png)

Here is a 50x50 with an attempted 200 rooms rendered with opengl:

![](http://i.imgur.com/LPEtZSP.png)
