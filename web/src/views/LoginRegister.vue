<template>
  <div class="login">
    <div class="box">
      <el-input
        class="username"
        type="text"
        v-model="username"
        :placeholder="$t('message.username')"
      />
      <el-input
        class="password"
        type="password"
        v-model="pwd"
        :placeholder="$t('message.pwd')"
        @keyup.enter="submit"
        show-password
      />
      <el-button
        class="submit-btn"
        style="width: 300px"
        @click="submit"
        >{{ $t('message.sign') }}</el-button
      >
    </div>
  </div>
</template>

<script>
import { Login } from '@/api/login-register';
import { ref } from 'vue';
import { useI18n } from 'vue-i18n';
import { openToast } from 'toast-ts';
import { settingStore } from '@/store';
import { Md5 } from 'ts-md5';
import router from '@/router';

export default {
  setup() {
    const username = ref('');
    const pwd = ref('');
    const t = useI18n()['t'];
    const store = settingStore();

    async function submit() {
      if (username.value.length < 4 || username.value.length > 32) {
        openToast(t('message.usrtext'));
        return;
      } else if (/^[a-z0-9A-Z_]+$/.test(username.value) === false) {
        openToast(t('message.illegal_username'));
        return;
      } else if (pwd.value.length < 8) {
        openToast(t('message.pwdtext'));
        return;
      }

      let pwdhash = Md5.hashStr(pwd.value + 'salt-9aSO(UIf89!(*@&12');
      let resp = await Login({ username: username.value, pwd: pwdhash });
      if (resp === null) {
        console.log(1);
      }
      if (pwdhash !== '1') {
        store.$state.name = username.value;
        router.push('/chat');
      } else {
        openToast(t('message.wrongpwd'));
      }
    }

    return {
      username,
      pwd,
      submit,
    };
  },
};
</script>

<style lang="scss" scoped>
.login {
  display: flex;
  justify-content: center; /* 水平居中 */
  align-items: center; /* 垂直居中 */
  min-height: 100vh; /* 保持内容垂直居中 */
}

@media screen and (min-width: 961px) {
  .box {
    background: ivory;
    position: relative;
    display: flex;
    flex-flow: column;
    width: 500px;
    height: 800px;
    border-radius: 5px;
    //background-color: white;
    overflow: hidden;
    box-shadow:
      0 2px 4px rgba(0, 0, 0, 0.12),
      0 0 6px rgba(0, 0, 0, 0.04);
    align-items: center;
    padding-top: 240px;
    box-sizing: border-box;
  }
}

@media (max-width: 960px) {
  .box {
    background: #2c3e50;
    position: relative;
    display: flex;
    flex-flow: column;
    width: 100vw;
    height: 100vh;
    //background-color: white;
    overflow: hidden;
    box-shadow:
      0 2px 4px rgba(0, 0, 0, 0.12),
      0 0 6px rgba(0, 0, 0, 0.04);
    align-items: center;
    padding-top: 30vh;
    box-sizing: border-box;
  }
}

.username {
  margin-bottom: 10px;
  width: 300px;
}

.password {
  margin-bottom: 30px;
  width: 300px;
}

.submit-btn {
  width: 100px;
}
</style>
