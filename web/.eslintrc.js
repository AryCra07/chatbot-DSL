module.exports = {
  root: true,
  env: {
    node: true
  },
  'extends': [
    'plugin:vue/vue3-essential',
    'eslint:recommended',
    '@vue/typescript/recommended'
  ],
  parserOptions: {
    ecmaVersion: 2020
  },
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 1 : 0,
    'no-debugger': process.env.NODE_ENV === 'production' ? 2 : 0,
    '@typescript-eslint/explicit-module-boundary-types': 'off',
    'quotes': [1, 'single'],
  }
}
