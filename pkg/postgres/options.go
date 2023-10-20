package postgres

import (
	"fmt"
	"os"
)

func getConfigurationPostgres() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("pst_host"), os.Getenv("pst_user"), os.Getenv("pst_password"),
		os.Getenv("pst_dbname"), os.Getenv("pst_port"), os.Getenv("pst_sslmode"),
		os.Getenv("pst_timeZone"))
}
