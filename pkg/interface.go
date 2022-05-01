package main

// Scratch padding ideas

// Transformation
// type Transformation interface {
// 	Register(register Register)
// 	Create(documents Documents)
// }

// type Documents interface {
// 	Get(name string) Document
// }

// type Document interface {
// 	Select(path string) interface{}
// }

// type Register interface {

// }

// type JoinedDocument struct {

// }

// // To prevent people from jumping out and grabing documents mid create. These documents can be lazily loaded so only if required are they generated. And even lazy subloaded.
// func Register(context string,registration string) {
// 	// registration.Lazy("A", document.http("HTTP address"), schema)
// 	// registration.Load("B", document.dynamodb("table", "id"), schema)
// 	// registration.From("A","C", func(document Document) {
// 	// 	return document.http(document.get("$.band_id"))
// 	// })
// }

// func Create(documents string) {
// 	// documents.ensure("A") -- starts loading documents on a background thread

// 	/*
// 	Transformation

// 	*/
// }

// // Document statuses
// Loaded, Used, Registered
