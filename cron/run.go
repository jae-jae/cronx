package cron

func Run() {
	cmd := NewCMD()
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
