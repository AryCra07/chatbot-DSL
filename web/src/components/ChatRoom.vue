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
        <span class="title">
          {{ this.botName }}
        </span>
        <span class="subtitle">
          {{ this.serviceMode }}
        </span>
      </div>
      <div class="control">
        <span
          class="back"
          @click="goBack"
        >
          {{ $t('back') }}
        </span>
      </div>
    </div>

    <div class="main">
      <el-scrollbar
        ref="historyScrollbar"
        class="chat"
      >
        <div
          ref="historyContent"
          style=""
        >
          <div
            v-for="message in messageHistory"
            :key="message.id"
            :class="'message' + ' ' + (message.isBot ? '' : 'user-message')"
          >
            <el-avatar
              v-if="message.isBot"
              class="avatar"
              style="min-width: 40px"
              shape="circle"
              :size="40"
              :src="require('../assets/avatar.jpg')"
            />

            <div
              v-if="message.isBot"
              class="bubble bot-bubble"
            >
              <span>
                {{ message.content }}
              </span>
            </div>
            <div
              v-if="message.isBot"
              class="time bot-time"
            >
              <span class="time-upper" />
              <span>
                {{ parseTime(message.time) }}
              </span>
            </div>

            <div
              v-if="!message.isBot"
              class="bubble user-bubble"
            >
              <span>
                {{ message.content }}
              </span>
            </div>
            <div
              v-if="!message.isBot"
              class="time user-time"
            >
              <span class="time-upper" />
              <span>
                {{ parseTime(message.time) }}
              </span>
            </div>
          </div>
        </div>
      </el-scrollbar>
    </div>
    <div class="control-bar">
      <el-input
        :disabled="disabled"
        v-model="message"
        @keyup.enter="sendMessage()"
      />
      <el-button
        @click="sendMessage()"
        :disabled="disabled"
        class="send-btn"
        circle
        >发送
      </el-button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { MessagePacket } from '@/types/ChatTypes';

export default defineComponent({
  name: 'ChatBox',
  data() {
    return {
      disabled: false,
      message: '',
      latestMessageID: 0,
      messageHistory: [] as MessagePacket[],
    };
  },
  // setup() {
  //
  // },
  props: {
    botName: { type: String, default: '' },
    serviceMode: { type: String, default: '' },
    schema: { type: String, default: '' },
    error: { type: String, default: '' },
  },
  mounted() {
    let resp;
    try {
    } catch (error) {
      console.error(error);
    }
  },
  methods: {
    fillZero(num: number): string | number {
      return (num < 10 ? '0' : '') + num;
    },
    parseTime(time: number) {
      const date: Date = new Date(time);
      let hour: number = date.getHours();
      let minute: number = date.getMinutes();
      return (
        this.fillZero(hour) +
        ':' +
        this.fillZero(minute) +
        ' ' +
        (hour >= 12 ? 'PM' : 'AM')
      );
    },
    // to do
    goBack(): void {
      const parent = this.$parent;
      // parent.stage = 0;
    },
    sendMessage() {
      if (this.message === '') {
        return;
      }
      this.messageHistory.push({
        id: this.latestMessageID++,
        content: this.message,
        time: Date.now(),
        isBot: false,
      });
      this.message = '';
      this.$nextTick(() => {
        const historyContent = this.$refs.historyContent as HTMLElement;
        historyContent.scrollTop = historyContent.scrollHeight;
      });
    },
  },
});
</script>

<style scoped></style>
