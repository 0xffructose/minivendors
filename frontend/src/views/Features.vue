<template>
  <div class="w-full min-h-screen bg-[#F9F9F9] text-black font-sans">
    
    <section class="h-screen w-full flex flex-col bg-white">
      
        <Navbar />

      <main class="flex-1 flex flex-col md:flex-row w-full items-stretch">
        
        <div class="w-full md:w-1/2 p-8 md:p-16 lg:p-24 flex flex-col justify-center border-b md:border-b-0 md:border-r border-gray-300 bg-white z-10">
          <h1 class="text-5xl md:text-6xl lg:text-[5rem] font-black leading-[1.05] tracking-tight mb-8">
            UNCOMPROMISING<br/>CAPABILITIES.
          </h1>
          <p class="text-sm lg:text-base text-gray-600 max-w-xl leading-relaxed font-medium uppercase">
            Engineered for precision and scale. A developer-first e-commerce primitive 
            providing raw structural power without the bloat of traditional monoliths.
          </p>
        </div>

        <div class="w-full md:w-1/2 flex items-center justify-center architectural-grid relative overflow-hidden bg-[#f4f4f4]">
          <div class="absolute inset-0 bg-gradient-to-r from-white/40 via-transparent to-black/5"></div>
          <div class="absolute inset-0 bg-gradient-to-b from-white/20 via-transparent to-transparent"></div>
        </div>

      </main>
    </section>

    <section class="w-full min-h-[60vh] grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 bg-white border-t border-gray-300">
      
      <div class="p-12 lg:p-16 border-b lg:border-b-0 md:border-r border-gray-300 flex flex-col hover:bg-gray-50 transition-colors">
        <div class="mb-12">
          <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zm10 0a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zm10 0a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z"></path>
          </svg>
        </div>
        <h3 class="text-2xl lg:text-3xl font-black tracking-tight mb-6 uppercase leading-tight">
          HIGH-<br/>PERFORMANCE<br/>API
        </h3>
        <p class="text-sm text-gray-600 leading-relaxed">
          Native support for REST and GraphQL. Built for raw throughput with minimal overhead.
        </p>
      </div>

      <div class="p-12 lg:p-16 border-b lg:border-b-0 lg:border-r border-gray-300 flex flex-col hover:bg-gray-50 transition-colors">
        <div class="mb-12">
          <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.055 11H5a2 2 0 012 2v1a2 2 0 002 2 2 2 0 012 2v2.945M8 3.935V5.5A2.5 2.5 0 0010.5 8h.5a2 2 0 012 2 2 2 0 104 0 2 2 0 012-2h1.064M15 20.488V18a2 2 0 012-2h3.064M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path>
          </svg>
        </div>
        <h3 class="text-2xl lg:text-3xl font-black tracking-tight mb-6 uppercase leading-tight">
          GLOBAL EDGE<br/>NETWORK
        </h3>
        <p class="text-sm text-gray-600 leading-relaxed">
          Sub-50ms latency globally. Data replicated across tier-1 data centers for zero-downtime reliability.
        </p>
      </div>

      <div class="p-12 lg:p-16 border-b md:border-b-0 md:border-r border-gray-300 flex flex-col hover:bg-gray-50 transition-colors">
        <div class="mb-12">
          <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"></path>
          </svg>
        </div>
        <h3 class="text-2xl lg:text-3xl font-black tracking-tight mb-6 uppercase leading-tight">
          SECURE-BY-<br/>DESIGN
        </h3>
        <p class="text-sm text-gray-600 leading-relaxed">
          SOC2 Type II and PCI DSS Level 1 compliant out of the box. End-to-end encryption at rest and in transit.
        </p>
      </div>

      <div class="p-12 lg:p-16 flex flex-col hover:bg-gray-50 transition-colors">
        <div class="mb-12">
          <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 3v2m6-2v2M9 19v2m6-2v2M5 9H3m2 6H3m18-6h-2m2 6h-2M7 19h10a2 2 0 002-2V7a2 2 0 00-2-2H7a2 2 0 00-2 2v10a2 2 0 002 2zM9 9h6v6H9V9z"></path>
          </svg>
        </div>
        <h3 class="text-2xl lg:text-3xl font-black tracking-tight mb-6 uppercase leading-tight">
          AI-POWERED<br/>OPTIMIZATION
        </h3>
        <p class="text-sm text-gray-600 leading-relaxed">
          Advanced machine learning models that optimize your e-commerce workflows, from predictive inventory to automated customer insights.
        </p>
      </div>

    </section>

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