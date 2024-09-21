package services

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/importcjj/sensitive"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	Filter *sensitive.Filter
	lock   sync.Mutex

	eventTimes   = make(map[string]time.Time)
	debounceTime = 1 * time.Second
)

func InitSensitive() {
	dir, _ := os.Getwd()
	dirPath := filepath.Join(dir, "dict")
	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalln("[InitSensitive] load dict err:", err)
	}

	Filter = sensitive.New()
	for _, file := range files {
		sensitiveFile := filepath.Join(dirPath, file.Name())
		err = Filter.LoadWordDict(sensitiveFile)
		if err != nil {
			log.Fatalln("[InitSensitiveWord] load sensitive file err:", err, ", file:", sensitiveFile)
		}
	}
}

func WaterDict() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer watcher.Close()

	dirPath, _ := os.Getwd()
	dirPath = filepath.Join(dirPath, "dict")
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		for {
			select {
			case err, ok := <-watcher.Errors:
				if !ok {
					fmt.Println(fmt.Sprintf("watcher errors: %s", err.Error()))
					return
				}
			case event, ok := <-watcher.Events:
				if !ok {
					fmt.Println("watcher event not ok")
					return
				}

				if !strings.Contains(event.Op.String(), "no events") {
					lock.Lock()
					lastEventTime, exists := eventTimes[event.Name]
					currentTime := time.Now()
					if !exists || currentTime.Sub(lastEventTime) > debounceTime {
						eventTimes[event.Name] = currentTime
						lock.Unlock()

						sensitiveFile := filepath.Join(dirPath, event.Name)
						err = Filter.LoadWordDict(sensitiveFile)
						if err != nil {
							log.Fatalln("[InitSensitiveWord] load sensitive file err:", err, ", file:", sensitiveFile)
						}
					} else {
						lock.Unlock()
					}
				}
			}
		}
	}()

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".txt") {
			_ = watcher.Add(filepath.Join("dict", file.Name()))
		}
	}

	<-make(chan struct{})
}
