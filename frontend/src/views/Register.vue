<template>
  <div class="w-full min-h-screen bg-[#F9F9F9] text-black font-sans flex flex-col">

    <main class="flex-1 flex items-center justify-center p-6 sm:p-12">
      
      <div class="relative w-full max-w-[480px] bg-white border-1 border-black">

        <div class="flex flex-col relative z-10">
          
          <div class="p-10 sm:p-12 pb-8">
            <h1 class="text-3xl font-black tracking-tight mb-2 uppercase">REGISTER</h1>
            <p class="text-sm text-gray-600 mb-12">REGISTER TO MINIVENDORS.</p>

            <form @submit.prevent="handleRegister">
              <div class="flex flex-col gap-8 mb-8">
                
                <div>
                  <label for="name" class="block text-[10px] font-bold tracking-widest uppercase mb-4 text-black">
                    FULL NAME
                  </label>
                  <input 
                    id="name"
                    type="text" 
                    v-model="form.name"
                    required
                    placeholder="Enter full corporate or personal name" 
                    class="w-full border-b border-gray-400 pb-2 text-sm outline-none focus:border-black placeholder-gray-500 bg-transparent transition-colors" 
                  />
                </div>

                <div>
                  <label for="email" class="block text-[10px] font-bold tracking-widest uppercase mb-4 text-black">
                    E-MAIL
                  </label>
                  <input 
                    id="email"
                    type="email" 
                    v-model="form.email"
                    required
                    placeholder="Enter secure email address" 
                    class="w-full border-b border-gray-400 pb-2 text-sm outline-none focus:border-black placeholder-gray-500 bg-transparent transition-colors" 
                  />
                </div>

                <div>
                  <label for="password" class="block text-[10px] font-bold tracking-widest uppercase mb-4 text-black">
                    PASSWORD
                  </label>
                  <input 
                    id="password"
                    type="password" 
                    v-model="form.password"
                    required
                    placeholder="Generate secure passphrase" 
                    class="w-full border-b border-gray-400 pb-2 text-sm outline-none focus:border-black placeholder-gray-500 bg-transparent transition-colors" 
                  />
                </div>

                <div>
                  <label for="confirmPassword" class="block text-[10px] font-bold tracking-widest uppercase mb-4 text-black">
                    CONFIRM PASSWORD
                  </label>
                  <input 
                    id="confirmPassword"
                    type="password" 
                    v-model="confirmPassword"
                    required
                    placeholder="Repeat secure passphrase" 
                    class="w-full border-b border-gray-400 pb-2 text-sm outline-none focus:border-black placeholder-gray-500 bg-transparent transition-colors" 
                  />
                </div>

              </div>

              <div v-if="errorMessage" class="mb-4 text-[10px] font-bold tracking-widest text-red-500 uppercase">
                {{ errorMessage }}
              </div>

              <div v-if="successMessage" class="mb-4 text-[10px] font-bold tracking-widest text-green-600 uppercase">
                {{ successMessage }}
              </div>

              <p class="text-[11px] text-gray-600 leading-relaxed mb-8 pr-4">
                By proceeding, you agree to the MINIVENDORS Infrastructure Protocol and Network Directives.
              </p>

              <button 
                type="submit" 
                :disabled="isLoading"
                class="w-full bg-black text-white py-4 text-xs font-bold tracking-widest hover:bg-gray-800 transition-colors uppercase disabled:opacity-50 disabled:cursor-not-allowed">
                {{ isLoading ? 'CREATING ACCOUNT...' : 'CREATE ACCOUNT' }}
              </button>
            </form>
          </div>

          <router-link to="/login" class="border-t border-black p-6 flex justify-center items-center gap-3 text-[10px] sm:text-xs font-bold tracking-widest hover:bg-gray-50 transition-colors uppercase text-black bg-white">
            <svg class="w-4 h-4 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1"></path>
            </svg>
            AUTHENTICATE EXISTING PROFILE
          </router-link>

        </div>
      </div>

    </main>

  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';
import { useAuthStore } from '../stores/auth'; // Kendi dosya yoluna göre ayarla
import { useRouter } from 'vue-router';
import type { AxiosError } from 'axios';

const authStore = useAuthStore();
const router = useRouter();

const form = reactive({
  name: '',
  email: '',
  password: ''
});
const confirmPassword = ref('');

const errorMessage = ref('');
const successMessage = ref('');
const isLoading = ref(false);

const handleRegister = async () => {
  if (form.password !== confirmPassword.value) {
    errorMessage.value = 'Şifreler uyuşmuyor!';
    return;
  }

  isLoading.value = true;
  errorMessage.value = '';
  successMessage.value = '';
  
  try {
    await authStore.register(form);
    successMessage.value = 'Kayıt başarılı! Giriş sayfasına yönlendiriliyorsunuz...';
    
    setTimeout(() => {
      router.push('/login');
    }, 1500);

  } catch (error: unknown) {
    const err = error as AxiosError<{error: string}>;
    if (err.response && err.response.data && err.response.data.error) {
      errorMessage.value = err.response.data.error;
    } else {
      errorMessage.value = 'Kayıt sırasında beklenmeyen bir hata oluştu.';
    }
  } finally {
    isLoading.value = false;
  }
};
</script>