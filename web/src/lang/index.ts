import {createI18n} from 'vue-i18n'
import enCustomLocale from './en'
import zhCustomLocale from './zh'
import elementEnLocale from 'element-plus/es/locale/lang/en'
import elementZhLocale from 'element-plus/es/locale/lang/zh-cn'

const jointMessages = {
    en: {
        ...enCustomLocale,
        ...elementEnLocale
    },
    zh: {
        ...zhCustomLocale,
        ...elementZhLocale
    }
}

export function getUserLanguage(): string {

    return 'zh';
}

const i18n = createI18n({
    locale: navigator.language,
    messages: jointMessages
})

export default i18n;
