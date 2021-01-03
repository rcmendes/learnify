<template>
  <div>
    <h1>{{ msg }}</h1>
    <ul>
      <li v-for="(quiz, idx) in quizzes" :key="idx">
        <span>{{ quiz.tag }}</span>
        <div>{{ quiz.palavra }}</div>
        <div>{{ quiz.mot }}</div>
        <div>{{ quiz.audio }}</div>
        <img :src="quiz.url"/>
      </li>
    </ul>
    <button @click="loadQuizzes">Carregar</button>
  </div>
</template>

<script>
export default {
  name: "hello-world",
  props: {
    msg: String,
  },
  data() {
    return {
      quizzes: []
    };
  },
  methods: {
   loadQuizzes() {
     this.axios.get("http://localhost:8080/quizzes").then((response)=> {
       this.quizzes = response.data
     })
   }
  },
  mounted() {
    this.loadQuizzes()
  }
};
</script>
<style scoped>
* {
  box-sizing: border-box;
}

ul {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  flex-wrap: wrap;
}
li {
  display: block;
  
  width: 400px;
  padding: 10px;
  margin: 5px;
  border: 1px solid;
  border-radius: 8px;
  box-shadow: 4px 4px rgba(50, 50, 50, 0.5);
}

li > span {
  color: #35495e;
  font-weight: 700;
  font-size: 32px;
  margin-bottom: 4px;
  display: block;
}

li > div {
  font-weight: 300;
  font-size: 20px;
  color: #8e99a2;
}
img {
    width: 100px;
    border: 1px solid;

}
</style>
