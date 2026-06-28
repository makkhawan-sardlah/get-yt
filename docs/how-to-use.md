# how-to-use

- [check version](#check-version)
- [download video](#download-video)
- [download sound](#download-sound)
- [download videod (upgrade)](#download-video-upgrade)
- [output video file](#output-video-file)

## check version

```bash
get-yt -V
# Version: 2.0
```

## download video

```bash
get-yt -Dmp4 https://www.youtube.com/watch?v=123456789ab
#or
get-yt -Dmp4 123456789ab
```

But you won't get the best video quality.

[download videod (upgrade)](#download-video-upgrade)

## download sound

```bash
get-yt -Dm4a https://www.youtube.com/watch?v=123456789ab
#or
get-yt -Dm4a 123456789ab
```

## download video (upgrade)

**This function was removed in version 2.0.**

**However, this function will be reintroduced after its performance and security have been improved.**

What is mp4u?

This process combines m4a and m4v files to achieve the highest video quality.

```bash
get-yt -Dmp4u https://www.youtube.com/watch?v=123456789ab
#or
get-yt -Dmp4u 123456789ab
```

## output video file

```bash
get-yt -Dmp4O https://www.youtube.com/watch?v=123456789ab file1.mp4 https://www.youtube.com/watch?v=123456789ab file2.mp4
#or
get-yt -Dmp4O 123456789ab file1.mp4 123456789ab file2.mp4
```

## clean cache

```bash
get-yt -C
```
