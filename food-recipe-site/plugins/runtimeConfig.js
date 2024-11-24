export default defineNuxtPlugin((nuxtApp) => {
    const config = useRuntimeConfig();
    console.log("Public API URL:", config.public.apiUrl);
});
