<template>
  <div class="min-h-screen py-28">
    <div class="z-20 fixed w-full h-auto px-5 border-b border-solid border-gray-400 bg-white pt-3 pb-4">
      <div class="">
        <div class="mr-3 flex">
          <div v-if="users" class="w-64 mr-5">
            <div class="text-xs text-gray-600 mt-2">user: {{ placeholder.user }}</div>
            <div class="select" :class="{'is-open-user': isOpen.user}">
              <span class="placeholder" @click="openToggle('USER')">{{ placeholder.user }}</span>
              <ul v-for="(u, index) in users" :key="index">
                <li @click="selectUser(u)">{{ u.name }}</li>
              </ul>
            </div>
          </div>
          <div v-if="gachas" class="w-64 mr-5">
            <div class="text-xs text-gray-600 mt-2">gacha: {{ placeholder.gacha }}</div>
            <div class="select" :class="{'is-open-gacha': isOpen.gacha}">
              <span class="placeholder" @click="openToggle('GACHA')">{{ placeholder.gacha }}</span>
              <ul v-for="(g, index) in gachas" :key="index">
                <li @click="selectGacha(g)">{{ g.gachaId }}</li>
              </ul>
            </div>
          </div>
          <div class="mt-1">
            <div class="text-xs text-gray-600 mt-2">Times（ガチャを引く回数）: {{ times }}</div>
            <input class="p-1 w-72 rounded-sm text-gray-600 border border-solid border-gray-400" type="number" v-model="times">
          </div>
        </div>
        <div class="flex mt-3">
          <div class="mr-2">
            <button class="button button-blue" data-ripple-light="true" @click="recieveToken">RECEIVE TOKEN</button>
          </div>
          <div>
            <button class="button button-green" data-ripple-light="true" :disabled="!gachaId || !token || times === 0" @click="drawGacha">draw gacha</button>
          </div>
        </div>
      </div>
      <div class="mt-3 text-sm italic">
        <template v-if="!loading && !fetched">
          <p class="text-gray-500">please draw gacha</p>
        </template>
        <template v-else-if="loading && fetched">
          <p class="text-blue-500">fetching...</p>
        </template>
        <template v-else-if="!loading && fetched && !isError">
          <p class="text-green-500">fetch finished</p>
          <button class="text-xs" @click="externalLink(transaction)">
            transaction detail: <span class="text-gray-700 hover:text-yellow-700 font-semibold">{{ transaction }}</span>
          </button>
        </template>
        <template v-else-if="!loading && fetched && isError">
          <p class="text-green-500">fetch failed</p>
        </template>
      </div>
    </div>
    <template v-if="characters.length > 0">
      <div class="pt-60 px-5 w-full">
        <div class="mt-10 border border-solid border-gray-500 text-gray-700 text-sm">
          <div class="flex py-2 border-b border-solid border-gray-500">
            <div class="ml-3 w-10" />
            <div class="w-1/12">ID</div>
            <div class="w-1/12">CharacterID</div>
            <div class="w-2/12">Name</div>
            <div class="w-2/12">Emmition Rate</div>
            <div class="w-3/12">Status</div>
          </div>
          <div v-for="(c, index) in characters" :key="index" class="">
            <div class="flex py-1 items-center">
              <div class="ml-3 w-10">{{ index + 1 }}</div>
              <div class="w-1/12">{{ c.userCharacterId }}</div>
              <div class="w-1/12">{{ c.characterId }}</div>
              <div class="w-2/12">{{ c.name }}</div>
              <div class="w-2/12">{{ c.emissionRate }}%</div>
              <div class="w-3/12">{{ c.status }}</div>
              <div class="relative border border-solid border-gray-400">
                <img :src="c.imgUrl" alt="pokemon" class="w-20 h-20" />
                <div class="absolute text-xs bottom-2 px-2 w-full bg-gray-400 text-white bg-opacity-80">{{ c.name }}</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </template>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
import { vendorABI } from '../../vendorAbi';
import gachaRepository from '~/repositories/gachaRepository'
import smartContractRepository from '~/repositories/smartContractRepository'
import Web3 from 'web3';
const web3 = new Web3(Web3.givenProvider);

const tokenContractAddress = '0x6a7edAd9c7f49Bf215Add73e5d8F8Cb550177297'
const vendorContractAddress = '0x7dBa3cc9bDf7B3F79dcDD90B0c19768190a5aC5b'
const myAddress = '0x6941cee0e87cb8ABE7A1985bf24c4f54CFeE9785'
const toAddress = '0xec64414617F2B65bB4a7adD57e82a1c5CF53B328'

export default {
  data () {
    return {
      characters: [],
      times: 0,
      loading: false,
      fetched: false,
      transaction: null,
      addresses: {
        tokenContractAddress: tokenContractAddress,
        vendorContractAddress: vendorContractAddress,
        myAddress: myAddress,
        toAddress: toAddress
      },
      userCharactersIds: null,
      vendorAbi: vendorABI,
      contract: null,
      isOpen: {
        user: false,
        gacha: false
      },
      placeholder: {
        user: 'select user',
        gacha: 'select gacha'
      },
      highlight: false
    }
  },
  computed: {
    ...mapGetters('gacha', [
      'gachaId',
      'token',
      'gachas'
    ]),
    ...mapGetters('user', [
      'users',
    ]),
    ...mapGetters('web3', [
      'tokenContract',
      'myWalletAddress',
      'vendorContractAddress',
      'gachaWalletAddress'
    ]),
  },
  async mounted () {
    try {
      this.contract = new web3.eth.Contract(this.vendorAbi, vendorContractAddress);
    } catch (e) {
      console.log(e)
    }
  },
  methods: {
    ...mapActions('gacha', [
      'selectGachaId',
      'setUserToken',
    ]),
    ...mapActions('web3', [
      'setTokenBalance'
    ]),
    async drawGacha () {
      this.fetched = true
      this.loading = true
      this.isError = false
      try {
        await this.getGachaCharacters()
        const payload = await this.getTransferTokenTransactionPayload()
        const request = await this.createTransaction(payload)
        const tx = await web3.eth.sendTransaction(request)
        console.log('tx', tx)
        await this.updateStatus(this.userCharactersIds, "success")
        this.transaction = `https://goerli.etherscan.io/tx/${tx.transactionHash}`
      } catch (e) {
        this.isError = true
        await this.updateStatus(this.userCharactersIds, "failed")
      } finally {
        await this.fetchTokenBalance()
        this.times = 0
        this.loading = false
        if (!this.isError) {
          await this.$store.commit('web3/setHighlight', true)
          setTimeout(async () => {
            await this.$store.commit('web3/setHighlight', false)
          }, 6000);
        }
      }
    },
    externalLink(url) {
      window.open(url, '_blank')
    },
    async getTransferTokenTransactionPayload() {
      try {
        const transferAmount = 1 * this.times
        console.log('transferAmount: ', transferAmount)
        const req = {
          fromAddress: myAddress,
          toAddress: toAddress,
          contractAddress: tokenContractAddress,
          amount: Number(web3.utils.toWei(transferAmount.toString(), "ether"))
        }
        const { data } = await smartContractRepository.getTransferTokenTransactionPayload(req)
        return data.transactionPayload
      } catch (e) {
        console.log(e)
        alert(e.response.data.error);
      }
    },
    async createTransaction(payload) {
      try {
        const accounts = await web3.eth.getAccounts();
        const maxFeePerGas = Number(web3.utils.toBN(payload.maxFeePerGas).toString())
        const maxPriorityFeePerGas = Number(web3.utils.toBN(payload.maxPriorityFeePerGas).toString())
        const request = {
          from: accounts[0],
          to: payload.to,
          gas: payload.gas,
          value: payload.value,
          maxFeePerGas:
            maxFeePerGas > maxPriorityFeePerGas ? maxFeePerGas : maxPriorityFeePerGas,
          maxPriorityFeePerGas: maxPriorityFeePerGas,
          nonce: payload.nonce,
          chainId: payload.chainId,
          input: payload.input
        }
        console.log("request: ", request)
        return request
        // 参考
        // https://goerli.etherscan.io/tx/0x2c74a240ca53e6411a33a0a1def610ae4855c3d7bcb9184b243342507225e713
      } catch (e) {
        console.log(e)
      }
    },
    async getGachaCharacters() {
      try {
        const { data } = await gachaRepository.draw(this.token, this.gachaId, { times: Number(this.times) })
        this.characters = data.result
        this.userCharactersIds = data.result.map(uc => uc.userCharacterId)
      } catch (e) {
        console.log(e)
      }
    },
    async updateStatus(userCharactersIds, status) {
      try {
        const params = {
          status: status,
          userCharacterIds: userCharactersIds
        }
        const { data } = await gachaRepository.update(
          this.token,
          params
        )
        this.characters.map(c => {
          c.status = data.results[0].status
        })
      } catch (e) {
        console.log(e)
      }
    },
    async recieveToken() {
      try {
        this.loading = true
        const payload = await this.getPayloadForTokenPurchase()
        const request = await this.createTransaction(payload)
        const tx = await web3.eth.sendTransaction(request)
        this.transaction = `https://goerli.etherscan.io/tx/${tx.transactionHash}`
      } catch (err) {
        this.isError = true
        console.error(err);
        alert("Error purchasing tokens");
      } finally {
        this.loading = false
      }
    },
    async getPayloadForTokenPurchase() {
      try {
        const req = {
          fromAddress: myAddress,
          contractAddress: vendorContractAddress,
        }
        const { data } = await smartContractRepository.getBuyTokenTransactionPayload(req)
        return data.transactionPayload
      } catch (err) {
        console.error(err);
      }
    },
    async selectUser (user) {
      this.placeholder.user = user.name
      this.isOpen.user = false
      await this.setUserToken({ token: user.token })
    },
    async selectGacha (gacha) {
      this.placeholder.gacha = gacha.gachaId
      this.isOpen.gacha = false
      await this.selectGachaId({ gachaId: gacha.gachaId })
    },
    openToggle (type) {
      if (type === 'USER') {
        this.isOpen.user = !this.isOpen.user
      }
      if (type === 'GACHA') {
        this.isOpen.gacha = !this.isOpen.gacha
      }
    },
    async fetchTokenBalance() {
      const myTokenBalance = await this.tokenContract.balanceOf(this.myWalletAddress).call()
      const vendorContractBalance = await this.tokenContract.balanceOf(this.vendorContractAddress).call()
      const gachaWalletBalance = await this.tokenContract.balanceOf(this.gachaWalletAddress).call()
      await this.setTokenBalance({ me: true, balance: myTokenBalance })
      await this.setTokenBalance({ gachaVendor: true, balance: vendorContractBalance })
      await this.setTokenBalance({ gachaWallet: true, balance: gachaWalletBalance })
    }
  }
}
</script>


<style lang="scss" scoped>

.select{
  position: relative;
  display: block;
  margin: 0 auto;
  width: 100%;
  max-width: 300px;
  color: #cccccc;
  text-align: left;
  // user-select: none;
  -webkit-touch-callout: none;

  .placeholder{
    position: relative;
    display: block;
    background-color: #393d41;
    z-index: 1;
    padding: 1em;
    border-radius: 2px;
    font-size: 12px;
    cursor: pointer;

    &:hover{
      background: darken(#393d41,2%);
    }

    &:after{
      position: absolute;
      right: 1em;
      top: 50%;
      transform: translateY(-50%);
      font-family: 'FontAwesome';
      content: '\f078';
      z-index: 10;
    }
  }

  &.is-open-gacha{
    .placeholder:after{
      content: '\f077';
    }
    ul{
      display: block;
    }
  }

  &.is-open-user{
    .placeholder:after{
      content: '\f077';
    }
    ul{
      display: block;
    }
  }

  &.select--white{
    .placeholder{
      background: #fff;
      color: #999;
      &:hover{
        background: darken(#fff,2%);
      }
    }
  }

  ul{
    display: none;
    position: absolute;
    overflow: hidden;
    overflow-y: auto;
    width: 100%;
    background: rgb(235, 235, 235);
    border-radius: 2px;
    top: 100%;
    left: 0;
    list-style: none;
    margin: 5px 0 0 0;
    padding: 0;
    z-index: 100;
    max-height: 100px;

    li{
      display: block;
      text-align: left;
      padding: 0.5em 1em 0.5em 1em;
      color: #999;
      cursor: pointer;
      font-size: 12px;

      &:hover{
        background: #4ebbf0;
        color: #fff;
      }
    }
  }
}
</style>