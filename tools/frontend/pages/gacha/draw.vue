<template>
  <div class="min-h-screen py-28">
    <div class="z-20 fixed w-full h-auto px-5 border-b border-solid border-gray-400 bg-white pt-3 pb-4">
      <div class="flex items-center">
        <div class="mr-3">
          <div class="text-xs text-gray-500">
            <p>tokenContractAddress: {{ addresses.tokenContractAddress }}</p>
            <p>vendorContractAddress: {{ addresses.vendorContractAddress }}</p>
            <p>myAddress: {{ addresses.myAddress }}</p>
            <p>toAddress: {{ addresses.toAddress }}</p>
          </div>
          <div>
            <div class="text-xs text-gray-600 mt-2">GachaId: {{ gachaId }}</div>
          </div>
          <div class="mt-1">
            <div class="text-xs text-gray-600 mt-2">Token: {{ token }}</div>
          </div>
          <div class="mt-1">
            <div class="text-xs text-gray-600 mt-2">Times（ガチャを引く回数）: {{ times }}</div>
            <input class="p-1 w-72 rounded-sm text-gray-600 border border-solid border-gray-400" type="number" v-model="times">
          </div>
          <div class="text-sm italic">
            <template v-if="!loading && !fetched">
              <p class="text-gray-500">please draw gasha</p>
            </template>
            <template v-else-if="loading && fetched">
              <p class="text-blue-500">fetching...</p>
            </template>
            <template v-else-if="!loading && fetched && !isError">
              <p class="text-green-500">fetch finished</p>
            </template>
            <template v-else-if="!loading && fetched && isError">
              <p class="text-green-500">fetch failed</p>
            </template>
          </div>
        </div>
        <div>
          <button
            class="text-sm cursor-pointer rounded-full h-24 w-24 text-white"
            :class="{'bg-green-300': !gachaId || !token || times === 0, 'bg-green-500 hover:bg-green-700': gachaId && token && times !== 0 }"
            :disabled="!gachaId || !token || times === 0"
            @click="drawGacha"
          >
            ガチャを引く
          </button>
        </div>
      </div>
      <button class="text-xs text-gray-400" @click="externalLink(transaction)">
        transaction details: {{ transaction }}
      </button>
    </div>
    <template v-if="characters.length > 0">
      <div class="pt-52 px-5 w-full">
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
import { mapGetters } from 'vuex'
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
      userCharactersIds: null
    }
  },
  computed: {
    ...mapGetters('gacha', [
      'gachaId',
      'token'
    ]),
  },
  methods: {
    async drawGacha () {
      this.fetched = true
      this.loading = true
      this.isError = false
      try {
        await this.getGachaCharacters()
        const payload = await this.getTransferTokenTransactionPayload()
        const request = await this.createTransaction(payload)
        const tx = await web3.eth.sendTransaction(request)
        await this.updateStatus(this.userCharactersIds, "success")
        this.transaction = `https://goerli.etherscan.io/tx/${tx.transactionHash}`
      } catch (e) {
        this.isError = true
        await this.updateStatus(this.userCharactersIds, "failed")
        console.log(e)
      } finally {
        this.times = 0
        this.loading = false
      }
    },
    externalLink(url) {
      window.open(url, '_blank')
    },
    async getTransferTokenTransactionPayload() {
      try {
        const transferAmount = 0.1
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
    }
  }
}
</script>
