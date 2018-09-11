sciondSock := sciond.GetDefaultSCIONDPath(&config.LocalAddr.IA)
dispatcher := getDefaultDispatcherSock()
if err := snet.Init(config.LocalAddr.IA, sciondSock, dispatcher); err != nil {
	log.Crit("Unable to initialize SCION network", "err", err)
}
