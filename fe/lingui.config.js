import { defineConfig } from "@lingui/cli";
import { allLocales } from './src/i18n.ts';

export default defineConfig({
  sourceLocale: "en",
  locales: allLocales,
  catalogs: [
    {
      path: "<rootDir>/src/locales/{locale}/messages",
      include: ["src"],
    },
  ],
});
