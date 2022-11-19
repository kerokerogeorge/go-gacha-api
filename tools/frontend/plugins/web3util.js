import Web3 from 'web3';
const web3 = new Web3(Web3.givenProvider);
import { tokenABI, vendorABI } from '../vendorAbi';

class Web3Util {
  static async getAccounts() {
    return await web3.eth.getAccounts();
  }

  static async getBalance(address) {
    try {
      return await web3.eth.getBalance(address, (error, balance) => {
        return web3.utils.toBN(balance).toString()
      });
    } catch (e) {
      console.error(e)
    }
  }

  static async fromWei(balance) {
    return await web3.utils.fromWei(balance, 'ether')
  }

  static async createContract(abi, contractAddress) {
    return new web3.eth.Contract(abi, contractAddress);
  }

  static tokenContract(contractAddress) {
    console.log('contractAddress')
    console.log(contractAddress)
    const tokenContract = this.createContract(tokenABI, contractAddress)
    return tokenContract
  }

  // static test2 () {
  //   const test = this.test()
  //   console.log(test)
  // }
}

export default (context, inject) => {
  inject('web3util', Web3Util)
}