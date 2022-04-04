package build

type BuildInfoError struct{}

func (e BuildInfoError) Error() string {
	return "unable to read build information"
}
