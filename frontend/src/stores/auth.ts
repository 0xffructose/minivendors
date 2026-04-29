import { defineStore } from "pinia";
import apiClient from '../plugins/axios';

interface User {
    id: string;
    email: string;
    name: string;
}

interface AuthState {
    user: User | null;
    token: string | null;
}

export const useAuthStore = defineStore("auth", {
    state: (): AuthState => ({
        user: null,
        token: localStorage.getItem("token") || null,
    }),

    getters: {
        isAuthenticated: (state) => !!state.token,
    },

    actions: {
        async login(credentials: Record<string, string>) {
            try {   
                const response = await apiClient.post("/login", credentials);
                this.token = response.data.token;
                this.user = response.data.user;

                localStorage.setItem("token", this.token!);
            } catch(error) {
                throw error;
            }
        },



        async register(userData: Record<string, string>) {
            try {
                await apiClient.post('/register', userData);
            } catch (error) {
                throw error;
            }
        },

        logout() {
            this.token = null;
            this.user = null;

            localStorage.removeItem('token');
        }
    }
})