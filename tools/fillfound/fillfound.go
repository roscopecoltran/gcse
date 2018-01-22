package main

import (
	"log"

	"github.com/golangplus/time"

	"github.com/roscopecoltran/gcse"
	"github.com/roscopecoltran/gcse/store"
	"github.com/roscopecoltran/gcse/utils"

	sppb "github.com/roscopecoltran/gcse/proto/spider"
)

func doFill() error {
	cDB := gcse.LoadCrawlerDB()
	return cDB.PackageDB.Iterate(func(pkg string, val interface{}) error {
		ent, ok := val.(gcse.CrawlingEntry)
		if !ok {
			log.Printf("Wrong entry, ignored: %+v", ent)
			return nil
		}
		site, path := utils.SplitPackage(pkg)
		return store.AppendPackageEvent(site, path, "unknown", ent.ScheduleTime.Add(-10*timep.Day), sppb.HistoryEvent_Action_None)
	})
}

func main() {
	if err := doFill(); err != nil {
		log.Fatalf("doFill failed: %v", err)
	}
}
