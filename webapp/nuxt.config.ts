// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },

  app: {
  },

  ssr: false,
  modules: ["@nuxtjs/tailwindcss"]
})