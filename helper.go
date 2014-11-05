package main

import (
  "github.com/howeyc/fsnotify"
  "log"
)

func chicken() {
  println("LALALALA")
}

func watch_bash_history() {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        println("some error")
        log.Fatal(err)
    }

    println("creating done")
    done := make(chan bool)

    // Process events
    go func() {
        counter := 0
        for {
            select {
            case ev := <-watcher.Event:
                counter += 1
                if (counter > 2) {
                  done <- true
                }

                println(counter)
                log.Println("event:", ev)
            case err := <-watcher.Error:
                log.Println("error:", err)
            }
        }
    }()

    err = watcher.Watch("shloom")
    if err != nil {
        log.Fatal(err)
    }

    msg := <-done

    println("printing msg")
    println(msg)

    watcher.Close()
}
