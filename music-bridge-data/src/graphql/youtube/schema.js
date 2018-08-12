const typeDefs = `
type YoutubeUserAuth{
  accessToken: String!
  expiry: String!
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

type YoutubeSearchResultWrapper{
results: [YoutubeSearchResult],
}

type YoutubeSearchResult{
videoId: String,
description: String,
title: String
}
input YoutubeSongInput{
  artist: String!,
  title: String!
}

type Mutation {
  youtubeSearchSongs(accessToken: String!, songs: [YoutubeSongInput]): [YoutubeSearchResultWrapper]
}
`;
const queryDefs = `
  youtubeAuthUrl(redirect: String): AuthUrl!
  youtubeAuth(code: String!, redirect: String): YoutubeUserAuth
  youtubeCreatePlaylist(accessToken: String!, title: String!, description: String, privacyStatus: String): YoutubeCreatePlaylist!
  youtubeAddToPlaylist(accessToken: String!, playlistId: String!, videoId: String!, position: String ): YoutubeVideo!
`;
module.exports = {
  typeDefs,
  queryDefs,
};
