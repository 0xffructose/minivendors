<template>
    <nav class="h-24 shrink-0 flex items-center justify-between px-8 border-b border-gray-300 bg-white">
        <div class="font-bold text-2xl tracking-tighter">
            <router-link to="/">MINIVENDORS</router-link>
        </div>
        
        <div class="flex gap-8 text-sm font-semibold">
            <router-link to="/features" class="text-gray-500 hover:text-black" exact-active-class="!text-black border-b-2 border-black">FEATURES</router-link>
            <router-link to="/pricing" class="text-gray-500 hover:text-black" exact-active-class="!text-black border-b-2 border-black">PRICING</router-link>
            <a href="#" class="text-gray-500 hover:text-black">DOCS</a>
        </div>

        <div v-if="!authStore.isAuthenticated" class="flex gap-4">
            <router-link to="/login" class="border border-black px-6 py-2 text-sm font-bold flex items-center justify-center hover:bg-gray-100">
                LOG IN
            </router-link>
            <router-link to="/register" class="bg-black text-white px-6 py-2 text-sm font-bold flex items-center justify-center hover:bg-gray-800">
                GET STARTED
            </router-link>
        </div>

        <div v-else class="flex items-center gap-4">
            <span class="text-sm font-semibold text-gray-600 mr-2">
                Merhaba, {{ authStore.user?.name || 'Kullanıcı' }}
            </span>
            <router-link to="/dashboard" class="text-sm font-bold text-gray-500 hover:text-black">
                DASHBOARD
            </router-link>
            <button @click="handleLogout" class="border border-red-500 text-red-500 px-6 py-2 text-sm font-bold hover:bg-red-50 transition-colors">
                LOG OUT
            </button>
        </div>
    </nav>
</template>

<script setup lang="ts">
import { useAuthStore } from '../stores/auth'; // Store yolunun projenle eşleştiğinden emin ol
import { useRouter } from 'vue-router';

const authStore = useAuthStore();
const router = useRouter();

const handleLogout = () => {
    authStore.logout();
    router.push('/login');
};
</script>