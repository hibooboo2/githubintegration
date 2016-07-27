package githubapi

var log *logger

func init() {
	log = newLogger()
}
