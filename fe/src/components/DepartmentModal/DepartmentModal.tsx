import { useLingui } from '@lingui/react/macro';
import { useMutation } from '@tanstack/react-query';
import { Button, Group, Textarea, TextInput } from '@mantine/core';
import { useForm } from '@mantine/form';
import { ContextModalProps } from '@mantine/modals';
import {
  createDepartmentMutationOptions,
  updateDepartmentMutationOptions,
} from '@/api/departments';
import { Department } from '@/api/model';
import { appGradient } from '@/theme';


interface DepartmentModalProps {
  department?: Department;
}

export const DepartmentModal = ({ context, id, innerProps }: ContextModalProps<DepartmentModalProps>) => {
  const { t } = useLingui();

  const { mutate: create, isPending: isCreatePending } = useMutation(createDepartmentMutationOptions);
  const { mutate: update, isPending: isUpdatePending } = useMutation(updateDepartmentMutationOptions);
  const isPending = isCreatePending || isUpdatePending;

  const department = innerProps.department;

  const form = useForm({
    mode: 'uncontrolled',
    initialValues: {
      name: department ? department.name : '',
      note: department?.note ? department.note : '',
    },
  });

  const handleSubmit = (values: typeof form.values) => {
    if (department) {
      update({ id: department.id, ...values})
    } else {
      create(values);
    }
    context.closeModal(id);
  };

  return (
    <form onSubmit={form.onSubmit(handleSubmit)}>
      <TextInput
        label={t`Name`}
        key={form.key('name')}
        data-autofocus
        {...form.getInputProps('name')}
      />
      <Textarea
        label={t`Note`}
        minRows={3}
        autosize
        mt="sm"
        key={form.key('note')}
        {...form.getInputProps('note')}
      />
      <Group justify="end">
        <Button type="submit" mt="md" variant="gradient" gradient={appGradient} loading={isPending}>
          {department ? t`Save Changes` : t`Create Department`}
        </Button>
      </Group>
    </form>
  );
};
