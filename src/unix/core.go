// SPDX-License-Identifier: 0BSD
// Author: Makkhawan Sardlah

package main

import (
	"github.com/kkdai/youtube/v2"
)

type SetUp struct {
	Video   string
	Out     bool
	Id      string
	PathOut string
}

func AutoDownload(setup SetUp) {
	var out []string
	if setup.Out == true {
		if len(setup.PathOut) == 0 {
			return
		}
		out = []string{setup.PathOut}
	}
	switch setup.Video {
	case "mp4":
		Download_mp4(setup.Id, out)
	case "m4a":
		Download_m4a(setup.Id, out)
	case "m4v":
		Download_m4v(setup.Id, out)
	}
}

func Download_mp4(id string, out []string) {
	videoID := Check_id(id)

	if videoID == "nil" {
		return
	}

	client, video, err := Create_client(videoID)

	if err != nil {
		return
	}

	formats := video.Formats.WithAudioChannels()

	if len(out) == 0 {
		Create_file(video.Title+".mp4", video, &formats[0], &client)
	} else {
		for i := 0; i < len(out); i++ {
			Create_file(out[i], video, &formats[0], &client)
		}
	}
}

func Download_m4a(id string, out []string) {
	videoID := Check_id(id)

	if videoID == "nil" {
		return
	}

	client, video, err := Create_client(videoID)

	if err != nil {
		return
	}

	formats := video.Formats.Type("audio")

	var targetFormat *youtube.Format
	for i := range formats {
		f := &formats[i]
		if f.FPS == 0 && f.AudioChannels > 0 {
			targetFormat = f
			break
		}
	}

	if len(out) == 0 {
		Create_file(video.Title+".m4a", video, targetFormat, &client)
	} else {
		for i := 0; i < len(out); i++ {
			Create_file(out[i], video, targetFormat, &client)
		}
	}
}

func Download_m4v(id string, out []string) {
	videoID := Check_id(id)

	if videoID == "nil" {
		return
	}

	client, video, err := Create_client(videoID)

	if err != nil {
		return
	}

	formats := video.Formats.Type("video")
	var targetFormat *youtube.Format

	for i := range formats {
		f := &formats[i]
		if f.FPS > 0 && f.AudioChannels == 0 {
			targetFormat = f
			break
		}
	}

	if len(out) == 0 {
		Create_file(video.Title+".m4v", video, targetFormat, &client)
	} else {
		for i := 0; i < len(out); i++ {
			Create_file(out[i], video, targetFormat, &client)
		}
	}
}
