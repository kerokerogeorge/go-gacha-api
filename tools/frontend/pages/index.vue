<template>
  <div class="min-h-screen pt-28">
    <div class="px-5 py-10">
      <div class="border border-solid gorder-gray-600">
        <div class="flex py-2 border-b border-solid border-gray-600">
          <div class="w-1/3 pl-3">API</div>
          <div class="w-28">METHOD</div>
          <div class="w-auto">説明</div>
        </div>
        <div class="text-gray-500 text-sm">
          <div class="flex py-2"><div class="w-1/3 pl-2">/user/list</div><div class="w-28">GET</div><div>ユーザー一覧取得</div></div>
          <div class="flex py-2"><div class="w-1/3 pl-2">/user</div><div class="w-28">POST</div><div>ユーザ-作成</div></div>
          <div class="flex py-2"><div class="w-1/3 pl-2">/user</div><div class="w-28">GET</div><div>ユーザ-取得</div></div>
          <div class="flex py-2"><div class="w-1/3 pl-2">/user/characters</div><div class="w-28">GET</div><div>ユーザ所持キャラクター一覧取得</div></div>
          <div class="flex py-2"><div class="w-1/3 pl-2">/character/list</div><div class="w-28">GET</div><div>キャラクター一覧取得</div></div>
          <div class="flex py-2"><div class="w-1/3 pl-2">/character</div><div class="w-28">POST</div><div>キャラクター作成</div></div>
          <div class="flex py-2"><div class="w-1/3 pl-2">/character/emmition_rates</div><div class="w-28">GET</div><div>ガチャの持つキャラクターを排出率とともに取得するAPI</div></div>
          <div class="flex py-2"><div class="w-1/3 pl-2">/character/:characterId</div><div class="w-28">DELETE</div><div>キャラクター削除</div></div>
          <div class="flex py-2"><div class="w-1/3 pl-2">/gacha/list</div><div class="w-28">GET</div><div>キャラクター一覧取得</div></div>
          <div class="flex py-2"><div class="w-1/3 pl-2">/gacha</div><div class="w-28">POST</div><div>ガチャ作成</div></div>
          <div class="flex py-2"><div class="w-1/3 pl-2">/gacha/:gachaId</div><div class="w-28">GET</div><div>一件のガチャ取得</div></div>
          <div class="flex py-2"><div class="w-1/3 pl-2">/gacha/draw/:gachaId</div><div class="w-28">POST</div><div>ガチャ実行</div></div>
          <div class="flex py-2"><div class="w-1/3 pl-2">/:gachaId</div><div class="w-28">DELETE</div><div>ガチャ削除</div></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import userRepository from '~/repositories/userRepository'

export default {
  data() {
    return {
      res: null
    }
  },
  async mounted () {
    try {
      await this.fetchUsers()
    } catch (e) {
      console.log(e)
    }
  },
  methods: {
    async fetchUsers() {
      try {
        const { data } = await userRepository.getUsers()
        this.res = data
      } catch (e) {
        console.log(e)
      }
    }
  }
}
</script>
