package models

import "github.com/hertz-contrib/websocket"

var Conns = make(map[int]map[*websocket.Conn]bool)
