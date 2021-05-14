package YTElkProducer

type Client interface {
	AddDoc(doc interface{})
	// AddDocAsync add generic doc to elk
	AddDocAsync(doc interface{})

	AddLog(log interface{})
	// AddLogAsync add whole log body as "log" field with auto timestamp to elk
	AddLogAsync(log interface{})
}
