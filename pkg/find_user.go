package pkg

import (
	"os"

	"github.com/gildemberg-santos/crudmongodb/pkg"
	"github.com/joho/godotenv"
)

func FindByIdClient(id string) interface{} {
	filter := map[string]interface{}{
		"id-usuario": id,
	}

	mongodb := pkg.MongoDB{
		SettingLog: true,
	}

	mongodb.Connect()

	godotenv.Load()

	return mongodb.FindOne(filter, os.Getenv("MONGO_DB"), os.Getenv("MONGO_COLLECTION"))
}
