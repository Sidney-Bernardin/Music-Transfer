# MusicTransfer

Transfers a Youtube Music playlist to a Spotify playlist.

## Install

#### Clone the repo.
```
git clone https://github.com/Sidney-Bernardin/MusicTransfer.git
cd MusicTransfer
```

#### Environment variables
Once inside the MusicTransfer directory, you're going to need to set these environment variables.<br>

<strong>YOUTUBE_PLAYLIST_ID:</strong> To get your Youtube playlist's ID, just view your playlist on youtube.com. The playlist's ID will be in the URL bar.<br>
<strong>SPOTIFY_PLAYLIST_ID:</strong> To get your Spotify playlist's ID, just create a new playlist on open.spotify.com, once created click on the playlist  and the ID will be in the URL bar.<br>
<strong>YOUTUBE_API_KEY:</strong> To get your Youtube API key, just generate one in the Google Cloud Platform API credentials page. Make sure that the Youtube API is enabled in your GCP project.<br>
<strong>SPOTIFY_TOKEN:</strong> To get your Spotify token, go the <a href="https://developer.spotify.com/console/get-album/" target="_blank">this page</a>. Click on the get token button, check the playlist-modify-public scope then click request token. Your token will expire after an hour.<br>

## Run.
```
go build && ./MusicTransfer
```
