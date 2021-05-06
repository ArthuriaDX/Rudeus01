package interfaces

type WMarkDown interface {
	Append(md WMarkDown) WMarkDown
	ToString() string
}
