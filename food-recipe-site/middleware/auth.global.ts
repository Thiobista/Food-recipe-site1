export default defineNuxtRouteMiddleware(() => {
    const config = useRuntimeConfig();
    console.log("Middleware config:", config.public.apiUrl);
  });
  