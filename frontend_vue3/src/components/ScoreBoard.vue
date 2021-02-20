<template>
  <div>
    <h1>{{ msg }}</h1>
    
  </div>
</template>

<script>

export default {
  name: "score-board",
  props: {
    // msg: String,
  },
  data() {
    return {
      msg: "Score Board"
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
