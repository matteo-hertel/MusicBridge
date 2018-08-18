<template>

    <p v-if="textDependency" class="text-center lead">{{ showIfTrue }}</p>
    <p v-else class="text-center lead"> {{ showIfFalse }} </p>

</template>

<script>
    export default {
        props: ["textDependency", "showIfTrue", "showIfFalse"],
    };
</script>