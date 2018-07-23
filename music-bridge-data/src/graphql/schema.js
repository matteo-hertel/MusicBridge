const {makeExecutableSchema} = require('graphql-tools');
const youtube = require('./youtube');
const spotify = require('./spotify');

const queryDefs = `
type Query{
  ${spotify.schema.queryDefs}
  ${youtube.schema.queryDefs}
}
`;
module.exports = makeExecutableSchema({
  typeDefs: [spotify.schema.typeDefs, youtube.schema.typeDefs, queryDefs],
  resolvers: {
    Query: {
      ...spotify.resolvers,
      ...youtube.resolvers,
    },
  },
});
