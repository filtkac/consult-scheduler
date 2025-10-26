import '@mantine/core/styles.css';
import '@mantine/notifications/styles.css';

import { i18n } from '@lingui/core';
import { I18nProvider } from '@lingui/react';
import { QueryClientProvider } from '@tanstack/react-query';
import { localStorageColorSchemeManager, MantineProvider } from '@mantine/core';
import { ModalsProvider } from '@mantine/modals';
import { queryClient } from '@/api/client';
import { getLocale, loadCatalog } from '@/i18n';
import { modals } from '@/modals';
import { Router } from './Router';
import { theme } from './theme';

await loadCatalog(getLocale());

export default function App() {
  return (
    <I18nProvider i18n={i18n}>
      <QueryClientProvider client={queryClient}>
        <MantineProvider
          theme={theme}
          colorSchemeManager={localStorageColorSchemeManager()}
          defaultColorScheme="light"
        >
          <ModalsProvider modals={modals}>
            <Router />
          </ModalsProvider>
        </MantineProvider>
      </QueryClientProvider>
    </I18nProvider>
  );
}
