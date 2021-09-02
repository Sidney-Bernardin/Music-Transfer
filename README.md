# MusicTransfer

Transfers a Youtube Music playlist to a Spotify playlist.

## Install
#### Environment variables
You are going to need to set these env-vars set before installing.<br>
YOUTUBE_API_KEY, YOUTUBE_PLAYLIST_ID, SPOTIFY_TOKEN, SPOTIFY_PLAYLIST_ID.

#### Getting the tokens and playlist IDs.
* To get your Youtube playlist's ID, just view your playlist on youtube.com. The playlist's ID will be in the URL.<br>
* To get your Spotify playlist's ID, just view your playlist on open.spotify.com. The playlist's ID will be in the URL.<br>
* To get your Youtube API key, just make one in the Google Cloud Platform API credentials page. Make sure that the Youtube API is enabled in your GCP project.<br>
* To get your Spotify token, go the this page https://developer.spotify.com/console/get-album/. Click on the get token button, don't check any scopes then click request token. Your token will expire after an hour.<br>

#### Clone the repo.
```
git clone https://github.com/Sidney-Bernardin/MusicTransfer.git
```

## Run.
cd into the repo then run with
```
go build && ./MusicTranfer
```
