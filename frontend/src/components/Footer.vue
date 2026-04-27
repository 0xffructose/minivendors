<template>
    <footer class="w-full bg-white flex flex-col border-t border-gray-300">
      
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
</template>


<script>
export default {
    name: "Footer"
}
</script>