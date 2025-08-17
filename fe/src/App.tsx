


import '@mantine/core/styles.css';

import { i18n } from '@lingui/core';
import { I18nProvider } from '@lingui/react';
import { localStorageColorSchemeManager, MantineProvider } from '@mantine/core';
import { getLocale, loadCatalog } from '@/i18n';
import { Router } from './Router';
import { theme } from './theme';


await loadCatalog(getLocale());

export default function App() {
  return (
    <MantineProvider
      theme={theme}
      colorSchemeManager={localStorageColorSchemeManager()}
      defaultColorScheme="light"
    >
      <I18nProvider i18n={i18n}>
        <Router />
      </I18nProvider>
    </MantineProvider>
  );
}
