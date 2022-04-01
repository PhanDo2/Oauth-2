package initialization

import (
	"go-klikdokter/app/api/transport"
	"go-klikdokter/app/registry"
	"go-klikdokter/helper/database"
	"net/http"

	"github.com/go-kit/log"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func DbInit() (*gorm.DB, error) {
	// Init DB Connection
	db, err := database.NewConnectionDB(viper.GetString("database.driver"), viper.GetString("database.dbname"),
		viper.GetString("database.host"), viper.GetString("database.username"), viper.GetString("database.password"),
		viper.GetInt("database.port"))
	if err != nil {
		return nil, err
	}

	//Define auto migration here

	// example Seeder
	// for i := 0; i < 1000; i++ {
	// 	fmt.Println("dijalankan")
	// 	product := entity.Product{}
	// 	err := faker.FakeData(&product)
	// 	db.Create(&product)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }

	return db, nil
}

func InitRouting(db *gorm.DB, logger log.Logger) *http.ServeMux {
	// Service registry
	userSvc := registry.RegisterAuthService(db, logger)
	// Transport initialization

	userHttp := transport.UserHttpHandler(userSvc, log.With(logger, "UserTransportLayer", "HTTP"))
	// Routing path
	mux := http.NewServeMux()
	mux.Handle("/users/", userHttp)

	return mux
}
