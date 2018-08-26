<template>
    <div class="col">
        <div class="row">
            <div class="col">
                <h1 class="text-center"><span class="youtube-pulse pulse-repeat youtube-text">Loading</span></h1>
            </div>
        </div>
        <div class="row">
            <div class="col">
                <ConditionalBlock
                        :condition="youtubeAccessToken"
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
      "youtube/getAccessTokenFromUrl",
      this.$route.query.code
    );
  },
    computed: {
        youtubeAccessToken: function() {
            return this.$store.state.youtube.accessToken;
        }
    },
    watch: {
        youtubeAccessToken: function (newToken, oldToken) {
            if (!oldToken && newToken) {
                setTimeout(()=>{
                    this.$router.push({'path': '/select-playlist'});
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
