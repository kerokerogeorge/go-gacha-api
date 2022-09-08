<template>
  <div class="fixed h-screen w-48 z-30 top-0 border-r border-solid border-gray-400 bg-white">
    <div class="h-full pt-32">
      <div class="text-gray-600 hover:bg-blue-50 cursor-pointer py-3 px-5" :class="{'current': $route.path === '/user/create'}">
        <NuxtLink :to="{'path':'/user/create'}" class="">
          ユーザー作成
        </NuxtLink>
      </div>
      <div class="text-gray-600 hover:bg-blue-50 cursor-pointer py-3 px-5" :class="{'current': $route.path === '/user/list'}">
        <NuxtLink :to="{'path':'/user/list'}" class="">
          ユーザー一覧
        </NuxtLink>
      </div>
      <div class="text-gray-600 hover:bg-blue-50 cursor-pointer py-3 px-5" :class="{'current': $route.path === '/gacha'}">
        <NuxtLink :to="{'path':'/gacha'}" class="">
          ガチャ一覧
        </NuxtLink>
      </div>
      <div class="text-gray-600 hover:bg-blue-50 cursor-pointer py-3 px-5" :class="{'current': $route.path === '/gacha/draw'}">
        <NuxtLink :to="{'path':'/gacha/draw'}" class="">
          ガチャを引く
        </NuxtLink>
      </div>
      <div class="text-gray-600 hover:bg-blue-50 cursor-pointer py-3 px-5" :class="{'current': $route.path === '/character/create'}">
        <NuxtLink :to="{'path':'/character/create'}" class="">
          キャラクター作成
        </NuxtLink>
      </div>
      <div class="text-gray-600 hover:bg-blue-50 cursor-pointer py-3 px-5" :class="{'current': $route.path === '/character/list'}">
        <NuxtLink :to="{'path':'/character/list'}" class="">
          キャラクター一覧
        </NuxtLink>
      </div>
    </div>
  </div>
</template>

<style scoped>
.current {
  background-color: rgb(229, 252, 241);
}
</style>