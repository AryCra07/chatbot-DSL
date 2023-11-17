import { createI18n } from 'vue-i18n';
import ehLocale from './en';
import zhLocale from './zh';
import elementEnLocale from 'element-plus/es/locale/lang/en';
import elementZhLocale from 'element-plus/es/locale/lang/zh-cn';

const jointMessages = {
  en: {
    ...ehLocale,
    ...elementEnLocale,
  },
  'zh-CN': {
    ...zhLocale,
    ...elementZhLocale,
  },
};

export function getUserLanguage(): string {
  return 'zh';
}

const i18n = createI18n({
  allowComposition: true,
  locale: navigator.language,
  messages: jointMessages,
});

// const i18n = new VueI18n({
//   // set locale
//   // options: en | zh | es
//   locale: navigator.language,
//   // set locale messages
//   jointMessages,
// });

export default i18n;
