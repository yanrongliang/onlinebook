package sysinit

func init() {
	sysinit()
	dbinit()
	dbinit("r")
	dbinit("uaw", "uar")
}
