package main

import (
    "log"
    "os/exec"
    "strconv"
    "regexp"
    "time"
    "os"
    "runtime"
    "bootic_server_stats/udp"
)

type ServerInfo struct {
  
}

func (self ServerInfo) AvgLoad() (float64) {
  
  out, err := exec.Command("uptime").Output()
  
  if err!=nil {
      return 0.0
  }
  
  cpus := runtime.NumCPU()
  
  line := string(out)
  
  reg, _ := regexp.Compile(`(?:load averages:)\s+([0-9]+\.[0-9]+)\s+`)
  r := reg.FindStringSubmatch(line)
  
  var load string

  for _, v := range r {
    load = v
  }
  
  loadFloat, _ := strconv.ParseFloat(load, 64)
  
  // we want CPU utilization for all available CPUs
  return loadFloat / float64(cpus)
  
}

func main() {
  udp_host      := os.Getenv("DATAGRAM_IO_UDP_HOST")
  intervalStr   := os.Getenv("INTERVAL")
	
  udp.Init(udp_host)
  
  server := new(ServerInfo)
  hostname, _ := os.Hostname()
  duration, err := time.ParseDuration(intervalStr)
  if err!=nil {
    panic("INTERVAL cannot be parsed")
  }
  ticker := time.NewTicker(duration) // 10 secs.
  
  log.Println("Reporting server load for", hostname, "every", intervalStr)
  
  for {
    select {
    case <- ticker.C:
      data := make(map[string]interface{})

      data["app"] = "server_stats"
      data["account"] = hostname
      data["status"] = server.AvgLoad()
      
      udp.Send("load_avg", data)
    }
  }

}