<template>
    <div class="container">
        <div class="row full-height align-items-center">
            <div class="col">
                <div class="row">
                    <div class="col">
                        <h1 class="text-center">We're checking with <span class="youtube-pulse youtube-text">YouTube</span> if it's all good</h1>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                        <p class="text-center lead">Hang on tight!</p>
                        <p class="text-center">
<nuxt-link to="select-playlist" :disabled="!this.$store.state.youtube.accessToken"


class="btn btn-primary btn-lg">
                        Let's goo!
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
      "youtube/getAccessTokenFromUrl",
      this.$route.query.code
    );
  }
};
</script>

<style>
</style>
