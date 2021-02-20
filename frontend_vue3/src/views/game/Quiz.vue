<template>
  <div class="container">
    <div class="hud">
      <div class="score">100 points</div>
      <div class="player">player</div>
    </div>
    <div class="audio">audio</div>
    <div class="question">question</div>
    <div class="image">
      <!-- <img src="http://localhost:8080/quizzes/123/image"/> -->
    </div>
    <div class="answer">answer</div>
  </div>
</template>
<script>
export default {};
</script>
<style scoped>
.hud {
  display: flex;
  flex-direction: row;
  justify-content:space-evenly;
  margin-top: 20px;
}
</style>