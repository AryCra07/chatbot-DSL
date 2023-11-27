<template>
  <div class="login">
    <div class="box">
      <el-input
        class="username"
        type="text"
        v-model="username"
        :placeholder="this.$t('user_name')"
      />
      <el-input
        class="password"
        type="password"
        v-model="pwd"
        :placeholder="this.$t('pass_word')"
        @keyup.enter="submit"
        show-password
      />
      <el-button
        class="submit-btn"
        style="width: 300px"
        @click="submit"
        >{{ this.$t('login') }}
      </el-button>
    </div>
  </div>
</template>

<script lang="ts">
import { Login } from '@/api/login-register';
import { ref } from 'vue';
import { useI18n } from 'vue-i18n';
import { settingStore } from '@/store';
import { Md5 } from 'ts-md5';
import router from '@/router';
import { ElMessageBox } from 'element-plus';

export default {
  setup() {
    const username = ref('');
    const pwd = ref('');
    const t = useI18n()['t'];
    const store = settingStore();

    async function submit() {
      if (username.value.length < 2 || username.value.length > 20) {
        await ElMessageBox.alert(t('username_text'));
        return;
      } else if (pwd.value.length < 6) {
        await ElMessageBox.alert(t('password_text'));
        return;
      }

      let password_hash = Md5.hashStr(pwd.value + 'salt-ji8*y89dQuiGYG');

      let resp = await Login({
        data: { name: username.value, password: password_hash },
      });
      if (resp.code === 0) {
        store.$state.name = username.value;
        store.$state.token = resp.data.token;
        console.log(store);
        console.log('LR:' + store.$state.token);
        await router.push({ path: '/chat' });
      } else {
        console.log(resp);
        await ElMessageBox.alert(t('login_fail'));
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
  height: 100vh;
  overflow: hidden;
  background-color: white;
}

@media screen and (min-width: 961px) {
  .box {
    //background: linear-gradient(to bottom left, #baccd9, #11659a);
    background-color: #c7d2d4;
    position: relative;
    display: flex;
    flex-flow: column;
    width: 500px;
    height: 800px;
    border-radius: 5px;
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
    background: linear-gradient(#1781b5, #cdd1d3);
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
