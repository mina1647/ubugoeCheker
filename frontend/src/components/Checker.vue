<script setup lang="ts">
import { ref } from 'vue'

const data = ref<string[]>()
const query = ref<string>('')
const imageUrlRef = ref('https://example.com/usericon.png')

const getUbugoe = async (name: string) => {
  const response = await fetch(`/api/ubugoe/${name}`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
  })
  if (!response.ok) {
    data.value = ['いないよ']
    return
  }
  imageUrlRef.value = `https://q.trap.jp/api/v3/public/icon/${name}`

  data.value = await response.json()
}
</script>

<template>
  <div class="checker">
    <h1>うぶごえチェッカー</h1>
    <!-- <p>ここにうぶごえチェッカーの内容が入ります。</p> -->
    <input type="text" v-model="query" placeholder="traQIDを入力してください" class="nameinput" />
    <button @click="getUbugoe(query)">検索</button>

    <div v-for="item in data" :key="item" class="item">
      <div class="icon">
        <img :src="imageUrlRef" class="usericon" />
      </div>

      <div class="text">
        {{ item }}
      </div>
    </div>
  </div>
</template>

<style>
.checker {
  background-color: #fff;
  color: #000;
  width: 100%;
  height: 100vh;
  top: 10vh;
  left: 0;
  position: fixed;
  padding: 1rem;
}

.nameinput {
  width: 30%;
}

/* .answer{
  width: 100%;
  height: 100%;

} */

.item {
  display: flex;
  width: 100%;
  padding: 1rem;
  justify-content: center;
  align-items: center;
}

.text {
  /* display: flex; */
  width: 85%;
  height: 100%;
  /* text-align: center; */
}

.icon {
  /* display: flex; */
  top: 0;
  width: 4rem;
  height: 100%;
  padding: 0.5rem;
}

.usericon {
  border-radius: 50%;
  width: 3rem;
  height: 3rem;
}
</style>
