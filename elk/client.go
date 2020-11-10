package elk

type Client interface {
	// AddDocAsync add generic doc to elk
	AddDocAsync(doc interface{})
	// AddLogAsync add whole log body as "log" field with auto timestamp to elk
	AddLogAsync(log interface{})
}
