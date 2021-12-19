package api

type Symbol struct {
	Symbol   string  `bson:"symbol"`
	SymbolID string  `bson:"symbol_id"`
	Name     string  `bson:"name"`
	LTP      float32 `bson:"ltp"`
}
