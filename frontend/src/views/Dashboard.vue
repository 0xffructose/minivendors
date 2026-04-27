<template>
  <div class="w-full h-screen bg-[#F9F9F9] text-black font-sans flex overflow-hidden">

    <aside class="w-64 h-full bg-white border-r border-gray-300 flex flex-col shrink-0 hidden md:flex">
      
      <div class="h-16 shrink-0 flex items-center px-6 border-b border-gray-300 gap-3">
        <div class="w-8 h-8 bg-black flex items-center justify-center text-white text-[10px] font-bold">IC</div>
        <div>
          <div class="font-bold text-sm tracking-widest leading-none mb-1">INFRA_CORE</div>
          <div class="text-[10px] text-gray-500 font-mono leading-none">v2.4.0-STABLE</div>
        </div>
      </div>

      <nav class="flex flex-col py-6 flex-1 gap-1">
        <a href="#" class="flex items-center gap-4 px-6 py-4 bg-black text-white font-bold text-xs tracking-widest">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z"></path></svg>
          DASHBOARD
        </a>
        <a href="#" class="flex items-center gap-4 px-6 py-4 text-gray-600 hover:text-black hover:bg-gray-50 font-bold text-xs tracking-widest transition-colors">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 002-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"></path></svg>
          INFRASTRUCTURE
        </a>
        <a href="#" class="flex items-center gap-4 px-6 py-4 text-gray-600 hover:text-black hover:bg-gray-50 font-bold text-xs tracking-widest transition-colors">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4"></path></svg>
          INVENTORY
        </a>
        <a href="#" class="flex items-center gap-4 px-6 py-4 text-gray-600 hover:text-black hover:bg-gray-50 font-bold text-xs tracking-widest transition-colors">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"></path></svg>
          ORDERS
        </a>
        <a href="#" class="flex items-center gap-4 px-6 py-4 text-gray-600 hover:text-black hover:bg-gray-50 font-bold text-xs tracking-widest transition-colors">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path></svg>
          ANALYTICS
        </a>
        <a href="#" class="flex items-center gap-4 px-6 py-4 text-gray-600 hover:text-black hover:bg-gray-50 font-bold text-xs tracking-widest transition-colors">
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"></path><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"></path></svg>
          SETTINGS
        </a>
      </nav>

      <div class="p-6 border-t border-gray-300 flex flex-col gap-8">
        <button class="bg-black text-white w-full py-4 text-xs font-bold tracking-widest hover:bg-gray-800 transition-colors">
          DEPLOY NEW NODE
        </button>
        <div class="flex flex-col gap-4">
          <a href="#" class="flex items-center gap-3 text-xs font-bold text-gray-500 tracking-widest hover:text-black">
             <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path></svg>
             DOCUMENTATION
          </a>
          <a href="#" class="flex items-center gap-3 text-xs font-bold text-gray-500 tracking-widest hover:text-black">
             <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
             SUPPORT
          </a>
        </div>
      </div>
    </aside>

    <div class="flex-1 flex flex-col h-full overflow-hidden bg-[#F9F9F9]">
      
      <header class="h-16 shrink-0 bg-white border-b border-gray-300 flex items-center justify-between px-6 lg:px-8">
        
        <div class="flex items-center border border-gray-300 px-3 py-2 w-full max-w-xs bg-white">
          <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
          <input type="text" placeholder="QUERY_DATA..." class="bg-transparent text-xs outline-none w-full ml-2 font-mono placeholder-gray-400" />
        </div>

        <div class="hidden lg:flex gap-8 text-[10px] font-bold text-gray-500 tracking-widest uppercase">
          <a href="#" class="hover:text-black">NETWORK STATUS</a>
          <a href="#" class="hover:text-black">LOGS</a>
          <a href="#" class="hover:text-black">SYSTEM HEALTH</a>
        </div>

        <div class="flex items-center gap-4 lg:border-l border-gray-300 lg:pl-6 h-full">
           <button class="text-black hover:text-gray-500"><svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"></path></svg></button>
           <button class="text-black hover:text-gray-500"><svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.121 17.804A13.937 13.937 0 0112 16c2.5 0 4.847.655 6.879 1.804M15 10a3 3 0 11-6 0 3 3 0 016 0zm6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg></button>
           <button class="bg-black text-white px-6 py-2.5 text-[10px] font-bold tracking-widest hover:bg-gray-800 ml-2">UPGRADE</button>
        </div>
      </header>

      <main class="flex-1 overflow-y-auto p-6 lg:p-12">
        
        <div class="flex flex-col sm:flex-row justify-between items-start sm:items-end mb-8 gap-4">
          <div>
            <h1 class="text-3xl lg:text-4xl font-black tracking-tight mb-2 uppercase">SYSTEM_OVERVIEW</h1>
            <p class="text-xs text-gray-500 font-mono">Metrics generated at 14:32:00 UTC</p>
          </div>
          <button class="border border-black px-4 py-2 text-xs font-bold tracking-widest hover:bg-gray-100 flex items-center gap-2 bg-white">
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"></path></svg> 
            EXPORT REPORT
          </button>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
          <div class="bg-white border border-gray-300 p-6 flex flex-col justify-between h-36">
             <div class="flex justify-between items-center text-[10px] font-bold tracking-widest text-gray-500 uppercase">
                TOTAL REVENUE <svg class="w-4 h-4 text-black" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
             </div>
             <div>
                <div class="text-3xl font-black tracking-tighter mb-1">$124,500.00</div>
                <div class="text-[10px] text-black font-mono">↗ +12.5% vs LAST_PERIOD</div>
             </div>
          </div>
          <div class="bg-white border border-gray-300 p-6 flex flex-col justify-between h-36">
             <div class="flex justify-between items-center text-[10px] font-bold tracking-widest text-gray-500 uppercase">
                TOTAL ORDERS <svg class="w-4 h-4 text-black" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"></path></svg>
             </div>
             <div>
                <div class="text-3xl font-black tracking-tighter mb-1">1,432</div>
                <div class="text-[10px] text-gray-500 font-mono">↗ +5.2% vs LAST_PERIOD</div>
             </div>
          </div>
          <div class="bg-white border border-gray-300 p-6 flex flex-col justify-between h-36">
             <div class="flex justify-between items-center text-[10px] font-bold tracking-widest text-gray-500 uppercase">
                AVG ORDER VALUE <svg class="w-4 h-4 text-black" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"></path></svg>
             </div>
             <div>
                <div class="text-3xl font-black tracking-tighter mb-1">$86.94</div>
                <div class="text-[10px] text-gray-500 font-mono">→ -0.1% vs LAST_PERIOD</div>
             </div>
          </div>
          <div class="bg-white border border-gray-300 p-6 flex flex-col justify-between h-36">
             <div class="flex justify-between items-center text-[10px] font-bold tracking-widest text-gray-500 uppercase">
                CONVERSION RATE <svg class="w-4 h-4 text-black" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"></path></svg>
             </div>
             <div>
                <div class="text-3xl font-black tracking-tighter mb-1">3.24%</div>
                <div class="text-[10px] text-black font-mono">↗ +1.1% vs LAST_PERIOD</div>
             </div>
          </div>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6 mb-8">
           
           <div class="lg:col-span-2 bg-white border border-gray-300 p-6 flex flex-col h-[400px]">
              <div class="flex justify-between items-center border-b border-gray-300 pb-4 mb-6">
                 <h2 class="text-lg font-black tracking-tight uppercase">REVENUE_TIMELINE</h2>
                 <div class="flex text-[10px] font-bold border border-black">
                   <button class="bg-black text-white px-3 py-1">7D</button>
                   <button class="px-3 py-1 hover:bg-gray-100 border-x border-black">30D</button>
                   <button class="px-3 py-1 hover:bg-gray-100">90D</button>
                 </div>
              </div>
              
              <div class="flex-1 flex relative">
                  <div class="w-8 flex flex-col justify-between text-[10px] font-mono text-gray-400 pb-6">
                      <span>15k</span>
                      <span>10k</span>
                      <span>5k</span>
                      <span>0</span>
                  </div>
                  <div class="flex-1 border-b border-l border-gray-300 flex items-end justify-around px-2 pb-0 relative">
                      <div class="absolute w-full h-px bg-gray-100 bottom-[25%] left-0"></div>
                      <div class="absolute w-full h-px bg-gray-100 bottom-[50%] left-0"></div>
                      <div class="absolute w-full h-px bg-gray-100 bottom-[75%] left-0"></div>
                      
                      <div class="w-[8%] bg-black z-10" style="height: 45%"></div>
                      <div class="w-[8%] bg-black z-10" style="height: 65%"></div>
                      <div class="w-[8%] bg-black z-10" style="height: 35%"></div>
                      <div class="w-[8%] bg-black z-10" style="height: 85%"></div>
                      <div class="w-[8%] bg-black z-10" style="height: 55%"></div>
                      <div class="w-[8%] bg-black z-10" style="height: 90%"></div>
                      <div class="w-[8%] bg-black z-10" style="height: 100%"></div>
                  </div>
              </div>
              <div class="flex justify-around text-[10px] font-mono text-gray-400 mt-2 ml-8">
                  <span>MON</span><span>TUE</span><span>WED</span><span>THU</span><span>FRI</span><span>SAT</span><span>SUN</span>
              </div>
           </div>

           <div class="bg-white border border-gray-300 p-6 h-[400px] flex flex-col">
              <h2 class="text-lg font-black tracking-tight border-b border-gray-300 pb-4 mb-6 uppercase">TOP_SKUS</h2>
              <div class="flex flex-col gap-6 overflow-y-auto pr-2 flex-1">
                 
                 <div class="flex items-center gap-4">
                    <div class="w-14 h-14 bg-[#F9F9F9] border border-gray-300 flex items-center justify-center text-[8px] font-mono text-gray-400 shrink-0">IMG_01</div>
                    <div class="flex-1 min-w-0">
                       <div class="text-sm font-bold truncate">Mechanical Keyboard v2</div>
                       <div class="text-[10px] text-gray-500 font-mono truncate mt-0.5">SKU: KBD-M2-BLK</div>
                    </div>
                    <div class="text-right shrink-0">
                       <div class="text-sm font-bold">342</div>
                       <div class="text-[10px] text-gray-500 font-mono">Units</div>
                       <div class="text-[10px] font-bold text-gray-400 mt-0.5">$41,040</div>
                    </div>
                 </div>

                 <div class="flex items-center gap-4">
                    <div class="w-14 h-14 bg-[#F9F9F9] border border-gray-300 flex items-center justify-center text-[8px] font-mono text-gray-400 shrink-0">IMG_02</div>
                    <div class="flex-1 min-w-0">
                       <div class="text-sm font-bold truncate">Ergonomic Mouse Pro</div>
                       <div class="text-[10px] text-gray-500 font-mono truncate mt-0.5">SKU: MSE-E1-WHT</div>
                    </div>
                    <div class="text-right shrink-0">
                       <div class="text-sm font-bold">215</div>
                       <div class="text-[10px] text-gray-500 font-mono">Units</div>
                       <div class="text-[10px] font-bold text-gray-400 mt-0.5">$18,350</div>
                    </div>
                 </div>

                 <div class="flex items-center gap-4">
                    <div class="w-14 h-14 bg-[#F9F9F9] border border-gray-300 flex items-center justify-center text-[8px] font-mono text-gray-400 shrink-0">IMG_03</div>
                    <div class="flex-1 min-w-0">
                       <div class="text-sm font-bold truncate">Desk Mat Extended</div>
                       <div class="text-[10px] text-gray-500 font-mono truncate mt-0.5">SKU: MAT-X-GRY</div>
                    </div>
                    <div class="text-right shrink-0">
                       <div class="text-sm font-bold">189</div>
                       <div class="text-[10px] text-gray-500 font-mono">Units</div>
                       <div class="text-[10px] font-bold text-gray-400 mt-0.5">$6,470</div>
                    </div>
                 </div>

              </div>
           </div>
        </div>

        <div class="bg-white border border-gray-300 p-6">
           <div class="flex justify-between items-center border-b border-gray-300 pb-4 mb-6">
              <h2 class="text-lg font-black tracking-tight uppercase">TRANSACTION_LOG</h2>
              <button class="border border-black px-4 py-2 text-[10px] font-bold tracking-widest hover:bg-gray-50">VIEW_ALL</button>
           </div>
           
           <div class="overflow-x-auto">
             <table class="w-full text-left text-sm whitespace-nowrap">
                <thead class="text-[10px] font-bold text-gray-500 tracking-widest border-b border-gray-300 uppercase">
                   <tr>
                      <th class="pb-4 font-normal">Order_ID</th>
                      <th class="pb-4 font-normal">Timestamp</th>
                      <th class="pb-4 font-normal">Customer</th>
                      <th class="pb-4 font-normal">Total</th>
                      <th class="pb-4 font-normal">Status</th>
                      <th class="pb-4 font-normal text-right">Action</th>
                   </tr>
                </thead>
                <tbody class="font-mono text-xs">
                   <tr class="border-b border-gray-100 hover:bg-gray-50">
                      <td class="py-4 font-bold text-gray-600">#ORD-9934</td>
                      <td class="py-4 text-gray-400">2023-10-27 14:22Z</td>
                      <td class="py-4 font-sans font-bold text-black">Alice Freeman</td>
                      <td class="py-4 font-bold">$240.00</td>
                      <td class="py-4"><span class="border border-gray-300 text-[9px] px-2 py-1 bg-white font-bold font-sans tracking-widest text-gray-600">PENDING</span></td>
                      <td class="py-4 text-right text-gray-400 hover:text-black cursor-pointer tracking-widest">•••</td>
                   </tr>
                   <tr class="border-b border-gray-100 hover:bg-gray-50">
                      <td class="py-4 font-bold text-gray-600">#ORD-9933</td>
                      <td class="py-4 text-gray-400">2023-10-27 13:45Z</td>
                      <td class="py-4 font-sans font-bold text-black">Marcus Chen</td>
                      <td class="py-4 font-bold">$85.50</td>
                      <td class="py-4"><span class="bg-black text-white text-[9px] px-2 py-1.5 font-bold font-sans tracking-widest">SHIPPED</span></td>
                      <td class="py-4 text-right text-gray-400 hover:text-black cursor-pointer tracking-widest">•••</td>
                   </tr>
                   <tr class="border-b border-gray-100 hover:bg-gray-50">
                      <td class="py-4 font-bold text-gray-600">#ORD-9932</td>
                      <td class="py-4 text-gray-400">2023-10-27 11:12Z</td>
                      <td class="py-4 font-sans font-bold text-black">Sarah Jenkins</td>
                      <td class="py-4 font-bold">$1,250.00</td>
                      <td class="py-4"><span class="bg-black text-white text-[9px] px-2 py-1.5 font-bold font-sans tracking-widest">SHIPPED</span></td>
                      <td class="py-4 text-right text-gray-400 hover:text-black cursor-pointer tracking-widest">•••</td>
                   </tr>
                   <tr class="border-b border-gray-100 hover:bg-gray-50">
                      <td class="py-4 font-bold text-gray-600">#ORD-9931</td>
                      <td class="py-4 text-gray-400">2023-10-26 18:30Z</td>
                      <td class="py-4 font-sans font-bold text-black">David Kim</td>
                      <td class="py-4 font-bold">$45.00</td>
                      <td class="py-4"><span class="border border-gray-200 text-gray-400 text-[9px] px-2 py-1 bg-transparent font-bold font-sans tracking-widest">DELIVERED</span></td>
                      <td class="py-4 text-right text-gray-400 hover:text-black cursor-pointer tracking-widest">•••</td>
                   </tr>
                   <tr class="hover:bg-gray-50">
                      <td class="py-4 font-bold text-gray-600">#ORD-9930</td>
                      <td class="py-4 text-gray-400">2023-10-26 15:05Z</td>
                      <td class="py-4 font-sans font-bold text-black">Elena Rodriguez</td>
                      <td class="py-4 font-bold">$320.00</td>
                      <td class="py-4"><span class="border border-gray-200 text-gray-400 text-[9px] px-2 py-1 bg-transparent font-bold font-sans tracking-widest">DELIVERED</span></td>
                      <td class="py-4 text-right text-gray-400 hover:text-black cursor-pointer tracking-widest">•••</td>
                   </tr>
                </tbody>
             </table>
           </div>
        </div>

      </main>
    </div>

  </div>
</template>