import { Trans, useLingui } from '@lingui/react/macro';
import {
  IconBuilding,
  IconCalendarTime,
  IconClipboardList,
  IconClock,
  IconOld,
} from '@tabler/icons-react';
import { useQuery } from '@tanstack/react-query';
import { NavLink as ReactRouterNavLink } from 'react-router-dom';
import {
  AppShell,
  ComboboxItem,
  Flex,
  getGradient,
  Loader,
  NavLink,
  Select,
  Title,
  useMantineTheme,
} from '@mantine/core';
import { getDepartmentsQueryOptions } from '@/api/departments';
import { useAppStore } from '@/store/useAppStore';
import { appGradient } from '@/theme';

export function Navbar() {
  const theme = useMantineTheme();
  const { t } = useLingui();
  const { data: departments, isLoading: departmentsLoading } = useQuery(getDepartmentsQueryOptions);
  const { activeDepartment, setActiveDepartment } = useAppStore();

  const getDepartmentsForSelect = (): ComboboxItem[] => {
    return (
      departments
        ?.map((dep) => {
          if (dep.id) {
            return {
              value: dep.id.toString(),
              label: dep.name,
            };
          }
          return null;
        })
        .filter((dep) => !!dep) ?? []
    );
  };

  const handleChangeActiveDepartment = (value: string | null) => {
    const activeDepartmentToSet = departments?.find((dep) => dep.id?.toString() === value);
    if (activeDepartmentToSet) {
      setActiveDepartment(activeDepartmentToSet);
      return;
    }
    setActiveDepartment(null);
  };

  return (
    <>
      <Flex align="center" justify="center" gap="sm" p="md" bg={getGradient(appGradient, theme)}>
        <Flex flex="none">
          <IconClock size={32} stroke={2.3} color="white" />
        </Flex>
        <Title order={1} size="h3" c="white">
          <Trans>Consult scheduling</Trans>
        </Title>
      </Flex>
      <NavLink
        component={ReactRouterNavLink}
        to="/consults"
        label={t`Consults`}
        leftSection={<IconCalendarTime />}
      />
      <NavLink
        component={ReactRouterNavLink}
        to="/day-templates"
        label={t`Templates`}
        leftSection={<IconClipboardList />}
      />
      <NavLink
        component={ReactRouterNavLink}
        to="/departments"
        label={t`Departments`}
        leftSection={<IconBuilding />}
      />
      <NavLink
        component={ReactRouterNavLink}
        to="/patients"
        label={t`Patients`}
        leftSection={<IconOld />}
        mb="md"
      />
      <AppShell.Section>
        <Select
          data={getDepartmentsForSelect()}
          value={activeDepartment?.id?.toString()}
          onChange={handleChangeActiveDepartment}
          label={t`Active department`}
          placeholder={t`Pick one`}
          searchable
          disabled={departmentsLoading}
          rightSection={departmentsLoading ? <Loader size="xs" /> : null}
          m="xs"
        />
      </AppShell.Section>
    </>
  );
}
