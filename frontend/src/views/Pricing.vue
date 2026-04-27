<template>
  <div class="w-full min-h-screen bg-[#F9F9F9] text-black font-sans">
    
    <section class="h-screen w-full flex flex-col bg-white">
      
      <Navbar />

      <main class="flex-1 flex flex-col md:flex-row w-full items-stretch">
        <div class="w-full md:w-1/2 p-8 md:p-16 lg:p-24 flex flex-col justify-center border-b md:border-b-0 md:border-r border-gray-300 z-10 bg-white">
          <h1 class="text-5xl lg:text-7xl font-black leading-[1.1] tracking-tight mb-8">
            PRICING INFRASTRUCTURE.<br />
            SCALED FOR COMMERCE.
          </h1>
          <p class="text-lg lg:text-xl text-gray-600 max-w-xl leading-relaxed">
            Transparent, usage-based pricing designed to support your growth from zero 
            to millions of transactions. No hidden fees. No decorative constraints.
          </p>
        </div>

        <div class="w-full md:w-1/2 flex items-center justify-center architectural-grid relative overflow-hidden bg-[#f4f4f4]">
          <div class="absolute inset-0 bg-gradient-to-r from-white/40 via-transparent to-black/5"></div>
          <div class="absolute inset-0 bg-gradient-to-b from-white/20 via-transparent to-transparent"></div>
        </div>
      </main>

    </section>

    <section class="w-full min-h-[80vh] grid grid-cols-1 md:grid-cols-3 bg-white border-t border-gray-300">
      
      <div class="flex flex-col p-12 lg:p-24 border-b md:border-b-0 md:border-r border-gray-300 bg-white">
        <h2 class="text-3xl font-bold tracking-tight mb-4">STARTER</h2>
        <p class="text-sm text-gray-600 mb-12 h-10">For independent developers and small projects.</p>
        
        <div class="flex items-baseline gap-2 mb-10">
          <span class="text-6xl lg:text-7xl font-black tracking-tighter">$0</span>
          <span class="text-xs font-bold text-gray-500 tracking-widest">/ MONTH</span>
        </div>
        
        <hr class="border-black mb-10" />
        
        <ul class="flex flex-col gap-6 text-sm font-medium text-gray-800 mb-16 flex-1">
          <li class="flex items-center gap-4">
            <svg class="w-6 h-6 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
            Up to 1,000 API calls/mo
          </li>
          <li class="flex items-center gap-4">
            <svg class="w-6 h-6 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
            Standard Support
          </li>
          <li class="flex items-center gap-4">
            <svg class="w-6 h-6 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
            Basic Analytics
          </li>
        </ul>

        <button class="w-full border border-black py-6 text-xs font-bold tracking-widest hover:bg-gray-50 mt-auto">
          START FOR FREE
        </button>
      </div>

      <div class="flex flex-col p-12 lg:p-24 bg-black text-white relative">
        <div class="absolute top-0 right-0 bg-white text-black px-6 py-3 text-[10px] font-bold tracking-widest">
          RECOMMENDED
        </div>

        <h2 class="text-3xl font-bold tracking-tight mb-4 mt-6 lg:mt-0">PROFESSIONAL</h2>
        <p class="text-sm text-gray-400 mb-12 h-10">For scaling businesses and growing teams.</p>
        
        <div class="flex items-baseline gap-2 mb-10">
          <span class="text-6xl lg:text-7xl font-black tracking-tighter">$49</span>
          <span class="text-xs font-bold text-gray-500 tracking-widest">/ MONTH</span>
        </div>
        
        <hr class="border-gray-700 mb-10" />
        
        <ul class="flex flex-col gap-6 text-sm font-medium text-gray-300 mb-16 flex-1">
          <li class="flex items-center gap-4">
            <svg class="w-6 h-6 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
            Up to 100,000 API calls/mo
          </li>
          <li class="flex items-center gap-4">
            <svg class="w-6 h-6 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
            Priority Support (24h SLA)
          </li>
          <li class="flex items-center gap-4">
            <svg class="w-6 h-6 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
            Advanced Analytics & Webhooks
          </li>
          <li class="flex items-center gap-4">
            <svg class="w-6 h-6 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
            Custom Domains
          </li>
        </ul>

        <button class="w-full bg-white text-black py-6 text-xs font-bold tracking-widest hover:bg-gray-200 mt-auto">
          UPGRADE TO PRO
        </button>
      </div>

      <div class="flex flex-col p-12 lg:p-24 border-t md:border-t-0 md:border-l border-gray-300 bg-white">
        <h2 class="text-3xl font-bold tracking-tight mb-4">ENTERPRISE</h2>
        <p class="text-sm text-gray-600 mb-12 h-10">Custom infrastructure for high-volume platforms.</p>
        
        <div class="flex items-baseline gap-2 mb-10">
          <span class="text-6xl lg:text-7xl font-black tracking-tighter">CUSTOM</span>
        </div>
        
        <hr class="border-black mb-10" />
        
        <ul class="flex flex-col gap-6 text-sm font-medium text-gray-800 mb-16 flex-1">
          <li class="flex items-center gap-4">
            <svg class="w-6 h-6 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
            Unlimited API calls
          </li>
          <li class="flex items-center gap-4">
            <svg class="w-6 h-6 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
            Dedicated Account Manager
          </li>
          <li class="flex items-center gap-4">
            <svg class="w-6 h-6 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
            Custom SLA (99.99% Uptime)
          </li>
          <li class="flex items-center gap-4">
            <svg class="w-6 h-6 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"></path></svg>
            On-Premise Deployment Options
          </li>
        </ul>

        <button class="w-full border border-black py-6 text-xs font-bold tracking-widest hover:bg-gray-50 mt-auto">
          CONTACT SALES
        </button>
      </div>

    </section>

    <Footer />

  </div>
</template>

<script>
import Navbar from '../components/Navbar.vue';
import Footer from '../components/Footer.vue';

export default {
    components: {
      Navbar,
      Footer
    }
}
</script>

<style scoped>
/* Mimari / Blueprint tarzı arka plan ızgarası.
  Saf CSS ile oluşturulmuştur, sayfa hızını etkilemez.
*/
.architectural-grid {
  background-image: 
    /* İnce detay çizgileri */
    linear-gradient(rgba(0, 0, 0, 0.03) 1px, transparent 1px),
    linear-gradient(90deg, rgba(0, 0, 0, 0.03) 1px, transparent 1px),
    /* Daha belirgin taşıyıcı çizgiler */
    linear-gradient(rgba(0, 0, 0, 0.05) 2px, transparent 2px),
    linear-gradient(90deg, rgba(0, 0, 0, 0.05) 2px, transparent 2px);
  
  /* Çizgilerin aralık boyutları (Önce ince, sonra kalın çizgiler için) */
  background-size: 20px 20px, 20px 20px, 100px 100px, 100px 100px;
  background-position: center center;
}
</style>