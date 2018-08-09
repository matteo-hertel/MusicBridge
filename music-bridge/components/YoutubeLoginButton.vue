<template>
    <b-button
            :disabled="!this.$store.state.youtube.authUrl"
            :href="this.$store.state.youtube.authUrl"
            size="lg"
            variant="primary"
    >
        <span>{{ this.$store.state.youtube.test }}</span>
        <span v-if="this.$store.state.youtube.authUrl">Log in to YouTube</span>
        <span v-else>Just a second...</span>
    </b-button>
</template>
