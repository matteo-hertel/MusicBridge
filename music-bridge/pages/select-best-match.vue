<template>
    <div class="container">
      <div class="full-height">
        <div class="row  align-items-top">
            <div class="col">
                <div class="row">
                    <div class="col">
                        <h1 class="text-center">Select best match</h1>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                        <p class="text-center lead">We'll show 5 alternatives for each song, choose the best one!</p>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                      <b-container >
                          <p class="lead">Playlist Info</p>
                          <b-row>
                              <b-col>Name: </b-col>
                              <b-col>{{playlist.name}}</b-col>
                          </b-row>
                          <b-row>
                              <b-col>Public: </b-col>
                              <b-col>{{playlist.public}}</b-col>
                          </b-row>
                          <b-row>
                              <b-col>Tracks: </b-col>
                              <b-col>{{playlist.tracks.length}}</b-col>
                          </b-row>
                      </b-container>

                      <b-container>
                          <b-row>
                             <b-list-group>
                              <b-list-group-item v-for="(song, index) in playlist.tracks" v-bind:key="index">
                                {{ song.artist}} - {{song.name}}
                              </b-list-group-item>
                              </b-list-group>

                              <b-col v-if='!searchResults.length'>
                                   <p class='lead'>Finding the best matches, maybe we should start a dating app 🤔</p>
                              </b-col>
                              <b-col v-else>
                                    {{searchResults}}
                              </b-col>
                          </b-row>
                      </b-container>
                    </div>
                </div>
            </div>
        </div>
        </div>
    </div>
</template>

<script>
export default {
  mounted() {
    this.makeSearch();
  },
  methods: {
    async makeSearch() {
      return this.$apollo
        .mutate({
          mutation: require("~/graphql/SearchSongs.gql"),
          variables: {
            songs: this.$store.state.core.playlist.tracks.map(
              ({ artist, name }) => {
                return { artist, title: name };
              }
            ),
            accessToken: this.$store.state.youtube.accessToken
          }
        })
        .then(
          ({ data: { youtubeSearchSongs } }) =>
            (this.searchResults = youtubeSearchSongs)
        );
    }
  },
  computed: {
    playlist() {
      return this.$store.state.core.playlist || {};
    }
  },
  middleware: "authenticated",
  data() {
    return {
      searchResults: []
    };
  }
};
</script>

<style>
</style>
