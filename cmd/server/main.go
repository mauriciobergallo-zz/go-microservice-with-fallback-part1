package main

import (
  mockfileserver "github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/adapters/mock-file-server"
  "github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/adding"
  "github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/deleting"
  "github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/deletingFallback"
  "github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/handlers"
  "github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/listing"
  "github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/logging"
  "github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/storage/memory"
  "github.com/mauriciobergallo/go-microservice-with-fallback-part1/pkg/updating"
)

func main() {
  // Creating Logging Service
  l := logging.NewStdoutLogging("DEBUG")
  r := new(memory.Storage)

  fs := new(mockfileserver.FileServer)

  ls := listing.NewService(r, l)
  as := adding.NewService(r, l)
  ds := deleting.NewService(r, fs, l)
  us := updating.NewService(r, l)
  dfs := deletingFallback.NewService(r, fs, l)

  handlers.NewRestService(as, ds, us, ls, dfs)
}
