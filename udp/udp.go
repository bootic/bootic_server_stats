package udp

import (
  "net"
  "log"
  "time"
  "encoding/json"
)

type Event struct {
	Type string   `json:"type"`
	Time time.Time `json:"time"`
	Data map[string]interface{} `json:"data"`
}

var udpConn *net.UDPConn

func Init(hostAndPort string) {
  udpAddr, err := net.ResolveUDPAddr("udp", hostAndPort)
  if err != nil { panic("Could not connect to UDP server") }
  var err2 error
  udpConn, err2 = net.DialUDP("udp", nil, udpAddr) 
  if err2 != nil { panic(err2) }
}

func Send(typeStr string, data map[string]interface{}) {
  
  defer func() {
    if err := recover(); err != nil {
      log.Println("Goroutine failed:", err)
    }
  }()
  
  event := Event{
    Time: time.Now(),
    Type: typeStr,
    Data: data,
  }
  
  json, err := json.Marshal(event)
  
  if err != nil {
    log.Println("Could not marshal JSON:", err)
  }
  
  udpConn.Write(json)
}