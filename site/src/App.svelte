<script>
  import Router from "svelte-spa-router";
  import { user_routes } from "./lib/modules/user/views/routes";
  import { main_routes } from "./lib/modules/main/views/routes";
  import { question_routes } from "./lib/modules/question/views/routes";
</script>

<Router
  routes={{
    ...user_routes,
    ...question_routes,
    ...main_routes,
  }}
/>
