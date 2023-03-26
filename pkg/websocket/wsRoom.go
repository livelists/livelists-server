package websocket

type WsMessage struct {
	event string
	data  any
}
type WsRoom struct {
	messagesChannel chan WsMessage
}

func (wsR *WsRoom) Listen() *chan WsMessage {
	return &wsR.messagesChannel
}

func (wsR *WsRoom) Leave() {

}

func (wsR *WsRoom) Publish() {

}
