package er

type JZError string

func (e JZError) Error() string {
	return string(e)
}
