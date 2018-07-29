<template>
<div>
Youtube Auth Url: {{this.$store.state.youtube.authUrl}}
</div>
</template>

<script>
export default {
  mounted() {
    this.$store.dispatch("youtube/fetchAuthUrl");
  }
};
</script>

<style>
.done {
  text-decoration: line-through;
}
</style>
