const axios = require('axios');
const moment = require('moment');
const spotifyBridgeUrl = process.env.SPOTIFY_BRIDGE_URL;

const errorPassThrough = exc => {
  //console.error(exc);
  throw exc;
};
module.exports = {
  Query: {
    spotifyAuthUrl: async (root, {redirect}, context, info) => {
      try {
        const {data: url} = await axios.post(`${spotifyBridgeUrl}/auth-url`, {
          redirect,
        });
        return url;
      } catch (exc) {
        errorPassThrough(exc);
      }
    },
    spotifyAuth: async (root, {code, redirect}, context, info) => {
      try {
        const {data: auth} = await axios.get(
          `${spotifyBridgeUrl}/auth-callback?code=${code}&redirect=${redirect}`,
          {
            'Content-Type': 'application/x-www-form-urlencoded',
          },
        );
        console.log(auth);
        return {
          accessToken: auth.access_token,
          expiry: moment()
            .add(auth.expires_in, 'seconds')
            .toISOString(),
        };
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
