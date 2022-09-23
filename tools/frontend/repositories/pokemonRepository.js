import repository from './pokemonBaseRepository'

export default {
  getPokemon (id) {
    return repository.get(`/pokemon/${id}`)
  }
}