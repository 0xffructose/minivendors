<template>
    <!--<div class="container mx-auto px-4">
        <Navbar />
        <main class="flex">
            <div class="bg-amber-200 h-[75%] w-[50%]">
                
            </div>
            <div class="bg-red-200 h-[75%] w-[50%]"></div>
        </main>
    </div>-->
<div class="w-full min-h-screen bg-[#F9F9F9] text-black font-sans">
    
    <section class="h-screen w-full flex flex-col">
      
      <nav class="h-24 shrink-0 flex items-center justify-between px-8 border-b border-gray-300 bg-white">
        <div class="font-bold text-2xl tracking-tighter">MINIVENDORS</div>
        
        <div class="flex gap-8 text-sm font-semibold">
          <a href="#" class="border-b-2 border-black pb-1">SOLUTIONS</a>
          <a href="#" class="text-gray-500 hover:text-black">INFRASTRUCTURE</a>
          <a href="#" class="text-gray-500 hover:text-black">PRICING</a>
          <a href="#" class="text-gray-500 hover:text-black">DOCS</a>
        </div>

        <div class="flex gap-4">
          <button class="border border-black px-6 py-2 text-sm font-bold hover:bg-gray-100">LOG IN</button>
          <button class="bg-black text-white px-6 py-2 text-sm font-bold hover:bg-gray-800">GET STARTED</button>
        </div>
      </nav>

      <main class="flex-1 flex w-full">
        <div class="w-1/2 h-full p-16 flex flex-col justify-center border-r border-gray-300 bg-white">
          <span class="text-xs font-bold tracking-widest text-gray-500 border-b border-gray-300 w-max pb-2 mb-8">
            VERSION 2.0 RELEASED
          </span>
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
              READ DOCUMENTATION <span>→</span>
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
            EXPLORE API <span class="text-lg leading-none">↗</span>
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
            VIEW ARCHITECTURE <span class="text-lg leading-none">↗</span>
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
            SECURITY BRIEF <span class="text-lg leading-none">↗</span>
          </a>
        </div>

      </div>
    </section>

    <section class="w-full min-h-[75vh] flex flex-col md:flex-row">
      
      <div class="w-full md:w-1/2 p-16 lg:p-24 flex flex-col justify-center bg-black text-white">
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

      <div class="w-full md:w-1/2 p-16 lg:p-24 flex flex-col justify-center bg-[#0a0a0a]">
        <div class="w-full max-w-lg mx-auto">
          <div class="text-gray-600 font-mono text-xs mb-4 tracking-widest">
            // QUICK START
          </div>
          
          <div class="border border-gray-800 bg-black p-10 font-mono text-sm leading-loose shadow-2xl">
            <div class="flex gap-4">
              <span class="text-gray-600">$</span>
              <span class="text-gray-300">npm install @minivendors/sdk</span>
            </div>
            <div class="flex gap-4">
              <span class="text-gray-600">$</span>
              <span class="text-gray-300">mv init --template react</span>
            </div>
            <div class="flex gap-4">
              <span class="text-gray-600">$</span>
              <span class="text-gray-300">mv deploy --prod</span>
            </div>
            <div class="mt-8 text-gray-500">
              > Deployment successful. URL generated.
            </div>
          </div>
        </div>
      </div>

    </section>

    <footer class="w-full bg-white flex flex-col">
      
      <div class="grid grid-cols-2 lg:grid-cols-5 gap-12 p-16 w-full max-w-screen-2xl mx-auto">
        <div class="flex flex-col gap-4">
          <h4 class="text-xs font-bold tracking-widest mb-2">SOLUTIONS</h4>
          <a href="#" class="text-sm text-gray-600 hover:text-black">Marketplaces</a>
          <a href="#" class="text-sm text-gray-600 hover:text-black">Headless Commerce</a>
          <a href="#" class="text-sm text-gray-600 hover:text-black">Global Payments</a>
          <a href="#" class="text-sm text-gray-600 hover:text-black">Edge Delivery</a>
        </div>
        <div class="flex flex-col gap-4">
          <h4 class="text-xs font-bold tracking-widest mb-2">PLATFORM</h4>
          <a href="#" class="text-sm text-gray-600 hover:text-black">Infrastructure</a>
          <a href="#" class="text-sm text-gray-600 hover:text-black">API Reference</a>
          <a href="#" class="text-sm text-gray-600 hover:text-black">SDKs</a>
          <a href="#" class="text-sm text-gray-600 hover:text-black">Security</a>
        </div>
        <div class="flex flex-col gap-4">
          <h4 class="text-xs font-bold tracking-widest mb-2">RESOURCES</h4>
          <a href="#" class="text-sm text-gray-600 hover:text-black">Documentation</a>
          <a href="#" class="text-sm text-gray-600 hover:text-black">Changelog</a>
          <a href="#" class="text-sm text-gray-600 hover:text-black">Guides</a>
          <a href="#" class="text-sm text-gray-600 hover:text-black">Support</a>
        </div>
        <div class="flex flex-col gap-4">
          <h4 class="text-xs font-bold tracking-widest mb-2">COMPANY</h4>
          <a href="#" class="text-sm text-gray-600 hover:text-black">About Us</a>
          <a href="#" class="text-sm text-gray-600 hover:text-black">Careers</a>
          <a href="#" class="text-sm text-gray-600 hover:text-black">Contact</a>
          <a href="#" class="text-sm text-gray-600 hover:text-black">Partners</a>
        </div>
        <div class="flex flex-col lg:col-span-1 col-span-2 mt-4 lg:mt-0">
          <h4 class="text-xs font-bold tracking-widest mb-6">STAY UPDATED</h4>
          <input type="email" placeholder="EMAIL ADDRESS" class="border border-black px-4 py-3 w-full text-sm outline-none focus:ring-1 focus:ring-black placeholder-gray-400" />
          <button class="bg-black text-white px-4 py-3 w-full text-sm font-bold tracking-widest hover:bg-gray-800 mt-[-1px]">
            SUBSCRIBE
          </button>
        </div>
      </div>

      <div class="border-y border-gray-300 py-12 md:py-20 flex items-center justify-center overflow-hidden bg-white">
        <h2 class="text-[12vw] font-black leading-none tracking-tighter text-black select-none">
          MINIVENDORS
        </h2>
      </div>

      <div class="flex flex-col lg:flex-row w-full text-[10px] sm:text-xs font-bold tracking-widest text-gray-500">
        <div class="flex flex-wrap gap-4 sm:gap-8 p-6 border-b lg:border-b-0 lg:border-r border-gray-300 lg:w-1/3 justify-center lg:justify-start items-center">
          <a href="#" class="hover:text-black">TERMS OF SERVICE</a>
          <a href="#" class="hover:text-black">PRIVACY POLICY</a>
          <a href="#" class="hover:text-black">COOKIE SETTINGS</a>
        </div>
        
        <div class="flex gap-8 p-6 border-b lg:border-b-0 lg:border-r border-gray-300 lg:w-1/3 justify-center items-center">
          <a href="#" class="hover:text-black">GITHUB</a>
          <a href="#" class="hover:text-black">TWITTER</a>
          <a href="#" class="hover:text-black">DISCORD</a>
        </div>

        <div class="flex p-6 lg:w-1/3 justify-center lg:justify-end items-center text-gray-400">
          © 2026 MINIVENDORS. BUILT FOR SCALE.
        </div>
      </div>
    </footer>

  </div>
</template>

<script>
import Navbar from '../components/Navbar.vue';
export default {
    components: {
        Navbar
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