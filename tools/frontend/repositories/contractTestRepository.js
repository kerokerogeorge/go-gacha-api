import repository from './repository'

const resource = 'test'

export default {
  createSendERC20Transaction (params) {
    return repository.post(`${resource}/createSendERC20Transaction`, params)
  },
  createBuyTokenTransaction (params) {
    return repository.post(`${resource}/createBuyTokenTransaction`, params)
  },
  getTokenTransferTransactionPayload (params) {
    return repository.post(`${resource}/payload`, params)
  }
}