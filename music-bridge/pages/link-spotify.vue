<template>
    <div class="container">
        <div class="row full-height align-items-center">
            <div class="col">
                <div class="row">
                    <div class="col">
                        <h1 class="text-center"><span class="spotify-pulse spotify-text">Spotify</span></h1>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                        <p class="text-center lead">First, let's log in to your Spotify account.</p>
                        <p class="text-center">
<!--Move this to its own function -->
                            <b-button :href="spotifyUrl" class="btn btn-primary btn-lg">
                                Log in to Spotify
                            </b-button>
                            <nuxt-link to="link-youtube" class="btn btn-primary btn-lg">
                                Next
                            </nuxt-link>
                        </p>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
  computed: {
    spotifyUrl: function() {
      return this.$store.state.spotify.authUrl;
    }
  }
};
</script>

<style>
</style>
