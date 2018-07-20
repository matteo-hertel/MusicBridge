const axios = require('axios');
const spotifyBridgeUrl = process.env.SPOTIFY_BRIDGE_URL;
const youtubeBridgeUrl = process.env.YOUTUBE_BRIDGE_URL;

const errorPassThrough = exc => {
  console.error(exc);
  throw exc;
};
module.exports = {
  Query: {
    spotifyAuthUrl: async (root, _, context, info) => {
      try {
        const {data: url} = await axios.get(`${spotifyBridgeUrl}/auth-url`);
        return url;
      } catch (exc) {
        errorPassThrough(exc);
      }
    },
    spotifyAuth: async (root, {code}, context, info) => {
      try {
        const {data: userInfo} = await axios.get(
          `${spotifyBridgeUrl}/auth-callback?code=${code}`,
        );
        return userInfo;
      } catch (exc) {
        errorPassThrough(exc);
      }
    },

    spotifyPlaylists: async (root, {accessToken}, context, info) => {
      try {
        const {data: playlists} = await axios.get(
          `${spotifyBridgeUrl}/playlists`,
          {
            headers: {
              'X-Spotify-Token': accessToken,
            },
          },
        );
        return playlists;
      } catch (exc) {
        errorPassThrough(exc);
      }
    },
    youtubeAuthUrl: async (root, _, context, info) => {
      try {
        const {data: url} = await axios.get(`${youtubeBridgeUrl}/auth-url`);
        return url;
      } catch (exc) {
        errorPassThrough(exc);
      }
    },
    youtubeCreatePlaylist: async (
      root,
      {accessToken, title, description, privacyStatus},
      context,
      info,
    ) => {
      try {
        const data = await axios.post(
          `${youtubeBridgeUrl}/create-playlist`,
          {
            title,
            description,
            privacyStatus,
          },
          {
            headers: {
              'X-Youtube-Token': accessToken,
            },
          },
        );
        console.log(data);
        return JSON.stringify(data);
      } catch (exc) {
        errorPassThrough(exc);
      }
    },
  },
};
