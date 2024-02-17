# Building plugins
```bash
go build -buildmode=plugin -o plugins/yourPlugin.so yourPlugin.go
go build -buildmode=plugin -o plugins/myPlugin.so myPlugin.go
```

# Run it
```bash
go run main.go 
```

