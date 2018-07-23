const axios = require('axios');
const youtubeBridgeUrl = process.env.YOUTUBE_BRIDGE_URL;

const errorPassThrough = exc => {
  //console.error(exc);
  throw exc;
};
module.exports = {
  Query: {
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
        const {data} = await axios.post(
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
        return data;
      } catch (exc) {
        errorPassThrough(exc);
      }
    },
    youtubeAddToPlaylist: async (
      root,
      {accessToken, playlistId, videoId, position},
      context,
      info,
    ) => {
      try {
        const {data} = await axios.post(
          `${youtubeBridgeUrl}/add-to-playlist`,
          {
            playlistId,
            videoId,
            position,
          },
          {
            headers: {
              'X-Youtube-Token': accessToken,
            },
          },
        );
        return data;
      } catch (exc) {
        errorPassThrough(exc);
      }
    },
  },
  Mutation: {
    youtubeSearchSongs: async (root, {accessToken, songs}, context, info) => {
      try {
        const {data} = await axios.post(
          `${youtubeBridgeUrl}/add-to-playlist`,
          [...songs],
          {
            headers: {
              'X-Youtube-Token': accessToken,
            },
          },
        );
        return data;
      } catch (exc) {
        errorPassThrough(exc);
      }
    },
  },
};
