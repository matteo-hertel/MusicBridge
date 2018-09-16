<template>
  <div class="col">
    <div class="row">
      <div class="col">
        <h1 class="text-center">Music Bridge</h1>
      </div>
    </div>
    <div class="row">
    <div class="col">
      <p class="text-center lead">Success!🎉🎉🎉</p>
      <p class="text-center">
        Your playlist has been transfered to YouTube, here's the link to see it
      </p>
      <p class="text-center">
       <b-link :href='playlistUrl' target='_blank' >{{playlistUrl}}</b-link>
      </p>
      <p class="text-center">
       <b-button variant="primary" @click="goAgain">Go Again!</b-button>
      </p>
    </div>
  </div>
  </div>
</template>

<script>
export default {
  methods: {
    async goAgain() {
      await this.$store.commit("core/wipeState");
      const redirectUrl = this.$store.getters["core/stepUrl"];
      this.$router.push({ path: redirectUrl });
    }
  },
  computed: {
    playlistUrl() {
      return `https://www.youtube.com/playlist?list=${
        this.$store.state.core.createdPlaylistId
      }`;
    }
  }
};
</script>
