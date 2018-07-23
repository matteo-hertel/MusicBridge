const typeDefs = `
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

`;
const queryDefs = `
  youtubeAuthUrl: AuthUrl!
  youtubeCreatePlaylist(accessToken: String!, title: String!, description: String, privacyStatus: String): YoutubeCreatePlaylist!
  youtubeAddToPlaylist(accessToken: String!, playlistId: String!, videoId: String!, position: String ): YoutubeVideo!
`;
module.exports = {
  typeDefs,
  queryDefs,
};
