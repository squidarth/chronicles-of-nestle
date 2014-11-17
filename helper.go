package main

import (
  "github.com/howeyc/fsnotify"
  "log"
  "bufio"
  "os"
)

func chicken() {
  println("LALALALA")
}

func watch_bash_history(fileToWatch string) {
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
            case <-watcher.Event:
                counter += 1
                if (counter > 2) {
                  done <- true
                }

                println(counter)
                file, err := os.Open(fileToWatch)
                if err != nil {
                  println("error")
                }
                reader := bufio.NewReader(file)
                scanner := bufio.NewScanner(reader)

                scanner.Split(bufio.ScanLines)
                text := ""
                for scanner.Scan() {
                  text = scanner.Text()
                }
                log.Println("text: ", text)

            case err := <-watcher.Error:
                log.Println("error:", err)
            }
        }
    }()

    err = watcher.Watch(fileToWatch)
    if err != nil {
        log.Fatal(err)
    }

    msg := <-done

    println("printing msg")
    println(msg)

    watcher.Close()
}
