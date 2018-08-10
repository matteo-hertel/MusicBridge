<template>
    <div class="container">
        <div class="row full-height align-items-center">
            <div class="col">
                <div class="row">
                    <div class="col">
                        <h1 class="text-center">We're checking with <span class="spotify-pulse spotify-text">Spotify</span> if it's all good</h1>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                        <p class="text-center lead">Hang on tight!</p>
                        <p class="text-center">
           <nuxt-link :disabled="!this.$store.state.youtube.accessToken"
                      to="link-youtube" class="btn btn-primary btn-lg">
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
  mounted() {
    this.$store.dispatch(
      "spotify/getAccessTokenFromUrl",
      this.$route.query.code
    );
  }
};
</script>

<style>
</style>
