<template>
  <div class="pt-28 mx-10">
    <template v-if="characters">
      <div class="mt-10">
        <p>Characters</p>
      </div>
      <div class="mt-10 border border-solid border-gray-500 text-gray-700 text-sm">
        <div class="flex py-2 border-b border-solid border-gray-500">
          <div class="ml-3 w-10">ID</div>
          <div class="w-28">Name</div>
          <div class="w-28">Image</div>
          <div class="w-48">Created</div>
          <div class="ml-auto w-20"/>
        </div>
        <div v-for="(c, index) in characters" :key="index" class="">
          <div class="flex py-2">
            <div class="ml-3 w-10">{{ c.id }}</div>
            <div class="w-28">{{ c.name }}</div>
            <div class="w-28">
              <div class="relative w-20">
                <img :src="c.imgUrl" alt="pokemon" class="w-20 h-20 border border-solid border-gray-400" />
                <div class="absolute text-xs bottom-2 px-2 w-full bg-gray-400 text-white bg-opacity-80">{{ c.name }}</div>
              </div>
            </div>
            <div class="w-48">{{ $dayjs(c.createdAt).format('YYYY/MM/DD HH:mm') }}</div>
            <div class="ml-auto w-20">
              <button @click="deleteCharacter(c.id)">削除</button>
            </div>
          </div>
        </div>
      </div>
    </template>
  </div>
</template>

<script>
import characterRepository from '~/repositories/characterRepository'

export default {
  data () {
    return {
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
    async fetchCharacters () {
      try {
        const { data } = await characterRepository.getCharacters()
        this.characters = data.data
      } catch (e) {
        console.log(e)
      } finally {}
    },
    async deleteCharacter (characterId) {
      try {
        await characterRepository.delete(characterId)
      } catch (e) {
        console.log(e)
      } finally {
        await this.fetchCharacters()
      }
    }
  }
}
</script>
