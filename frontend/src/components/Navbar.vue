<template>
    
    <nav class="h-24 shrink-0 flex items-center justify-between px-8 border-b border-gray-300 bg-white">
        <div class="font-bold text-2xl tracking-tighter">MINIVENDORS</div>
        
        <div class="flex gap-8 text-sm font-semibold">
            <a href="#" class="border-b-2 border-black pb-1"><router-link to="/features">FEATURES</router-link></a>
            <a href="#" class="text-gray-500 hover:text-black"><router-link to="/pricing">PRICING</router-link></a>
            <a href="#" class="text-gray-500 hover:text-black">DOCS</a>
        </div>

        <div class="flex gap-4">
            <button class="border border-black px-6 py-2 text-sm font-bold hover:bg-gray-100">LOG IN</button>
            <button class="bg-black text-white px-6 py-2 text-sm font-bold hover:bg-gray-800">GET STARTED</button>
        </div>
    </nav>

</template>

<script>
export default {
    name: "Navbar"
}
</script>