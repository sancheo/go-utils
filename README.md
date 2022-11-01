# how to use

*log.go*
```
// info
log.Logger.Info().Msg(string(msg))

// error
log.Logger.Error().Err(err).Msg("")
```

*user.go*
```
userId := user.GenerateUserID()
```