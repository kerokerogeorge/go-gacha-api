import repository from './repository'

const resource = 'gacha'

export default {
  getGachas () {
    return repository.get(`${resource}/list`)
  },
  getGacha (gachaId) {
    return repository.get(`${resource}/${gachaId}`)
  },
  create () {
    return repository.post(`${resource}`)
  },
  delete (gachaId) {
    return repository.delete(`${resource}/${gachaId}`)
  },
  draw (token, gachaId, times) {
    return repository.post(`${resource}/draw/${gachaId}`, {
      times: times,
    },
    {
      headers: {
        "x-token": token
      }
    })
  }
}
