<template>

<div class="w-full min-h-screen bg-[#F9F9F9] text-black font-sans">
    
    <section class="h-screen w-full flex flex-col">
      
      <Navbar />

      <main class="flex-1 flex w-full">
        <div class="w-1/2 h-full p-16 flex flex-col justify-center border-r border-gray-300 bg-white">
          <h1 class="text-6xl font-black leading-tight tracking-tight mb-6">
            THE ARCHITECTURE OF<br />MODERN COMMERCE.
          </h1>
          <p class="text-lg text-gray-600 max-w-md mb-12">
            Minimalist, high-performance infrastructure for developers building the next generation of global marketplaces and independent storefronts.
          </p>
          <div class="flex gap-4 mt-auto lg:mt-0">
            <button class="bg-black text-white px-8 py-4 text-xs font-bold tracking-widest hover:bg-gray-800">
              INITIALIZE DEPLOYMENT
            </button>
            <button class="border border-gray-300 px-8 py-4 text-xs font-bold tracking-widest hover:bg-gray-50 flex items-center gap-2">
              READ DOCUMENTATION <span><i class="pi pi-arrow-right"></i></span>
            </button>
          </div>
        </div>

        <div class="w-1/2 h-full p-16 flex items-center justify-center bg-[#F3F3F3]">
          <div class="w-full max-w-lg h-80 bg-white border border-gray-300 shadow-[8px_8px_0px_0px_rgba(0,0,0,1)] p-8">
              </div>
        </div>
      </main>
    </section>

    <section class="w-full bg-white flex flex-col">
      
      <div class="border-b border-t border-gray-300 py-8 overflow-hidden relative w-full flex bg-white">
        <div class="flex w-[200%] animate-marquee">
          <div class="flex w-1/2 justify-around items-center shrink-0">
            <span class="text-2xl font-black text-gray-400 tracking-tighter">STRIPE</span>
            <span class="text-2xl font-black text-gray-400 tracking-tighter">AWS</span>
            <span class="text-2xl font-black text-gray-400 tracking-tighter">VERCEL</span>
            <span class="text-2xl font-black text-gray-400 tracking-tighter">ALGOLIA</span>
            <span class="text-2xl font-black text-gray-400 tracking-tighter">POSTGRES</span>
          </div>
          <div class="flex w-1/2 justify-around items-center shrink-0">
            <span class="text-2xl font-black text-gray-400 tracking-tighter">STRIPE</span>
            <span class="text-2xl font-black text-gray-400 tracking-tighter">AWS</span>
            <span class="text-2xl font-black text-gray-400 tracking-tighter">VERCEL</span>
            <span class="text-2xl font-black text-gray-400 tracking-tighter">ALGOLIA</span>
            <span class="text-2xl font-black text-gray-400 tracking-tighter">POSTGRES</span>
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-16 p-16 w-full max-w-screen-2xl mx-auto">
        <div class="flex flex-col">
          <div class="w-16 h-16 border border-black flex items-center justify-center mb-8 shrink-0">
            <svg class="w-6 h-6" fill="none" stroke="black" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"></path>
            </svg>
          </div>
          <h3 class="text-xl font-bold mb-4 tracking-tight">API-FIRST DESIGN</h3>
          <p class="text-sm text-gray-600 leading-relaxed mb-8 flex-1">
            Headless by default. Build bespoke storefronts with any framework while we handle the complex backend logic, cart state, and checkout flows.
          </p>
          <a href="#" class="text-xs font-bold tracking-widest flex items-center gap-2 hover:text-gray-500 w-max mt-auto uppercase">
            EXPLORE API <span class="text-md leading-none"><i class="pi pi-arrow-up-right"></i></span>
          </a>
        </div>

        <div class="flex flex-col">
          <div class="w-16 h-16 border border-black flex items-center justify-center mb-8 shrink-0">
            <svg class="w-6 h-6" fill="none" stroke="black" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 002-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"></path>
            </svg>
          </div>
          <h3 class="text-xl font-bold mb-4 tracking-tight">INFINITE SCALABILITY</h3>
          <p class="text-sm text-gray-600 leading-relaxed mb-8 flex-1">
            Deployed across a global edge network. Our infrastructure automatically provisions resources to handle traffic spikes without manual intervention.
          </p>
          <a href="#" class="text-xs font-bold tracking-widest flex items-center gap-2 hover:text-gray-500 w-max mt-auto uppercase">
            VIEW ARCHITECTURE <span class="text-md leading-none"><i class="pi pi-arrow-up-right"></i></span>
          </a>
        </div>

        <div class="flex flex-col">
          <div class="w-16 h-16 border border-black flex items-center justify-center mb-8 shrink-0">
            <svg class="w-6 h-6" fill="none" stroke="black" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path>
            </svg>
          </div>
          <h3 class="text-xl font-bold mb-4 tracking-tight">ZERO-TRUST SECURITY</h3>
          <p class="text-sm text-gray-600 leading-relaxed mb-8 flex-1">
            Enterprise-grade compliance out of the box. SOC2 certified, PCI DSS Level 1 compliant data vaults for sensitive customer information.
          </p>
          <a href="#" class="text-xs font-bold tracking-widest flex items-center gap-2 hover:text-gray-500 w-max mt-auto uppercase">
            SECURITY BRIEF <span class="text-md leading-none"><i class="pi pi-arrow-up-right"></i></span>
          </a>
        </div>

      </div>
    </section>

    <section class="w-full min-h-[75vh] flex flex-col md:flex-row">
      
      <div class="w-full md:w p-16 lg:p-24 flex flex-col justify-center bg-black text-white">
        <h2 class="text-5xl lg:text-7xl font-black leading-tight tracking-tight mb-6">
          READY TO DEPLOY?
        </h2>
        <p class="text-lg text-gray-400 max-w-md mb-12">
          Join thousands of developers building the future of commerce on MiniVendors. Start building locally in seconds, push to production in minutes.
        </p>
        <div class="flex flex-col sm:flex-row gap-4">
          <button class="bg-white text-black px-8 py-4 text-xs font-bold tracking-widest hover:bg-gray-200 transition-colors">
            CREATE FREE ACCOUNT
          </button>
          <button class="border border-gray-600 text-white px-8 py-4 text-xs font-bold tracking-widest hover:border-white transition-colors">
            CONTACT SALES
          </button>
        </div>
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
/* Soldan sağa kayan metin animasyonu */
.animate-marquee {
  animation: marquee-ltr 25s linear infinite;
}

@keyframes marquee-ltr {
  0% { transform: translateX(-50%); }
  100% { transform: translateX(0%); }
}
</style>
