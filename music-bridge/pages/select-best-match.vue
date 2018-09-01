<template>
    <div class="container">
      <div class="full-height">
        <div class="row  align-items-top">
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
                                <b-button @click="chooseResult(result)" block>Choose</b-button>
                            </b-col>
                        </b-row>
                    </b-col>
                </b-row>
                <b-row>
                    <b-col>
                        <hr>
                    </b-col>
                </b-row>
                <b-row class="control-row">
                    <b-col>
                        <b-row>
                            <b-col>
                                <b-button :disabled="!this.resultPointer" block @click="updatePointer(-1)">prev</b-button>
                            </b-col>
                            <b-col cols="8">
                                <b-progress height="100%" :max="this.totalResults">
                                    <b-progress-bar :value="this.resultPointer+1" show-value :label="progressLabel">
                                    </b-progress-bar>
                                </b-progress>
                            </b-col>
                            <b-col>
                                <b-button :disabled="this.resultPointer+1 == this.totalResults" block @click="updatePointer(1)">next</b-button>
                            </b-col>
                        </b-row>
                        <b-row v-if="this.resultPointer+1 == this.totalResults">
                            <b-col class="align-right">
                                <h2 class="text-right">All done?</h2>

                                <ConditionalBlock
                                        :condition="hasChosenAll"
                                >
                                    <div slot="true">
                                        <p class="text-right">Feel free to go back and change your choices.</p>
                                    </div>
                                    <div slot="false">
                                        <p class="text-right">You haven't chosen every song yet, you can go back and select it.</p>
                                    </div>
                                </ConditionalBlock>

                                <b-button :disabled="hasChosen" @click="initiateTransfer" class="float-right">Transfer my Playlist</b-button>
                            </b-col>
                        </b-row>
                    </b-col>
                </b-row>

            </div>
        </div>
        </div>
    </div>
</template>

<script>
    import Vue from 'vue';
    import ConditionalBlock from "~/components/ConditionalBlock.vue";

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
      },
      chooseResult(result) {
          Vue.set(this.chosenResults, this.resultPointer, result);
      },
      async initiateTransfer() {
          this.$store.dispatch('core/storeSongs', this.chosenResults);
          await this.$store.commit("core/incrementStep");
          const redirectUrl = this.$store.getters["core/stepUrl"];
          this.$router.push({ path: redirectUrl });
      }
  },
  computed: {
    playlist() {
      return this.$store.state.core.playlist || {};
    },
    currentItem() {
        return this.searchResults[this.resultPointer].results;
    },
      progressLabel() {
        return `${this.resultPointer+1} / ${this.totalResults}`
      },
      hasChosen() {
          return !this.chosenResults.length;
      },
      hasChosenAll() {
          return (this.chosenResults.filter(result => !!result).length === this.totalResults);
      }
  },
  middleware: "authenticated",
  data() {
    return {
      searchResults: [],
      chosenResults: [],
      resultPointer: 0,
      totalResults: 0
    };
  },
    components: {
        ConditionalBlock
    }
};
</script>

<style lang="scss">
    hr {
        border-top: 2px solid white;
    }
</style>
