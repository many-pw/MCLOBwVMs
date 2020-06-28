In this example we'll deploy the video sharing site [jjaa.me](https://jjaa.me/) using mclob.

This app has 3 parts:

1. worker - takes newly uploaded video files and runs them through ffmpeg, uploads results to cloud storage
2. http - handles incoming http requests for uploading or watching videos, reads from cloud storage
3. mysql - one central database that records users and videos details

On day one all 3 parts will run in one tiny VM.

```
```

