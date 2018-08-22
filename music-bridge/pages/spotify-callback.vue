<template>
            <div class="col">
                <div class="row">
                    <div class="col">
                        <h1 class="text-center"><span class="spotify-pulse pulse-repeat spotify-text">Loading</span></h1>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                        <ConditionalBlock
                                :condition="spotifyAccessToken"
                        >
                            <div slot="true">
                                <p class="text-center lead">Success! Redirecting you now...</p>
                            </div>

                            <div slot="false">
                                <p class="text-center lead">This will only take a few moments...</p>
                            </div>

                        </ConditionalBlock>
                    </div>
                </div>
            </div>
</template>

<script>

    import ConditionalBlock from "~/components/ConditionalBlock.vue";

export default {
  mounted() {
    this.$store.dispatch(
      "spotify/getAccessTokenFromUrl",
      this.$route.query.code
    );
  },
  computed: {
      spotifyAccessToken: function() {
          return this.$store.state.spotify.accessToken;
      }
  },
  watch: {
        spotifyAccessToken: function (newToken, oldToken) {
            if (!oldToken && newToken) {
                setTimeout(()=>{
                    this.$router.push({'path': '/link-youtube'});
                }, 2000);
            }
        }
    },
  components: {
        ConditionalBlock
    }
};
</script>

<style>
</style>
