package main

import (
	"fmt"

	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-os/db"
	"github.com/pborman/uuid"
)

func main() {
	cmd.Init()

	database := db.NewDB(
		db.Database("foo"),
		db.Table("bar"),
	)

	type Thing struct {
		Name string
	}

	record := db.NewRecord(uuid.NewUUID().String(), db.Metadata{"key": "value"}, &Thing{"dbthing"})

	fmt.Printf("Creating record: id: %s metadata: %+v bytes: %+v\n", record.Id(), record.Metadata(), string(record.Bytes()))
	if err := database.Create(record); err != nil {
		fmt.Println(err)
		return
	}

	rec, err := database.Read(record.Id())
	if err != nil {
		fmt.Println(err)
		return
	}

	thing := new(Thing)

	if err := rec.Scan(&thing); err != nil {
		fmt.Println("Error scanning read record", err)
		return
	}

	fmt.Printf("Read record: id: %s metadata: %+v bytes: %+v\n", rec.Id(), rec.Metadata(), thing)

	fmt.Println("Searching for metadata key:value")
	records, err := database.Search(
		db.WithMetadata(db.Metadata{"key": "value"}),
		db.WithLimit(10),
		db.WithOffset(0),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, record := range records {
		thing := new(Thing)

		if err := record.Scan(&thing); err != nil {
			fmt.Println("Error scanning record", record.Id(), err)
			return
		}

		fmt.Printf("Record: id: %s metadata: %+v bytes: %+v\n", record.Id(), record.Metadata(), thing)
	}

	fmt.Println("Deleting", record.Id())
	if err := database.Delete(record.Id()); err != nil {
		fmt.Println(err)
		return
	}
}
