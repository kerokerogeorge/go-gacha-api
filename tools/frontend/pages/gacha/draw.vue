<template>
  <div class="min-h-screen py-28">
    <div class="z-20 fixed w-full h-auto px-5 border-b border-solid border-gray-400 bg-white py-10">
      <div class="flex items-center">
        <div class="mr-3">
          <div>
            <div class="text-xs text-gray-600 mt-2">GachaId: {{ gachaId }}</div>
          </div>
          <div class="mt-1">
            <div class="text-xs text-gray-600 mt-2">Token: {{ token }}</div>
          </div>
          <div class="mt-1">
            <div class="text-xs text-gray-600 mt-2">Times（ガチャを引く回数）: {{ times }}</div>
            <input class="p-1 w-96 rounded-sm text-gray-600 border border-solid border-gray-400" type="number" v-model="times">
          </div>
        </div>
        <div>
          <button
            class="text-sm cursor-pointer rounded-full h-32 w-32 text-white"
            :class="{'bg-green-300': !gachaId || !token || times === 0, 'bg-green-500 hover:bg-green-700': gachaId && token && times !== 0 }"
            :disabled="!gachaId || !token || times === 0"
            @click="drawGacha"
          >
            ガチャを引く
          </button>
        </div>
      </div>
      <button class="text-xs text-gray-400" @click="externalLink(transaction)">
        transaction details: {{ transaction }}
      </button>
    </div>
    <template v-if="characters.length > 0">
      <div class="pt-52 px-5 w-full">
        <div class="mt-10 border border-solid border-gray-500 text-gray-700 text-sm">
          <div class="flex py-2 border-b border-solid border-gray-500">
            <div class="ml-3 w-10" />
            <div class="w-10">ID</div>
            <div class="w-48">Name</div>
            <div class="w-40">Emmition Rate</div>
          </div>
          <div v-for="(c, index) in characters" :key="index" class="">
            <div class="flex py-1 items-center">
              <div class="ml-3 w-10">{{ index + 1 }}</div>
              <div class="w-10">{{ c.characterId }}</div>
              <div class="w-48">{{ c.name }}</div>
              <div class="w-40">{{ c.emissionRate }}%</div>
              <div class="relative border border-solid border-gray-400">
                <img :src="c.imgUrl" alt="pokemon" class="w-20 h-20" />
                <div class="absolute text-xs bottom-2 px-2 w-full bg-gray-400 text-white bg-opacity-80">{{ c.name }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </template>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import gachaRepository from '~/repositories/gachaRepository'

export default {
  data () {
    return {
      characters: [],
      times: 0,
      loading: false,
      transaction: null
    }
  },
  computed: {
    ...mapGetters('gacha', [
      'gachaId',
      'token'
    ]),
  },
  methods: {
    async drawGacha () {
      // this.loading = true
      try {
        const { data } = await gachaRepository.draw(this.token, this.gachaId, Number(this.times))
        console.log(data)
        this.characters = data.result
        this.transaction = `https://goerli.etherscan.io/tx/${data.transaction}`
      } catch (e) {
        console.log(e)
      } finally {
        // this.loading = false
        this.times = 0
      }
    },
    externalLink(url) {
      window.open(url, '_blank')
    }
  }
}
</script>
