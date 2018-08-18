<template>
 <b-collapse :id="id" @show="show" @hide="hide"   accordion="my-accordion" role="tabpanel">
<div v-if="isVisible">
   <b-card-body >
         <p class="card-text">
<slot></slot>
 <b-form-checkbox-group  :v-model="selected">
      </b-form-checkbox-group>
             <b-embed type="iframe" :src='url' aspect="16by9" controls></b-embed>
         </p>
     </b-card-body>
</div>
 </b-collapse>
</template>

<script>
export default {
  props: ["id", "url", "selected"],
  methods: {
    hide() {
      this.isVisible = false;
    },
    show() {
      this.isVisible = true;
    }
  },
  data() {
    return {
      isVisible: false
    };
  }
};
</script>
