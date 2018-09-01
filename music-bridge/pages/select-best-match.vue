<template>
    <div class="container">
      <div class="full-height">
        <div class="row  align-items-top">
            <b-button @click="updatePointer(-1)">Previous</b-button>
            <div class="col">
                <div class="row">
                    <div class="col">
                        <h1 class="text-center">Select the best matches</h1>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                        <p class="text-center lead">We'll show 5 alternatives for each song, choose the best one!</p>
                    </div>
                </div>
                <b-row v-if='!searchResults.length'>
                    <b-col>
                        <p class='lead text-center'>Finding your best matches, maybe we should start a dating app 🤔</p>
                    </b-col>
                </b-row>
                <b-row v-if='searchResults.length'>
                    <b-col>
                        <b-row>
                            <b-col v-for="(result, index) in currentItem" v-bind:key="index">
                                <p>{{ result.title }}</p>
                            </b-col>
                        </b-row>
                        <b-row>
                            <b-col v-for="(result, index) in currentItem" v-bind:key="index">
                                <b-embed
                                        type="iframe"
                                        aspect="16by9"
                                        :src=get_youtube_video_source(result.videoId)
                                        allowfullscreen
                                ></b-embed>
                                <b-button block>Choose</b-button>
                            </b-col>
                        </b-row>
                    </b-col>
                </b-row>
                <b-row>
                    <b-col>
                        <b-progress :value="this.resultPointer" :max="this.totalResults" class="mb-3"></b-progress>
                    </b-col>
                </b-row>

            </div>
            <b-button @click="updatePointer(1)">Next</b-button>
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
      get_youtube_video_source(videoId) {
          return 'https://www.youtube.com/embed/' + videoId + '?rel=0'
      },
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
          {
              this.searchResults = youtubeSearchSongs;
              this.totalResults = youtubeSearchSongs.length;
          },

        );
    },
      updatePointer(operation) {
          const updatedPointer = this.resultPointer + operation;
          if (updatedPointer < 0) {
              this.resultPointer = 0;
              return;
          }

          if (updatedPointer > this.totalResults -1) {
              this.resultPointer = this.totalResults -1;
              return;
          }

          this.resultPointer = updatedPointer;
      }
  },
  computed: {
    playlist() {
      return this.$store.state.core.playlist || {};
    },
    currentItem() {
        return this.searchResults[this.resultPointer].results;
    },
  },
  middleware: "authenticated",
  data() {
    return {
      searchResults: [],
      chosenResults: [],
      resultPointer: 0,
      totalResults: 0
    };
  }
};
</script>

<style>
</style>
