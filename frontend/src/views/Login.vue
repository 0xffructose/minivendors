<template>
  <div class="w-full min-h-screen bg-[#F9F9F9] text-black font-sans flex flex-col">

    <main class="flex-1 flex items-center justify-center p-6 sm:p-12">
      
      <div class="relative w-full max-w-[420px] bg-white border-1 border-black">

        <div class="flex flex-col relative z-10">
          
          <div class="p-10 sm:p-12 pb-8">
            <h1 class="text-3xl font-black tracking-tight mb-2 uppercase">LOGIN</h1>
            <p class="text-sm text-gray-600 mb-12">Login to your profile.</p>

            <form @submit.prevent="handleLogin">
              <div class="flex flex-col gap-8 mb-6">
                
                <div>
                  <label for="email" class="block text-[10px] font-bold tracking-widest uppercase mb-4 text-black">
                    E-MAIL
                  </label>
                  <input 
                    id="email"
                    type="email" 
                    v-model="email"
                    required
                    placeholder="Enter secure email address" 
                    class="w-full border-b border-gray-400 pb-2 text-sm outline-none focus:border-black placeholder-gray-500 bg-transparent transition-colors" 
                  />
                </div>

                <div>
                  <div class="flex justify-between items-end mb-4">
                    <label for="password" class="block text-[10px] font-bold tracking-widest uppercase text-black">
                      PASSWORD
                    </label>
                    <a href="#" class="text-[9px] font-bold tracking-widest text-gray-500 hover:text-black uppercase transition-colors">
                      FORGOT PASSWORD?
                    </a>
                  </div>
                  <input 
                    id="password"
                    type="password" 
                    v-model="password"
                    required
                    placeholder="Enter secure passphrase" 
                    class="w-full border-b border-gray-400 pb-2 text-sm outline-none focus:border-black placeholder-gray-500 bg-transparent transition-colors" 
                  />
                </div>

              </div>

              <div v-if="errorMessage" class="mb-4 text-[10px] font-bold tracking-widest text-red-500 uppercase">
                {{ errorMessage }}
              </div>

              <button 
                type="submit" 
                :disabled="isLoading"
                class="w-full bg-black text-white py-4 text-xs font-bold tracking-widest hover:bg-gray-800 transition-colors uppercase mt-4 disabled:opacity-50 disabled:cursor-not-allowed">
                {{ isLoading ? 'LOGGING IN...' : 'LOGIN' }}
              </button>
            </form>

          </div>

          <router-link to="/register" class="border-t border-black p-6 flex justify-center items-center gap-3 text-[10px] sm:text-xs font-bold tracking-widest hover:bg-gray-50 transition-colors uppercase text-black bg-white">
            INITIALIZE NEW PROFILE
            <svg class="w-4 h-4 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7l5 5m0 0l-5 5m5-5H6"></path>
            </svg>
          </router-link>

        </div>
      </div>

    </main>

  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useAuthStore } from '../stores/auth';
import { useRouter } from 'vue-router';

const email = ref('');
const password = ref('');
const errorMessage = ref('');
const isLoading = ref(false);

const authStore = useAuthStore();
const router = useRouter();

const handleLogin = async () => {
  isLoading.value = true;
  errorMessage.value = '';
  
  try {
    await authStore.login({ email: email.value, password: password.value });
    router.push('/dashboard');
  } catch (error) {
    errorMessage.value = 'Giriş başarısız. Lütfen bilgilerinizi kontrol edin.';
  } finally {
    isLoading.value = false;
  }
};
</script>