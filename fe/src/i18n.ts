import { i18n } from '@lingui/core';

export const allLocales = ['cs', 'en'] as const;
type AvailableLocale = (typeof allLocales)[number];

const localeStorageKey = 'consult-scheduler-locale';
const defaultLocale: AvailableLocale = 'cs';

export function isLocale(str: string): str is AvailableLocale {
  return allLocales.includes(str as AvailableLocale);
}

export function getLocale(): AvailableLocale {
  const storedLocale = localStorage.getItem(localeStorageKey);
  if (storedLocale && isLocale(storedLocale)) {
    return storedLocale as AvailableLocale;
  }
  return defaultLocale;
}

export async function setLocale(locale: AvailableLocale) {
  if (isLocale(locale)) {
    localStorage.setItem(localeStorageKey, locale);
    await loadCatalog(locale);
  } else {
    throw Error(`Locale ${locale} not found`);
  }
}

export async function loadCatalog(locale: AvailableLocale) {
  const { messages } = await import(`./locales/${locale}/messages.po`);
  i18n.loadAndActivate({ locale, messages });
}
