package commandlib

//TODO: should be used as a library

type Command interface {
	Execute(service interface{}) (interface{}, error)
}

type CommandHandler interface {
	ExecuteCommand(cmd Command) (interface{}, error)
}

type commandHandler struct {
	service interface{}
}

func NewCommandHandler(service interface{}) CommandHandler {
	return &commandHandler{service: service}
}

func(c *commandHandler) ExecuteCommand(cmd Command) (interface{}, error) {
	return cmd.Execute(c.service)
}