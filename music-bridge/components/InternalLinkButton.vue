<template>

    <nuxt-link v-if="buttonDependency" :to="linkTo" class="btn btn-primary btn-lg">
        {{ buttonMessage }}
    </nuxt-link>

</template>

<script>
    export default {
        props: ["linkTo", "buttonDependency", "buttonMessage"],
    };
</script>