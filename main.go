package main

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"github.com/goccy/go-yaml"
	"github.com/seungyeop-lee/easycmd"
	"io"
	"log"
	"os"
	"os/signal"
	"time"
)

type Config struct {
	Schedulers []Scheduler `yaml:"schedulers"`
}

type Scheduler struct {
	Name    string `yaml:"name"`
	Command string `yaml:"command"`
	Cron    string `yaml:"cron"`
}

func main() {
	configFile, err := os.Open("example.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	config, err := io.ReadAll(configFile)
	if err != nil {
		log.Fatal(err)
	}

	c := Config{}
	err = yaml.Unmarshal(config, &c)
	if err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	scheduler, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range c.Schedulers {
		fmt.Printf("Add NewJob: %+v\n", s)
		_, err = scheduler.NewJob(
			gocron.CronJob(s.Cron, true),
			gocron.NewTask(
				func() {
					easycmd.New().Run(easycmd.Command(s.Command))
				},
			),
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	scheduler.Start()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = scheduler.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
}
