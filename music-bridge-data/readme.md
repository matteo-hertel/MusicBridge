## MusicBridgeData 
GraphQL layer to aggregate data for the music bridge 

### Setup
For local development make sure you have a `.env` file with the same vars as in `app_env.yaml.example`(and actual values in 😁)

```
npm install
```

### Run the Application
```
npm start
```

### Development
```
npm run dev
```

### Deployment
It's super easy to deploy to google app engine, rename `app_env.yaml.example` to  `app_env.yaml` and replace the placeholders with your env vars then

```
gcloud app deploy
```

and you're all sorted

