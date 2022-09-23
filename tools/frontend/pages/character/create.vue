<template>
  <div class="pt-28 px-10 w-full">
    <div class="">
      <div class="my-3">
        <button class="text-sm cursor-pointer rounded-md py-1 px-3 text-white bg-green-500 hover:bg-green-700" @click="createCharacter">キャラクターを作成する</button>
      </div>
    </div>
    <div class="border border-gray-400 border-solid mt-10 h-28">
      <div class="mx-4 my-3 text-gray-600 font-semibold text-lg">Response</div>
      <div v-if="res" class="m-4 text-xs">{{ res }}</div>
    </div>
    <div class="border border-gray-400 border-solid mt-10 h-28">
      <div class="mx-4 my-3 text-gray-600 font-semibold text-lg">Pokemon Response</div>
      <div v-if="pokemon" class="m-4 text-xs">{{ pokemon }}</div>
    </div>
    <template v-if="characters">
      <div class="mt-10">
        <p>Users</p>
      </div>
      <div class="flex flex-wrap">
        <div v-for="(c, index) in characters" :key="index" class="m-2 relative text-xs text-gray-500 border border-solid border-gray-500 rounded-lg">
          <img :src="c.imgUrl" alt="pokemon" class="w-28 h-28" />
          <div class="absolute bottom-0 left-0 p-1 bg-gray-400 text-white bg-opacity-80">{{ c.name }}</div>
        </div>
      </div>
    </template>
  </div>
</template>

<script>
import characterRepository from '~/repositories/characterRepository'
import pokemonRepository from '~/repositories/pokemonRepository'

export default {
  data () {
    return {
      character: {
        name: ''
      },
      characters: null,
      res: null,
      pokemonCount: 500,
      pokemon: {
        name: null,
        imgUrl: null
      },
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
        await this.fetchPokemon()
        console.log(this.pokemon)
        const { data } = await characterRepository.create(this.pokemon.name, this.pokemon.imgUrl)
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
    },
    async fetchPokemon () {
      try {
        const pokeId = Math.floor(Math.random() * this.pokemonCount)
        const { data } = await pokemonRepository.getPokemon(pokeId)
        Object.assign(this.pokemon, {
          name: data.name,
          imgUrl: data.sprites.front_default,
        })
      } catch (e) {
        console.log(e)
      }
    }
  }
}
</script>
