package main

import (
	"fmt"

	"github.com/gildemberg-santos/retrospectiva/pkg"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	var data bson.Raw = pkg.FindByIdClient("327").(bson.Raw)
	fmt.Println(data.Lookup("id-empresa").StringValue())
}
