<template>
  <div class="min-h-screen pt-28">
    <div class="fixed z-20 w-full h-52 px-5 border-b border-solid border-gray-400 bg-white">
      <div class="">
        <div class="my-3">
          <button class="bg-green-500 hover:bg-green-700 text-sm cursor-pointer rounded-md py-1 px-3 text-white" @click="createGacha">ガチャを作成する</button>
        </div>
      </div>
      <div class="border border-gray-400 border-solid mt-5 h-28">
        <div class="px-4 my-3 text-gray-600 font-semibold text-lg">Response</div>
        <div v-if="res" class="m-4 text-xs">{{ res }}</div>
      </div>
    </div>
    <template v-if="gachas">
      <div class="pt-52 px-5 w-full flex h-full">
        <div class="w-5/12 pt-3 pb-10 border-r border-solid border-gray-400">
          <p class="font-semibold text-gray-700">ガチャ一覧</p>
          <div v-for="(g, index) in gachas" :key="index" class="mt-2 text-xs">
            <div class="mr-3">ID: {{ g.gachaId }}</div>
            <div class="flex">
              <div class="mr-2">
                <button class="mt-2 py-1 px-2 text-xs rounded-md text-white bg-green-400 hover:bg-green-600" @click="fetchGacha(g.gachaId)">キャラを表示</button>
              </div>
              <div>
                <button class="mt-2 py-1 px-2 text-xs rounded-md text-white bg-red-400 hover:bg-red-600" @click="deleteGacha(g.gachaId)">削除</button>
              </div>
            </div>
          </div>
        </div>
        <div class="w-7/12">
          <div class="z-20 fixed top-80 flex bg-white items-center h-12 px-5 w-full border-b border-gray-400 border-solid">
            <div>
              <p class="text-gray-700 text-xs">ガチャID: {{ gachaId ? gachaId : '-' }}</p>
            </div>
          </div>
          <template v-if="characters.length > 0">
            <div class="px-5 sticky top-96 pt-14 pb-20">
              <div class="border border-solid border-gray-500 text-gray-700 text-sm">
                <div class="flex py-2 border-b border-solid border-gray-500">
                  <div class="ml-3 w-10">ID</div>
                  <div class="w-48">Name</div>
                  <div class="w-auto">Emmition Rate</div>
                </div>
                <div v-for="(c, index) in characters" :key="index" class="">
                  <div class="flex py-2">
                    <div class="ml-3 w-10">{{ c.characterId }}</div>
                    <div class="w-48">{{ c.name }}</div>
                    <div class="w-28">{{ c.emissionRate }} / {{ multiple }} =</div>
                    <div class="w-auto">{{ c.fixedRate ? c.fixedRate : '-' }}%</div>
                  </div>
                </div>
                <div class="border-t border-solid border-gray-500">
                  <div class="flex py-2">
                    <div class="w-60 pl-3">合計</div>
                    <div class="w-28">{{total}} / 100 = {{ multiple }}</div>
                    <div class="w-auto">{{ totalFixedRate }}%</div>
                  </div>
                </div>
              </div>
            </div>
          </template>
        </div>
      </div>
    </template>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
import gachaRepository from '~/repositories/gachaRepository'

export default {
  data () {
    return {
      characters: [],
      res: null,
      total: 0,
      emmitionRates: [],
      multiple: 0,
      isFetching: false,
      totalFixedRate: 0
    }
  },
  computed: {
    ...mapGetters('gacha', [
      'gachaId',
      'gachas'
    ]),
  },
  methods: {
    ...mapActions('gacha', [
      'fetchGachas'
    ]),
    async createGacha () {
      try {
        const { data } = await gachaRepository.create()
        this.res = data
      } catch (e) {
        console.log(e)
      } finally {
        await this.fetchGachas()
      }
    },
    async fetchGacha (gachaId) {
      try {
        this.reset()
        this.isFetching = true
        const { data } = await gachaRepository.getGacha(gachaId)
        this.characters = data.data.characters
        this.characters.forEach(c => {
          this.total += c.emissionRate
        })
        this.multiple = this.total / 100
        this.characters = this.characters.map(c => {
          const data = c.emissionRate / this.multiple
          c.fixedRate = parseFloat(data.toFixed(2))
          this.totalFixedRate += parseFloat(data.toFixed(2))
          return c
        })
        await this.selectGachaId({ gachaId: gachaId })
      } catch (e) {
        console.log(e)
      } finally {
        this.isFetching = false
      }
    },
    async deleteGacha (gachaId) {
      try {
        await gachaRepository.delete(gachaId)
        this.characters = []
      } catch (e) {
        console.log(e)
      } finally {
        await this.fetchGachas()
      }
    },
    reset () {
      this.total = 0
      this.totalFixedRate = 0
      this.multiple = 0
    }
  }
}
</script>
