<template>
  <div class="container">
    <router-link to="/">Dashboard</router-link>
    <router-link to="/game">Game</router-link>
    <router-link to="/score-board">Score Board</router-link>
    <router-link to="/about">About</router-link>
  </div>
</template>
<script>
export default {};
</script>
<style scoped>
.container {
  display: flex;
  justify-content: space-evenly;
}
</style>