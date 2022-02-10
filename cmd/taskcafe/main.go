package main

import (
	"github.com/olzhaskambar/taskcafe/internal/commands"
	_ "github.com/lib/pq"
)

func main() {
	commands.Execute()
}
