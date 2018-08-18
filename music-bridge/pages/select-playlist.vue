<template>
    <div class="container">
        <div class="row full-height align-items-top">
            <div class="col">
                <div class="row">
                    <div class="col">
                        <h1 class="text-center">Transfer your Playlists</h1>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                        <p class="text-center lead">Pick the playlist</p>
                    </div>
                    <div class="col">
                        <p class="text-center lead">Check the songs</p>
                    </div>
                    <div class="col" v-if='chosenSongs.length'>
                        <p class="text-center lead">Confirm YourSelection</p>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                        <b-form-select :options="playlistTitles" v-model="selectedPlaylist" id="playlistSelect">

                             <template slot="first">
                                     this slot appears above the options from 'options' prop
                                    <option :value="false" disabled>-- Please select an option --</option>
                                  </template>
                            </b-form-select>

                            <p class="text-center" v-if="selectedPlaylist !== false">
                            <ul>
                              <li v-for="(song, index) in computedSongs" v-bind:key="index">
                                {{ song.artist}} - {{song.name}}
                              </li>
                            </ul>
                            </p>
                            <b-button @click='makeSearch' v-if="selectedPlaylist">
                                Transfer Playlist
                            </b-button>
                    </div>
                    <div class="col" >
                        <div role="tablist" class="shadow-lg">
                            <b-card v-for="(song, index) in searchResults" v-bind:key="index" no-body class="mb-1">
                                <b-card-header header-tag="header" class="p-1" role="tab">
                                    <b-btn block  v-b-toggle="getAccordionID('accordion', index)" variant="info">{{ song.results[0].title }} - {{ song.results[0].artist }}</b-btn>
                                </b-card-header>
                            <div v-for="(video, i) in song.results" v-bind:key="i" no-body class="mb-1">
                              <LazyCollapse :url="getVideoUrl(video.videoId)" :id="getAccordionID('accordion', index)">
                                  <b-button @click='useMe(index, i)'>Select this</b-button>
                              </LazyCollapse>
                            </div>
                           </b-card>
                        </div>
                    </div>
                    <div class="col" v-if='chosenSongs.length'>
                            <ul>
                              <li v-for="(song, index) in chosenSongs" v-bind:key="index">
                                {{ song.title}} - {{song.videoId}}
                              </li>
                            </ul>
                            <b-button @click='transferPlaylist' >
                                Make it so!
                            </b-button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import LazyCollapse from "~/components/LazyCollapse";
export default {
  components: {
    LazyCollapse
  },
  mounted() {
    this.$apollo.addSmartQuery("playlists", {
      query: require("~/graphql/SpotifyPlaylists.gql"),
      update: ({ spotifyPlaylists }) => spotifyPlaylists,
      variables: {
        accessToken: this.$store.state.spotify.accessToken
      }
    });
  },
  methods: {
    useMe(searchResultsIndex, resultIndex) {
      this.chosenSongs.push(
        this.searchResults[searchResultsIndex].results[resultIndex]
      );
      this.searchResults.splice(searchResultsIndex, 1);
    },
    getAccordionID(prefix, index) {
      return `${prefix}-${index}`;
    },
    getVideoUrl(videoId) {
      return `https://www.youtube.com/embed/${videoId}`;
    },
    async transferPlaylist() {
      const createPlaylist = ({ name: title, public: privacy }) => {
        return this.$apollo
          .query({
            query: require("~/graphql/CreatePlaylist.gql"),
            fetchPolicy: "network-only",
            variables: {
              title,
              privacyStatus: privacy ? "public" : "private",
              accessToken: this.$store.state.youtube.accessToken
            }
          })
          .then(({ data: { youtubeCreatePlaylist } }) => youtubeCreatePlaylist);
      };
      const addToPlaylist = (playlistId, videoId) => {
        return this.$apollo.query({
          query: require("~/graphql/AddToPlaylist.gql"),
          fetchPolicy: "network-only",
          variables: {
            playlistId,
            videoId,
            accessToken: this.$store.state.youtube.accessToken
          }
        });
      };
      const playlist = await createPlaylist(
        this.playlists[this.selectedPlaylist]
      );
      this.chosenSongs.map(async ({ videoId }) => {
        const data = await addToPlaylist(playlist.id, videoId);
        console.log(data);
      });
    },
    makeSearch() {
      this.$apollo
        .mutate({
          mutation: require("~/graphql/SearchSongs.gql"),
          variables: {
            songs: this.selectedSongs().map(({ artist, name }) => {
              return { artist, title: name };
            }),
            accessToken: this.$store.state.youtube.accessToken
          }
        })
        .then(
          ({ data: { youtubeSearchSongs } }) =>
            (this.searchResults = youtubeSearchSongs)
        );
    },
    selectedSongs() {
      if (!this.playlists.length) return [];
      const tracks = this.playlists.filter((_, index) => {
        return index === this.selectedPlaylist;
      })[0].tracks;
      return tracks;
    }
  },
  computed: {
    computedSongs() {
      return this.selectedSongs();
    },
    playlistTitles() {
      if (!this.playlists.length) return [];

      return this.playlists.map((item, index) => ({
        value: index,
        text: `${item.name} (${item.tracks.length})`
      }));
    }
  },
  data() {
    return {
      selectedPlaylist: false,
      searchResults: [],
      chosenSongs: [],
      playlists: []
    };
  }
};
</script>

<style>
</style>
