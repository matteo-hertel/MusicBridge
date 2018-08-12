const _ = require('lodash');
const SpotifyWebApi = require('spotify-web-api-node');

const scopes = ['user-read-private', 'user-read-email'];

function spotifyApiProvider(spotifyApi) {
  async function getAllPlaylists() {
    async function enrichData(playlist) {
      playlist.tracks = await getPlaylistTracks(playlist.owner, playlist.id);
      return Promise.resolve(playlist);
    }
    const playlists = await getUserPlaylists();
    return await Promise.all(playlists.map(enrichData));
  }

  function setAccessToken(accessToken) {
    return spotifyApi.setAccessToken(accessToken);
  }

  function getAuthorizeURL() {
    return spotifyApi.createAuthorizeURL(scopes);
  }
  async function getUserInfo() {
    return await spotifyApi.getMe();
  }
  async function getApiToken(CALLBACK_TOKEN) {
    const auth = await spotifyApi.authorizationCodeGrant(CALLBACK_TOKEN);
    return auth.body;
  }

  function getUserPlaylists() {
    function processUserData(data) {
      return data.body.items.map(i => {
        const {
          id,
          name,
          public,
          owner: {id: owner},
        } = i;
        return {id, name, owner, public};
      });
    }
    return spotifyApi.getUserPlaylists().then(processUserData);
  }

  function getPlaylistTracks(user, id) {
    function processPlaylistData(data) {
      return data.body.tracks.items.map(i => ({
        name: _.get(i, 'track.name'),
        artist: _.get(i, 'track.artists[0].name'),
      }));
    }
    return spotifyApi.getPlaylist(user, id).then(processPlaylistData);
  }
  return {
    getAllPlaylists,
    getApiToken,
    getAuthorizeURL,
    getUserInfo,
    setAccessToken,
    spotifyApi,
  };
}
function makeSpotifyApi() {
  return new SpotifyWebApi({
    clientId: process.env.CLIENT_ID,
    clientSecret: process.env.CLIENT_SECRET,
    redirectUri: process.env.REDIRECT_URI,
  });
}
module.exports = {
  makeSpotifyApi,
  spotifyApiProvider,
};
