// SPDX-License-Identifier: 0BSD
// Author: Makkhawan Sardlah

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Args() {
	Args := os.Args
	var setup struct {
		Out   bool
		Video []SetUp
	}

	var Video string
	var Out bool

	Video = "nil"

	if len(Args) < 2 {
		fmt.Println("Args >> error: No Action")
		return
	}
	for i := 1; i < len(Args); i++ {
		if len(Args[i]) > 0 && Args[i][0] == '-' {
			argv := Args[i]
			for v := 0; v < len(argv); v++ {
				if argv[v] == 'V' {
					fmt.Printf("\033[33m Version:\033[36m %s \033[0m\n", version)
				} else if argv[v] == 'C' {
					cacheDir, err := os.UserCacheDir()
					if err == nil {
						os.RemoveAll(filepath.Join(cacheDir, "get-yt"))
					}
				} else if argv[v] == 'D' {
					if v > len(argv) {
						continue
					}
					v++
					if strings.HasPrefix(argv[v:], "mp4") {
						Video = "mp4"
						v += 2
					} else if strings.HasPrefix(argv[v:], "m4a") {
						Video = "m4a"
						v += 2
					} else if strings.HasPrefix(argv[v:], "m4v") {
						Video = "m4v"
						v += 2
					} else if strings.HasPrefix(argv[v:], "mp4u") {
						Video = "mp4u"
						v += 3
					}
				} else if argv[v] == 'O' {
					setup.Out = true
				}
			}
		} else if Video != "nil" {
			if setup.Out == true {
				if Out == true {
					setup.Video[len(setup.Video)-1].PathOut = Args[i]
				} else {
					setup.Video = append(setup.Video, SetUp{})
					setup.Video[len(setup.Video)-1].Video = Video
					setup.Video[len(setup.Video)-1].Out = true
					setup.Video[len(setup.Video)-1].Id = Args[i]
				}
				Out = !Out
			} else {
				setup.Video = append(setup.Video, SetUp{})
				setup.Video[len(setup.Video)-1].Video = Video
				setup.Video[len(setup.Video)-1].Out = false
				setup.Video[len(setup.Video)-1].Id = Args[i]
			}
		}
	}
	for i := 0; i < len(setup.Video); i++ {
		AutoDownload(setup.Video[i])
	}
}
