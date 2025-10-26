import pluginQuery from '@tanstack/eslint-plugin-query';
import mantine from 'eslint-config-mantine';
import pluginLingui from 'eslint-plugin-lingui';
import tseslint from 'typescript-eslint';

export default tseslint.config(
  ...mantine,
  ...pluginQuery.configs['flat/recommended'],
  pluginLingui.configs['flat/recommended'],
  { ignores: ['**/*.{mjs,cjs,js,d.ts,d.mts}', './src/locales'] },
  {
    files: ['**/*.story.tsx'],
    rules: { 'no-console': 'off' },
  }
);
