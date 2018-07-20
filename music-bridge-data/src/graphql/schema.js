const {makeExecutableSchema} = require('graphql-tools');
const resolvers = require('./resolvers');

const typeDefs = `
type SpotifyUserInfo{
  accessToken: String!,
  display_name: String
}

type SpotifyPlaylist{
  id: ID!,
  name: String!,
  owner: String!,
  public: Boolean,
  tracks: [SpotifyTrack]
}

type SpotifyTrack{
  name: String!,
  artist: String!
}

type YoutubePlaylistSnippet{
  channelId: String,
channelTitle: String,
title: String!
}

type YoutubePlaylistStatus{
  privacyStatus: String!
}

type YoutubeCreatePlaylist{
  id: String!,
  kind:  String,
  snippet: YoutubePlaylistSnippet!
  status: YoutubePlaylistStatus  
}

type AuthUrl{
  url: String!
}

type YoutubeVideoSnippet{
  title: String!
}

type YoutubeVideo{
  kind:  String,
  snippet: YoutubeVideoSnippet!
}

type Query{
  spotifyPlaylists(accessToken: String!): [SpotifyPlaylist],
  spotifyAuthUrl: AuthUrl!
  spotifyAuth(code: String!): SpotifyUserInfo

  youtubeAuthUrl: AuthUrl!
  youtubeCreatePlaylist(accessToken: String!, title: String!, description: String, privacyStatus: String): YoutubeCreatePlaylist!
  youtubeAddToPlaylist(accessToken: String!, playlistId: String!, videoId: String!, position: String ): YoutubeVideo!
}
`;
module.exports = makeExecutableSchema({
  typeDefs,
  resolvers,
});
