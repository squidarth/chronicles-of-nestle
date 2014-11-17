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
                file, err := os.Open("/Users/sidharthshanker/.zsh_history")
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

                log.Println(ev) 
            case err := <-watcher.Error:
                log.Println("error:", err)
            }
        }
    }()

    err = watcher.Watch("/Users/sidharthshanker/.zsh_history")
    if err != nil {
        log.Fatal(err)
    }

    msg := <-done

    println("printing msg")
    println(msg)

    watcher.Close()
}
