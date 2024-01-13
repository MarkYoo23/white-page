package exec

type Exec interface {
	Execute(args []string) error
}
