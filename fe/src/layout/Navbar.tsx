import { Trans, useLingui } from '@lingui/react/macro';
import {
  IconBuilding,
  IconCalendarTime,
  IconClipboardList,
  IconClock,
  IconOld,
} from '@tabler/icons-react';
import { NavLink as ReactRouterNavLink } from 'react-router-dom';
import { Flex, NavLink, Title } from '@mantine/core';

export function Navbar() {
  const { t } = useLingui();

  return (
    <>
      <Flex align="center" justify="center" gap="sm" p="md" bg="cyan.9">
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
      />
    </>
  );
}
