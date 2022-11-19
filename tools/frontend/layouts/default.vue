<template>
  <div class="w-screen">
    <UtilLoading v-if="loading" />
    <div class="flex items-center pl-56 h-28 fixed z-20 bg-white border-b border-solid border-gray-400 w-full">
      <div>
        <div class="flex">
          <div class="p-2 bg-green-50 rounded-lg text-xs mx-2">
            <p v-if="tokenName && symbol" class="text-sm font-semibold">{{ tokenName }}:{{ symbol }}</p>
            <p :class="{'highlight': highlight}">自分のトークン残高: {{ myTokenBalance ? myTokenBalance : '-' }}</p>
            <p>ガチャコントラクト保有トークン: {{ gachaVendorTokenBalance ? gachaVendorTokenBalance : '-' }}</p>
            <p>ガチャマスターウォレット保有トークン: {{ gachaWalletTokenBalance ? gachaWalletTokenBalance : '-' }}</p>
          </div>
          <div class="p-2 bg-blue-50 rounded-lg text-xs mx-2">
            <p>自分のETH残高: {{ myEtherBalance ? myEtherBalance : '-' }}</p>
            <p>ガチャコントラクト保有ETH: {{ gachaVendorEtherBalance ? gachaVendorEtherBalance : '-' }}</p>
          </div>
           <div class="p-2 bg-gray-100 rounded-lg text-xs">
            <p class="truncate">me: {{ myWalletAddress ? myWalletAddress : '-' }}</p>
            <p class="truncate">token contract: {{ tokenContractAddress ? tokenContractAddress : '-' }}</p>
            <p class="truncate">gacha contract: {{ vendorContractAddress ? vendorContractAddress : '-' }}</p>
            <p class="truncate">gacha wallet: {{ gachaWalletAddress ? gachaWalletAddress : '-' }}</p>
          </div>
        </div>
      </div>
    </div>
    <div class="relative">
      <CommonSidebar />
    </div>
    <main class="pl-48">
      <Nuxt />
    </main>
  </div>
</template>

<style scoped>
</style>

<script>
import { mapGetters, mapActions } from 'vuex'
import { tokenABI, vendorABI } from '../vendorAbi';
export default {
  data() {
    return {
      loading: false,
      vendorAbi: vendorABI,
      tokenAbi: tokenABI,
    }
  },
  computed: {
    ...mapGetters('token', [
      'isConnected'
    ]),
    ...mapGetters('gacha', [
      'gachaId',
      'token'
    ]),
    ...mapGetters('web3', [
      'tokenContractAddress',
      'vendorContractAddress',
      'gachaWalletAddress',
      'myWalletAddress',
      'myTokenBalance',
      'gachaVendorTokenBalance',
      'gachaWalletTokenBalance',
      'symbol',
      'tokenName',
      'myEtherBalance',
      'gachaVendorEtherBalance',
      'tokenContract',
      'highlight'
    ]),
  }
}
</script>

<style lang="scss" scoped>
.highlight {
  color: red;
  font-weight: 700;
  animation: fadein-anim 6s linear forwards;
}

@keyframes fadein-anim {
  100% {
    font-weight: 400;
    color: black;
  }
}
</style>