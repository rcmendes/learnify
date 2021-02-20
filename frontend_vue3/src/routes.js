import { createWebHistory, createRouter } from "vue-router";
// import Home from "@/views/Home.vue";
import Dashboard from "./views/Dashboard.vue";
import About from "./components/About.vue";
import ScoreBoard from "./components/ScoreBoard.vue";
import Game from "./views/game/Game.vue";
import Quiz from "./views/game/Quiz.vue";

const routes = [
    {
        path: "/",
        name: "Home",
        component: Dashboard,
    },
    {
        path: "/game",
        name: "Game",
        component: Game,
    },
    {
        path: "/score-board",
        name: "ScoreBoard",
        component: ScoreBoard,
    },
    {
        path: "/about",
        name: "About",
        component: About,
    },
    {
        path: "/quiz",
        name: "Quiz",
        component: Quiz,
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

export default router;