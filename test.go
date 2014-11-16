package main

import (
  "log"
  "os"
  "os/signal"
  "syscall"
  //"encoding/hex"
  rtl "github.com/jpoirier/gortlsdr"
)

func sig_abort(dev *rtl.Context) {
  //Make a Channel of os.Signal objects
  ch := make(chan os.Signal)
  signal.Notify(ch, syscall.SIGINT)
  //Wait for the Channel to return something.  Discarding
  //the retruned value since we know it's SIGINT
  <-ch
  //This code will execute after SIGINT is returned
  _ = dev.CancelAsync()
  dev.Close()
  os.Exit(0)
}

func main() {
  var err error
  var dev *rtl.Context

  if c:= rtl.GetDeviceCount(); c== 0 {
    log.Fatal("No devices found, exiting.\n")
  }

  if dev, err = rtl.Open(0); err != nil {
    log.Fatal("\tOpen Failed, existing\n")
  }
  //defer pushes a function onto a list that will be invoked
  //after the surrounding function returns (in this case main)
  defer dev.Close()
  go sig_abort(dev)

  //Set to 1090Mhz
  dev.SetSampleRate(2000000) //from rtl_adsb.c
  dev.SetTunerGainMode(false)
  err = dev.SetCenterFreq(1090000000)
  if err != nil {
    log.Printf("\tSetCenterFreq 1090Mhz Failed, error: %s\n", err)
  }

  dev.SetTestMode(true)
  dev.ResetBuffer()

  var buffer []byte = make([]uint8, rtl.DefaultBufLength)
  //var hexbuf []byte = make([]uint8, rtl.DefaultBufLength)
  n_read, err := dev.ReadSync(buffer, rtl.DefaultBufLength)
  if err != nil {
    log.Printf("\tReadSynch Failed - error %s\n", err)
  } else {
    //hex.Decode(hexbuf, buffer)
    log.Printf("\tReadSync %d\n", n_read)
    //log.Printf("\t%X\n", hexbuf)
  }
  if err == nil && n_read < rtl.DefaultBufLength {
    log.Printf("ReadSynch short read, %d samples lost\n", rtl.DefaultBufLength-n_read)
  }
  log.Printf("Exiting...\n")
}
