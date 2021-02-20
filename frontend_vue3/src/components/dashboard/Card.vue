<template>
  <div class="container" @click="onSelect">
    <h1 class="title">{{ title }}</h1>
    <div class="description">{{ description }}</div>
  </div>
</template>

<script>
export default {
  name: "dashboard-card",
  props: {
    title: String,
    description: String,
    onSelect: Function,
  },
  data() {
    return {};
  },
  methods: {},
};
</script>
<style scoped>
.container {
  display: flex;
  flex-direction: row;
  justify-content: center;

  width: 250px;
  height: 100px;
  padding: 10px;
  margin: 5px;
  border: 1px solid;
  border-radius: 8px;
  box-shadow: 4px 4px rgba(50, 50, 50, 0.5);
}

.container:hover {
    cursor: pointer;
}

.title {
  color: #35495e;
  font-weight: 700;
  font-size: 32px;
  margin-bottom: 4px;
  display: block;
  width: 100%;
}

.description {
  font-weight: 300;
  font-size: 20px;
  color: #8e99a2;
}
</style>
