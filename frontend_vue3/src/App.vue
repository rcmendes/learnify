<template>
  <div>
    <menu-bar />
    <router-view></router-view>
  </div>
</template>

<script>
import MenuBar from "./components/MenuBar.vue";
export default {
  name: "App",
  components: {
    MenuBar,
  },
};
</script>