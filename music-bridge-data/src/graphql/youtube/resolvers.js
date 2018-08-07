const axios = require('axios');
const youtubeBridgeUrl = process.env.YOUTUBE_BRIDGE_URL;

const errorPassThrough = exc => {
  //console.error(exc);
  throw exc;
};

module.exports = {
  Query: {
    youtubeAuthUrl: async (root, {redirect}, context, info) => {
      try {
        const {data: url} = await axios.post(`${youtubeBridgeUrl}/auth-url`, {
          redirect,
        });
        return url;
      } catch (exc) {
        errorPassThrough(exc);
      }
    },
    youtubeAuth: async (root, {code}, context, info) => {
      try {
        const {data} = await axios.get(
          `${youtubeBridgeUrl}/auth-callback?code=${code}`,
        );
        return data;
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
          `${youtubeBridgeUrl}/bulk-search`,
          songs,
          {
            headers: {
              'X-Youtube-Token': accessToken,
            },
          },
        );
        return data.map(mapToProp('items', extractSongInfo));
      } catch (exc) {
        errorPassThrough(exc);
      }
    },
  },
};
function extractSongInfo(song) {
  const {
    id: {videoId},
    snippet: {description, title},
  } = song;
  return {videoId, description, title};
}
function mapToProp(prop, fn) {
  return function map(obj) {
    return {results: obj[prop].map(fn)};
  };
}
