package db

import (
	"backend_master_class/db/util"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" // PostgreSQL sürücüsünü içe aktarır. This is because the “database/sql” package just provides a generic interface around SQL database. It needs to be used in conjunction with a database driver in order to talk to a specific database engine. We’re using postgres, so we will use lib/pq driver.
)

var testQueries *Queries // Bu, testler sırasında kullanılacak olan veritabanı sorgularını içerir.
var testDB *sql.DB       // Bu, veritabanı bağlantısını temsil eder.

func TestMain(m *testing.M) { // By convention, the TestMain function is the main entry point of all unit tests inside 1 specific golang package. // M is a type passed to a TestMain function to run the actual tests.
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource) // Open opens a database specified by its database driver name and a driver-specific data source name. Returns *DB and error
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB) // New: Bu, bir yapıcı fonksiyondur. DBTX arayüzünü alır ve yeni bir Queries yapısı döner. Bu yapı, testlerde kullanılacak veritabanı sorgularını içerir.

	os.Exit(m.Run()) // Run runs the tests. It returns an exit code to pass to os.Exit. // Tüm testleri çalıştırır (m.Run()) ve programı, testlerin sonucuna göre sonlandırır (os.Exit). m.Run işlevi, testlerin tamamlanması için gerekli olan tüm hazırlıkları yapar ve testleri yürütür.
}
