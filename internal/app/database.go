package app

func (a *App) MustRunPSQLMigration(MigrationPATH string) {
	if err := a.psqlDB.RunMigrate(
		"file://"+MigrationPATH,
		a.psqlDatabaseURL,
	); err != nil {
		panic(err)
	}
}
