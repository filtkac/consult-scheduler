import { useLingui } from '@lingui/react/macro';
import { IconMoon, IconSun, IconUserSearch } from '@tabler/icons-react';
import {
  ActionIcon,
  Autocomplete,
  Group,
  localStorageColorSchemeManager,
  MantineColorScheme,
  Select,
  useComputedColorScheme,
  useMantineColorScheme,
} from '@mantine/core';
import { getLocale, isLocale, setLocale } from '@/i18n';

export function Header() {
  const colorSchemeManager = localStorageColorSchemeManager();
  const { setColorScheme } = useMantineColorScheme();
  const colorScheme = useComputedColorScheme('light');

  const { t } = useLingui();
  const locale = getLocale();

  const handleToggleColorScheme = () => {
    const newColorScheme: MantineColorScheme = colorScheme === 'light' ? 'dark' : 'light';
    colorSchemeManager.set(newColorScheme);
    setColorScheme(newColorScheme);
  };

  const handleLocaleChange = async (value: string | null) => {
    if (value && isLocale(value)) {
      await setLocale(value);
    }
  };

  const Icon = colorScheme === 'light' ? IconMoon : IconSun;

  return (
    <Group align="center" justify="space-between">
      <Group align="center" justify="flex-start" p="sm">
        <Autocomplete rightSection={<IconUserSearch />} placeholder={t`Look up patient`} />
      </Group>
      <Group align="center" justify="flex-end" p="sm">
        <Select
          data={[
            { value: 'en', label: '🇺🇸' },
            { value: 'cs', label: '🇨🇿' },
          ]}
          value={locale}
          onChange={handleLocaleChange}
          size="sm"
          withCheckIcon={false}
          w="65"
          comboboxProps={{ width: 50, position: 'bottom-start' }}
        />
        <ActionIcon variant="default" size="lg" onClick={handleToggleColorScheme}>
          <Icon size={20} />
        </ActionIcon>
      </Group>
    </Group>
  );
}
