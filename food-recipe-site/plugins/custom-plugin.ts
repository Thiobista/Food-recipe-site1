export default defineNuxtPlugin(() => {
    const runtimeConfig = useRuntimeConfig();
    console.log(runtimeConfig.public.apiUrl);
});
