import repository from './repository'

const resource = 'contract'

export default {
  getTransferTokenTransactionPayload (params) {
    console.log(params)
    return repository.post(`${resource}/transfer`, params)
  },
}
