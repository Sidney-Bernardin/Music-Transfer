# MusicTransfer

Transfers a Youtube Music playlist to a Spotify playlist.

## Install

#### Clone the repo.
```
git clone https://github.com/Sidney-Bernardin/MusicTransfer.git
```

#### Environment variables
You are going to need to set these env-vars before installing.<br>
YOUTUBE_API_KEY, YOUTUBE_PLAYLIST_ID, SPOTIFY_TOKEN, SPOTIFY_PLAYLIST_ID.

TOUTUBE_PLAYLIST_ID: To get your Youtube playlist's ID, just view your playlist on youtube.com. The playlist's ID will be in the URL.<br>
SPOTIFY_PLAYLIST_ID: To get your Spotify playlist's ID, just view your playlist on open.spotify.com. The playlist's ID will be in the URL.<br>
YOUTUBE_API_KEY: To get your Youtube API key, just make one in the Google Cloud Platform API credentials page. Make sure that the Youtube API is enabled in your GCP project.<br>
SPOTIFY_TOKEN To get your Spotify token, go the this page https://developer.spotify.com/console/get-album/. Click on the get token button, don't check any scopes then click request token. Your token will expire after an hour.<br>

## Run.
cd into the repo then run with
```
go build && ./MusicTransfer
```
