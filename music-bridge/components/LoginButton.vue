<template>
    <b-button
            :disabled="!url"
            :href="url"
            v-if="buttonDependency"
            size="lg"
            variant="primary"
    >
        <span v-if="url">{{ buttonMessage }}</span>
        <span v-else>{{ waitMessage }}</span>
    </b-button>
</template>

<script>
    export default {
        props: ["url", "waitMessage", "buttonMessage", "buttonDependency"],
    };
</script>