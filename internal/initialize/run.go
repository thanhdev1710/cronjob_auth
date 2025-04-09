package initialize

func Run() {
	LoadConfig()
	InitLogger()
	InitPostgreSql()
	InitCronJobs()

	select {}
}
