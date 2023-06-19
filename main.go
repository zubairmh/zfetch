package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/user"

	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/urfave/cli/v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func n(size uint64) float64 {
	return roundFloat(float64(size)*math.Pow(1024, -3), 2)
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func main() {
	app := &cli.App{
		Name:  "ZFetch",
		Usage: "Display current system stats",
		Action: func(*cli.Context) error {
			u, _ := user.Current()
			fmt.Println(u.Username)
			v, _ := mem.VirtualMemory()
			parts, _ := disk.Partitions(true)
			var free, total uint64
			for _, part := range parts {
				device := part.Mountpoint
				usage, _ := disk.Usage(device)
				free += usage.Free
				total += usage.Total

			}

			// almost every return value is a struct
			fmt.Printf("Memory Total: %vGB, Free:%vGB, UsedPercent:%d%%\n", n(v.Total), n(v.Free), int(v.UsedPercent))
			fmt.Printf("Disk Total: %vGB, Free:%vGB", n(total), n(free))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
