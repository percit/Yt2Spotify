# Yt2Spotify
Change your youtube playlists to spotify, as yt is starting to get worse

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/mzdunek)

To make it work:
- go through "GettingAPIs.md" to enable everything
- install Go (I used Go 20)
- ``go build -o yt2Spotify`` 
- ``./yt2Spotify -g GoogleApiToken -c SpotifyClientID -s SpotifyClientSecret``
- press numbers from 1 to 11 to add yt song to spotify playlist or just add to txt file

Known issues:
- unfortunetely for now you have to hard code SpotifyClientSecret and SpotifyClientID in spotifyAuth/spotifyAuth.go (line 17,18), as reading from input doesn't work yet
- playlist must be public, won't work with private

TODO
- add nicer terminal "gui", maybe this: https://github.com/rivo/tview - not really needed
- add deleting from yt playlist when appending to spotify playlist - done
- fix error with reading SpotifyClientID and SpotifyClientSecret from flags