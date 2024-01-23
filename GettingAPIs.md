## Youtube
- Go to the Google Cloud Console: Visit the Google Cloud Console at https://console.cloud.google.com/ and sign in with your Google account.

- Create a new project: In the Google Cloud Console, click on the project drop-down menu at the top of the page and select "New Project". Enter a name for your project and click on the "Create" button. 

- Enable the YouTube Data API: Once your project is created, you need to enable the YouTube Data API for your project.

- In the Google Cloud Console, click on the navigation menu (☰) in the upper-left corner.
Scroll down and select "APIs & Services" > "Library" from the menu.
In the search bar, type "YouTube Data API" and select it from the results.
Click on the "Enable" button to enable the API for your project.
Set up API credentials: To authenticate your application and make authorized API calls, you need to create API credentials.

- In the Google Cloud Console, go to "APIs & Services" > "Credentials".
Click on the "Create Credentials" button and select "OAuth client ID".
Choose the application type based on your requirements (e.g., Web application, Desktop application, etc.).
Enter a name for the OAuth 2.0 client ID and configure the authorized JavaScript origins and redirect URIs if applicable.
- Click on the "Create" button to generate the OAuth client ID and client secret.
Note down the API credentials: After creating the OAuth client ID, you will see a dialog box displaying the client ID and client secret. Take note of these credentials as you will need them in your application for authentication.
- go to oauth credentials and click "Authorised redirect URIs" and set it as localhost
- add yourself as tester in "Credentials" section

## Spotify
- Go to the Spotify Developer Dashboard: Visit the Spotify Developer Dashboard at https://developer.spotify.com/dashboard/ and sign in with your Spotify account or create a new account if you don't have one.

- Create a new application: Once you are logged in to the Spotify Developer Dashboard, click on the "Create an App" button to create a new application. Use "http://localhost:8080/spotify" as Redirect URI 


## Getting youtube playlist ID:
Go to your target YouTube playlist on the browser. On the address bar, you will see something like this: https://www.youtube.com/watch?v=RLykC1VN7NY&list=PLFs4vir_WsTwEd-nJgVJCZPNL3HALHHpF. The playlist ID is the characters after “list=” so in the URL above, our playlist ID is PLFs4vir_WsTwEd-nJgVJCZPNL3HALHHpF.