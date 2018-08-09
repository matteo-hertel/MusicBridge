<template>
    <div class="container">
        <div class="row full-height align-items-center">
            <div class="col">
                <div class="row">
                    <div class="col">
                        <h1 class="text-center"><span class="youtube-pulse youtube-text">YouTube</span></h1>
                    </div>
                </div>
                <div class="row">
                    <div class="col">
                        <p class="text-center lead">Next, let's log in to your YouTube account.</p>
                        <p class="text-center">
                            <YoutubeLoginButton />
                            <nuxt-link to="select-playlist" class="btn btn-primary btn-lg">
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
    import YoutubeLoginButton from "~/components/YoutubeLoginButton.vue";

    export default {
      components: {
        YoutubeLoginButton,
      }
    };
</script>

<style>

</style>
