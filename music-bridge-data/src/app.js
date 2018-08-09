if (process.env.NODE_ENV !== 'production') {
  require('dotenv').config();
}
const express = require('express');
const graphqlHTTP = require('express-graphql');
const cors = require('cors');

const schema = require('./graphql/schema');

const app = express();

app.set('trust proxy', true);

app.get('/', (req, res) => res.send('Hello GraphQL 🎉'));

app.use(
  '/graphql',
  cors(),
  graphqlHTTP({
    graphiql: true,
    schema,
  }),
);

app.listen(process.env.PORT, () => {
  console.log(`App listening on port ${process.env.PORT}`);
});
