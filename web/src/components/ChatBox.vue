<template>
  <div class="chat-box">
    <div class="header">
      <el-avatar
          class="avatar"
          shape="square"
          :size="50"
          :src="require('../assets/avatar.jpg')"
      />
      <div class="user-name">
        Name
      </div>
      <div class="control">
        <span class="back" @click="goback">{{ $t('back') }}</span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent} from "vue";


export default defineComponent({
  name: "ChatBox",
  data() {
    return {};
  },
  setup() {
    return null;
  },
  props: {},
  mounted() {
  },
  methods: {}
})
</script>

<style scoped>

</style>
