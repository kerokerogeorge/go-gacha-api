<template>
  <div class="pt-28 mx-10">
    <div class="">
      <div class="my-3">
        <button class="text-sm cursor-pointer rounded-md py-1 px-3 text-white" :class="{'bg-green-300': !character.name, 'bg-green-500 hover:bg-green-700': character.name }" :disabled="!character.name" @click="createCharacter">キャラクターを作成する</button>
      </div>
      <input class="p-1 rounded-sm text-gray-600 border border-solid border-gray-400" type="text" v-model="character.name">
      <div class="text-xs text-gray-400 mt-2">NAME: {{ character.name }}</div>
    </div>
    <div class="border border-gray-400 border-solid mt-10 h-28">
      <div class="mx-4 my-3 text-gray-600 font-semibold text-lg">Response</div>
      <div v-if="res" class="m-4 text-xs">{{ res }}</div>
    </div>
    <template v-if="characters">
      <div class="mt-10">
        <p>Users</p>
      </div>
      <div v-for="(c, index) in characters" :key="index" class="m-3 flex">
        <div class="mr-3">ID: {{ c.id }}</div>
        <div>Name: {{ c.name }}</div>
      </div>
    </template>
  </div>
</template>

<script>
import characterRepository from '~/repositories/characterRepository'

export default {
  data () {
    return {
      character: {
        name: ''
      },
      characters: null,
      res: null
    }
  },
  async mounted () {
    try {
      await this.fetchCharacters()
    } catch (e) {
      console.log(e)
    }
  },
  methods: {
    async createCharacter () {
      try {
        const { data } = await characterRepository.create(this.character.name)
        this.res = data.data
      } catch (e) {
        console.log(e)
      } finally {
        this.character.name = ''
        await this.fetchCharacters()
      }
    },
    async fetchCharacters () {
      try {
        const { data } = await characterRepository.getCharacters()
        this.characters = data.data
      } catch (e) {
        console.log(e)
      } finally {}
    }
  }
}
</script>
