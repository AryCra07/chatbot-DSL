<template>
  <div class="chat-room">
    <div class="header">
      <el-avatar
        class="avatar"
        shape="square"
        :size="50"
        :src="require('../assets/img/logo.png')"
      />
      <div class="bot-name">
        <span class="title"> AryCra07 </span>
        <span class="subtitle">chat bot</span>
      </div>
      <el-button class="control">
        <span
          class="back"
          @click="goBack"
        >
          {{ $t('back') }}
        </span>
      </el-button>
    </div>

    <div style="height: 10px"></div>
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
              :src="require('../assets/img/logo.png')"
            />
            <el-avatar
              v-if="!message.isBot"
              class="avatar"
              style="min-width: 40px"
              shape="circle"
              :size="40"
              :src="require('../assets/img/avatar2.jpg')"
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
    <div style="height: 80px"></div>
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
import { fetchHello, fetchMessage, fetchTimer } from '@/api/chat';
import { settingStore } from '@/store';

export default defineComponent({
  name: 'ChatBox',
  data() {
    return {
      disabled: false,
      message: '',
      latestMessageID: 0,
      messageHistory: [] as any,
      timerId: 0 as any,
      store: settingStore(),
      reset: false,
      exit: false,
    };
  },
  props: {
    botName: { type: String, default: '' },
    serviceMode: { type: String, default: '' },
    schema: { type: String, default: '' },
    error: { type: String, default: '' },
  },
  async mounted() {
    // console.log(store.name)
    let fail = false;
    let resp;
    try {
      resp = await fetchHello();
      if (resp.code != 0) {
        fail = true;
      }
    } catch (error) {
      fail = true;
    }

    if (fail) {
      this.disabled = true;
      this.message = this.$t('service_error');
      return;
    }

    let msgList = resp.data.content;
    if (msgList !== null) {
      msgList.forEach((msg) => {
        this.loadMessage(true, msg, Date.now());
      });
    }
    this.startTimer(5000);
  },
  methods: {
    startTimer(ms: number = 2000) {
      let now_sec = 0;

      const timerHandler = async () => {
        if (this.exit) {
          return
        }
        if (this.reset) {
          now_sec = 0;
          this.reset = false
        }
        now_sec = now_sec + ms / 1000;
        console.log(now_sec);
        const resp = await fetchTimer({
          data: {
            last_time: Math.round(now_sec - ms / 1000),
            now_time: Math.round(now_sec),
          },
        });

        let msgList = resp.data.content;
        if (msgList !== null) {
          msgList.forEach((msg) => {
            this.loadMessage(true, msg, Date.now());
          });
        }

        if (resp.data.reset === true) {
          this.reset = true
        }

        if (resp.data.exit === true) {
          this.store.setToken('');
          await fetchTimer({
            data: {
              last_time: Math.round(now_sec - ms),
              now_time: Math.round(now_sec),
            },
          });
          // 设置标志为 false，不再继续执行定时器
          this.exit = true;
        }

        // 如果标志为 false，则继续执行定时器
        if (!this.exit) {
          setTimeout(timerHandler, ms);
        }
      };

      // 第一次启动定时器
      setTimeout(timerHandler, ms);
    },

    async loadMessage(bot: boolean, msg: string, time: number) {
      this.messageHistory.push({
        msg_id: this.latestMessageID++,
        isBot: bot,
        content: msg,
        time: time,
      });
      await this.$nextTick();
      let historyContent = this.$refs.historyContent as any;
      let historyScrollbar = this.$refs.historyScrollbar as any;
      historyScrollbar.setScrollTop(historyContent.clientHeight);
    },
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
      this.exit = true;
      this.store.setToken('');
      this.store.setUserName('');
      this.$router.push({ path: '/login' });
    },
    async sendMessage() {
      let msg: string = this.message;
      if (msg.length == 0) {
        return;
      }
      let time: number = new Date().getTime();
      this.message = '';
      await this.loadMessage(false, msg, time);

      try {
        let resp = await fetchMessage({
          data: { name: this.store.name, input: msg },
        });

        if (resp.code == 2) {
          msg = this.$t('expire');
          time = new Date().getTime();
          await this.loadMessage(true, msg, time);
        } else if (resp.code != 0) {
          msg = this.error;
          time = new Date().getTime();
          await this.loadMessage(true, msg, time);
        } else {
          this.reset = true;
          let data = resp.data.content;
          data.forEach((elem) => {
            this.loadMessage(true, elem, Date.now());
          });
        }
      } catch (error) {
        console.error(error);
        msg = this.error;
        time = new Date().getTime();
        await this.loadMessage(true, msg, time);
      }
    },
  },
});
</script>

<style lang="scss" scoped>
.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
  height: 80px;
  box-sizing: border-box;
  padding: 20px;
  box-shadow:
    0 2px 4px rgba(0, 0, 0, 0.2),
    0 0 6px rgba(0, 0, 0, 0.04);

  .avatar {
    flex: 0 0 auto;
  }

  .bot-name {
    display: flex;
    flex: 0 0 auto;
    margin-left: 20px;
    flex-flow: column;
    font-size: 20px;

    span {
      text-align: left;
    }

    .title {
    }

    .subtitle {
      font-size: 14px;
      color: #8f8f8f;
    }
  }

  .control {
    flex: 0 0 auto;
    text-align: right;
    margin-left: auto;

    .back {
      cursor: pointer;
    }
  }
}

.main {
  position: relative;
  flex: 1;
  width: 100%;
  overflow-y: auto;

  .chat {
    height: 100%;
    padding: 0 20px 0;

    .message {
      display: flex;
      padding: 0 0 20px;
      width: 100%;
      text-align: left;
    }

    .user-message {
      flex-flow: row-reverse;
    }

    .bubble {
      position: relative;
      width: auto;
      padding: 10px;
      background: #1491a8;
      -moz-border-radius: 10px;
      -webkit-border-radius: 10px;
      border-radius: 10px;
      color: white;
      font-family: 'JetBrains Mono', '黑体', 'sans-serif';

      span {
        white-space: pre-wrap;
        word-break: break-word;
      }
    }

    .bot-bubble::before {
      content: '';
      position: absolute;
      width: 0;
      height: 0;
      border-top: 13px solid transparent;
      border-right: 26px solid #1491a8;
      border-bottom: 13px solid transparent;
      margin: -3px 0 0 -25px;
    }

    .bot-bubble {
      margin-left: 25px;
    }

    .user-bubble::after {
      content: '';
      position: absolute;
      top: 0;
      right: 0;
      width: 0;
      height: 0;
      border-top: 13px solid transparent;
      border-left: 25px solid #f07c82;
      border-bottom: 13px solid transparent;
      margin: 6px -15px 0 0;
    }

    .user-bubble {
      margin-right: 25px;
      background-color: #f07c82;
    }

    .time {
      position: relative;
      display: flex;
      flex-flow: column;
      min-width: 70px;

      .time-upper {
        flex: 1;
      }

      span {
        color: #c7c7c7;
        font-size: 14px;
      }
    }

    .bot-time span {
      margin-left: 5px;
    }

    .user-time span {
      margin-right: 5px;
    }
  }
}

.control-bar {
  display: flex;
  flex-flow: row;
  justify-content: center;
  align-items: center;
  box-sizing: border-box;
  padding: 0 20px;
  width: 100%;
  height: 80px;
  border-top: 1px solid #c7c7c7;
  position: fixed;
  bottom: 0;
  background-color: white;

  .el-input :deep(.el-input__wrapper) {
    border-radius: 20px;
  }

  .send-btn {
    padding: 15px;
    margin-left: 20px;
  }
}
</style>
