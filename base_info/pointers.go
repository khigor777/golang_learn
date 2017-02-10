package base_info



var Name = "Hello"

func Change(name *string)  {
	*name += "test"
}
