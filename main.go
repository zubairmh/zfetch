package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/user"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/urfave/cli/v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func n(size uint64) string {
	v := roundFloat(float64(size)*math.Pow(1024, -3), 2)
	return fmt.Sprintf("%vGB", v)
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
			green := color.New(color.BgGreen).SprintFunc()
			fmt.Println(green(u.Username))
			fmt.Println()
			v, _ := mem.VirtualMemory()
			parts, _ := disk.Partitions(true)
			var free, total uint64
			for _, part := range parts {
				device := part.Mountpoint
				usage, _ := disk.Usage(device)
				free += usage.Free
				total += usage.Total

			}

			green_fg := color.New(color.FgHiGreen).SprintFunc()
			blue := color.New(color.BgHiBlue).SprintFunc()
			fmt.Println(blue("Memory"))
			fmt.Printf("Total: %s, Free: %s, Usage: %s\n\n", green_fg(n(v.Total)), green_fg(n(v.Free)), green_fg(fmt.Sprintf("%v%%", v.UsedPercent)))
			fmt.Println(blue("Disk"))
			fmt.Printf("Total: %s, Free: %s", green_fg(n(total)), green_fg(n(free)))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
