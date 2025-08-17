import mantine from 'eslint-config-mantine';
import tseslint from 'typescript-eslint';
import pluginLingui from 'eslint-plugin-lingui';

export default tseslint.config(
  ...mantine,
    pluginLingui.configs['flat/recommended'],
  { ignores: ['**/*.{mjs,cjs,js,d.ts,d.mts}', './src/locales'] },
  {
    files: ['**/*.story.tsx'],
    rules: { 'no-console': 'off' },
  }
);
