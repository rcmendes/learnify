<template>
  <div>
    <h1>{{ msg }}</h1>
    <div class="container">
      <Card
        v-for="(quiz, idx) in quizzes"
        :key="idx"
        :title="quiz.category"
        :description="''"
        :on-select="()=> onSelectCategory(quiz)"
      />
    </div>
  </div>
</template>

<script>
import Card from "../components/dashboard/Card.vue";
export default {
  name: "dashboard",
  components: {
    Card,
  },
  props: {
    msg: String,
  },
  data() {
    return {
      quizzes: [],
    };
  },
  methods: {
    onSelectCategory(quiz) {
        console.log('Selected category:', quiz.category)
        this.$router.push('/game')
    },
  },
  mounted() {
    this.axios.get("http://localhost:8080/quizzes").then((response) => {
        let data = response.data;
        const filteredList = []
        data.forEach(q => {
            const exists = filteredList.filter(element => q.category == element.category).length != 0
            if (!exists) {
                filteredList.push(q)
            }
        });

        this.quizzes = filteredList;
      });
  },
};
</script>
<style scoped>
.container {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  flex-wrap: wrap;
}
</style>
