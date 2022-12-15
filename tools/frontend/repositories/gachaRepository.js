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
  draw (token, gachaId, params) {
    return repository.post(`${resource}/draw/${gachaId}`,
    params,
    {
      headers: {
        'x-token': `${token}`
      }
    })
  },
  update (token,  params) {
    return repository.put(`/user_character`,
    params,
    {
      headers: {
        "x-token": token
      }
    },
    )
  },
  drawWithTransaction (token, gachaId, params) {
    console.log('called!!')
    return repository.post(`${resource}/draw_with_transaction/${gachaId}`,
    params,
    {
      headers: {
        'x-token': `${token}`
      }
    })
  },
}
