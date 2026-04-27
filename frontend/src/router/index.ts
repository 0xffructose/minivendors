import { createRouter, createWebHistory, type RouteRecordRaw } from "vue-router";
import Home from "../views/Home.vue";

const routes: Array<RouteRecordRaw> = [
    {
        path: "/",
        name: "home",
        component: Home
    },
    {
        path: "/pricing",
        name: "pricing",
        component: () => import("../views/Pricing.vue")
    },
    {
        path: "/dashboard",
        name: "dashboard",
        component: () => import("../views/Dashboard.vue")
    },
    {
        path: "/about",
        name: "about",
        component: () => import("../views/About.vue")
    }
];

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes
});

export default router;