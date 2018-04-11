package rpc

type Resolver interface {
	Resolve(binarizer Binarizer) (Binarizer, error)
}
