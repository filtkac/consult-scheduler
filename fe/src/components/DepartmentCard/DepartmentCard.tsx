import { MouseEvent } from 'react';
import { Trans, useLingui } from '@lingui/react/macro';
import { IconEdit, IconTrash } from '@tabler/icons-react';
import { useMutation } from '@tanstack/react-query';
import { ActionIcon, Card, Group, Text, Title } from '@mantine/core';
import { modals } from '@mantine/modals';
import { deleteDepartmentMutationOptions } from '@/api/departments';
import { Department } from '@/api/model';
import { useAppStore } from '@/store/useAppStore';
import classes from './DepartmentCard.module.css';

interface DepartmentCardProps {
  department: Department;
}

export function DepartmentCard({ department }: DepartmentCardProps) {
  const { t } = useLingui();
  const { mutate: deleteDepartment, isPending: isDeletePending } = useMutation(
    deleteDepartmentMutationOptions
  );
  const { activeDepartment, setActiveDepartment } = useAppStore();

  const { name: departmentName } = department;

  const handleChangeActiveDepartment = (department: Department) => {
    setActiveDepartment(department);
  };

  const handleDepartmentEdit = (event: MouseEvent) => {
    event.stopPropagation();
    modals.openContextModal({
      modal: 'department',
      title: t`Edit Department`,
      innerProps: { department },
    });
  };

  const handleDepartmentDelete = (event: MouseEvent) => {
    event.stopPropagation();
    modals.openConfirmModal({
      title: t`Delete Department`,
      centered: true,
      children: (
        <Text>
          <Trans>Are you sure you want to delete the department "{departmentName}"?</Trans>
        </Text>
      ),
      labels: {
        confirm: t`Delete`,
        cancel: t`Cancel`,
      },
      confirmProps: { color: 'red', loading: isDeletePending },
      onConfirm: () => {
        if (activeDepartment?.id === department.id) {
          setActiveDepartment(null);
        }
        if (department.id) {
          deleteDepartment(department.id);
        }
      },
    });
  };

  return (
    <Card
      className={activeDepartment?.id === department.id ? classes.selectedCard : classes.card}
      key={department.id}
      withBorder
      shadow="sm"
      onClick={() => handleChangeActiveDepartment(department)}
    >
      <Card.Section withBorder p="md">
        <Title order={4} mb="xs">
          {department.name}
        </Title>
        <Text size="sm" c="dimmed" style={{ whiteSpace: 'pre-line' }}>
          {department.note}
        </Text>
      </Card.Section>
      <Card.Section>
        <Group justify="flex-end" p="sm" gap="4">
          <ActionIcon variant="subtle" onClick={handleDepartmentEdit}>
            <IconEdit />
          </ActionIcon>
          <ActionIcon variant="subtle" color="red" onClick={handleDepartmentDelete}>
            <IconTrash />
          </ActionIcon>
        </Group>
      </Card.Section>
    </Card>
  );
}
