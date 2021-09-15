package mq

type Instance interface {
	Push()
	MustEmbedDefaultInstance()
}

type DefaultInstance struct{}

func (instance *DefaultInstance) MustEmbedDefaultInstance() {
}

type Option interface {
	MustEmbedDefaultOption()
}

type DefaultOption struct{}

func (option *DefaultOption) MustEmbedDefaultOption() {
}
