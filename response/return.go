package response

type ReturnData struct {
	Map   map[string]any
	Data  any
	Msg   string
	Error bool
}

func ReturnSuccess(data any, msg ...string) ReturnData {
	var returnData ReturnData
	if len(msg) != 0 {
		returnData.Msg = msg[0]
	}
	returnData.Data = data
	returnData.Error = true
	return returnData
}
func ReturnError(msg string) ReturnData {
	var returnData ReturnData
	returnData.Msg = msg
	returnData.Error = false

	return returnData
}
