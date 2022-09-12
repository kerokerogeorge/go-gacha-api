import repository from './repository'

const resource = 'character'

export default {
  getCharacters () {
    return repository.get(`${resource}/list`)
  },
  create (name) {
    return repository.post(`${resource}`, {
      name
    })
  },
  getCharactersWithEmmitionRates (gachaId) {
    return repository.get(`${resource}/emmition_rates`, {
      gachaId
    })
  },
  delete (characterId) {
    return repository.delete(`${resource}/${characterId}`)
  },
}