package main

import (
	handler "Go-Fundamentals/src/advance/api"
	"Go-Fundamentals/src/basics"
	"fmt"
)

func main() {

	fmt.Println("Hello World")

	basics.VariableExp()
	//basics.ConditionalExp()
	//basics.SwitchExp()
	//basics.FunctionExp()
	//basics.LoopExp()
	//basics.ArrayExp()
	//basics.SliceExp()
	//basics.MapExp()
	//basics.StructExp()
	//basics.MethodExp()
	//basics.GoroutineExp()
	//basics.ChannelExp()
	//basics.SynchonizationUsingChannelExp()
	//basics.MutexExp()

	handler.HandleRequests()

}
