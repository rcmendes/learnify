<template>
  <div class="container">
    <div class="title-bar">
      <h1>{{ category }}</h1>
    </div>
    <div class="quiz-list">
      <div
        class="quiz-list-item"
        v-for="quiz in quizzes"
        :key="quiz.uuid"
        @click="() => onSelectQuiz(quiz)"
      >
        {{quiz.palavra}}
        <img :src="quiz.image_url"/>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  components: {},
  data() {
    return {
      category: "Animais",
      quizzes: [],
    };
  },
  methods: {
    onSelectQuiz(quiz) {
      console.log("Selected quiz:", quiz.name);
      this.$router.push("/quiz");
    },
  },
  mounted() {
    const url = `http://localhost:8080/quizzes`
    // const url = `http://localhost:8080/quizzes?category=animais`
    this.axios.get(url).then((response) => {
        this.quizzes = response.data;
        console.log("Quiz list:\n", this.quizzes);
      });
  },
};
</script>
<style scoped>
.quiz-list {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  flex-wrap: wrap;
  margin-top: 20px;
}

.quiz-list-item {
  border: 1px solid;
  min-height: 50px;
  min-width: 200px;
  padding: 10px;
  margin: 2px 4px;
}

.quiz-list-item:hover {
  background-color: rgba(26, 71, 24, 0.5);
  color: white;
  cursor: pointer;
  font-weight: bold;
}

.quiz-list-item > img {
  width: 200px;
  
}
</style>