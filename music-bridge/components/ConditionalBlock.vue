<template>

        <div>
            <slot v-if="condition" name="true"></slot>
            <slot v-else name="false"></slot>
        </div>

</template>

<script>
    export default {
        props: ["condition"],
    };
</script>