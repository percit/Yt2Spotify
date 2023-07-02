import os
import googleapiclient.discovery
from google.oauth2 import service_account

# Set up the YouTube Data API credentials
credentials = service_account.Credentials.from_service_account_file(
    'path/to/service_account_credentials.json',
    scopes=['https://www.googleapis.com/auth/youtube.readonly']
)
youtube = googleapiclient.discovery.build('youtube', 'v3', credentials=credentials)

# Define the playlist ID
playlist_id = 'YOUR_PLAYLIST_ID'

# Get the playlist items
items = []
next_page_token = None

while True:
    request = youtube.playlistItems().list(
        part='snippet',
        playlistId=playlist_id,
        maxResults=50,
        pageToken=next_page_token
    )
    response = request.execute()

    items.extend(response['items'])
    next_page_token = response.get('nextPageToken')

    if not next_page_token:
        break

# Extract video titles
video_titles = [item['snippet']['title'] for item in items]

# Write video titles to a text file
output_file = 'playlist_titles.txt'

with open(output_file, 'w') as f:
    f.write('\n'.join(video_titles))

print(f"Playlist titles written to {os.path.abspath(output_file)}")