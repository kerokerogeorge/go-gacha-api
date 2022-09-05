import repository from './repository'

export default {
  getUsers () {
    return repository.get('/user/list')
  }
}