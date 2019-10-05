package main

import (
	"log"
	"os"

	dp "github.com/gino8070/dynamodb-stream-pipe"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	os.Exit(_main())
}

func _main() int {
	endpoint := kingpin.Flag("endpoint", "dynamodb endpoint(optional).").String()
	table := kingpin.Flag("table", "dynamodb table name.").Required().String()
	command := kingpin.Flag("command", "execute command. ex --command=wc").Required().String()
	args := kingpin.Flag("args", "comma separated command args(optional). ex --args=-l").String()
	kingpin.Parse()

	app, err := dp.NewApp(*endpoint, *table, *command, *args)
	if err != nil {
		log.Println(err)
		return 1
	}
	if err = app.Run(); err != nil {
		log.Println(err)
		return 1
	}
	return 0
}
