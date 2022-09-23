import repository from './repository'

export default {
  getUsers () {
    return repository.get('/user/list')
  },
  getUser (token) {
    return repository.get('/user', {
      headers: {
        'x-token': `${token}`
      }
    })
  },
  createUser (name, address) {
    return repository.post('/user', {
      name,
      address
    })
  },
  getUserCharacters (token) {
    return repository.get('/user/characters', {
      headers: {
        'x-token': `${token}`
      }
    })
  },
  update (token, params) {
    return repository.put('/user', {
      headers: {
        'x-token': `${token}`
      }
    },
    { params })
  },
  delete (token) {
    return repository.delete('/user', {
      headers: {
        'x-token': `${token}`
      }
    })
  }
}