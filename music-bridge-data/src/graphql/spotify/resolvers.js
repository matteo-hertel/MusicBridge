const axios = require('axios');
const spotifyBridgeUrl = process.env.SPOTIFY_BRIDGE_URL;

const errorPassThrough = exc => {
  //console.error(exc);
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
  },
  Mutation: {},
};
