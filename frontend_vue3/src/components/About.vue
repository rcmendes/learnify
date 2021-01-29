<template>
  <div>
    <h1>{{ msg }}</h1>
    
  </div>
</template>

<script>

export default {
  name: "about",
  props: {
    // msg: String,
  },
  data() {
    return {
      msg: "About"
    };
  },
  methods: {
  },
  mounted() {
  }
};
</script>
<style scoped>
</style>
