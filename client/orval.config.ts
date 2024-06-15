import { defineConfig } from "orval";

export default defineConfig({
  backend: {
    input: {
      target: "../docs/openapi.yml",
    },
    output: {
      target: "./src/generated/api/openapi.ts",
      baseUrl: "/api",
      clean: true,
      client: "vue-query",
    },
  },
});
