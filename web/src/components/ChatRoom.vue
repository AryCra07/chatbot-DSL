<template>
  <div class="chat-room">
    <div class="header">
      <el-avatar
          class="avatar"
          shape="square"
          :size="50"
          :src="require('../assets/avatar.jpg')"
      />
      <div class="bot-name">
        <span class="title">{{ this.botName }}</span>
        <span class="subtitle">{{ this.serviceMode }}</span>
      </div>
      <div class="control">
        <span class="back" @click="goback">{{ $t('back') }}</span>
      </div>
    </div>

    <div class="main">
      <el-scrollbar ref="historyScrollbar" class="chat">
        <div
            ref="historyContent"
            style=""
        >
          <div
              v-for="msg in history"
              :key="msg.id"
              :class="'message' + ' ' + (msg.bot) "
          >
          </div>
        </div>
      </el-scrollbar>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent} from 'vue';
import {MessagePacket} from '@/types/ChatTypes';
export default defineComponent({
  name: 'ChatBox',
  data() {
    return {
      isAvailable: true,
      message: '',
      latestMessageID: 0,
      messageHistory: [] as MessagePacket[],
    };
  },
  setup() {
    return null;
  },
  props: {
    botName: {type: String, default: ''},
    serviceMode: {type: String, default: ''},
    schema: {type: String, default: ''}
  },
  mounted() {
    let i = 0;
    return i;
  },
  methods: {}
})
</script>

<style scoped>

</style>
