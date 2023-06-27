# how to use
```
go get github.com/sancheo/go-utils
```

*log.go*
```
// info
log.Logger.Info().Msg(string(msg))

// error
log.Logger.Error().Err(err).Msg("")
```

*ids.go*
```
id := ids.GenerateID()
```