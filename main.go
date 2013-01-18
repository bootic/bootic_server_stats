package main

import (
    "log"
    "os/exec"
    "strconv"
    "regexp"
    "strings"
)

type ServerInfo struct {
  
}

func (self ServerInfo) Hostname() (hname string) {
  out, err := exec.Command("hostname").Output()
  
  if err!=nil {
      hname = ""
      return
  }
  
  hname = string(out)
  hname = strings.Replace(hname, "\n", "", -1)
  return
}

func (self ServerInfo) AvgLoad() (loadFloat float64) {
  
  out, err := exec.Command("uptime").Output()
  
  if err!=nil {
      return
  }
  
  line := string(out)
  
  reg, _ := regexp.Compile(`(?:load averages:)\s+([0-9]+\.[0-9]+)\s+`)
  r := reg.FindStringSubmatch(line)
  
  var load string

  for _, v := range r {
    load = v
  }
  
  loadFloat, _ = strconv.ParseFloat(load, 64)
  
  return
  
}

func main() {
  foo := new(ServerInfo)
  log.Println("Current server load for", foo.Hostname(), "is", foo.AvgLoad())
}