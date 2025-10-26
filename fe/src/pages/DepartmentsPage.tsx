import { Trans, useLingui } from '@lingui/react/macro';
import { IconPlus } from '@tabler/icons-react';
import { useQuery } from '@tanstack/react-query';
import { Box, Button, LoadingOverlay, SimpleGrid } from '@mantine/core';
import { modals } from '@mantine/modals';
import { getDepartmentsQueryOptions } from '@/api/departments';
import { DepartmentCard } from '@/components/DepartmentCard/DepartmentCard';
import { appGradient } from '@/theme';

export function DepartmentsPage() {
  const { t } = useLingui();
  const { data: departments, isLoading } = useQuery(getDepartmentsQueryOptions);

  return (
    <Box pos="relative">
      <LoadingOverlay visible={isLoading} />
      <Button
        fullWidth
        variant="gradient"
        gradient={appGradient}
        leftSection={<IconPlus />}
        mb="sm"
        onClick={() => {
            modals.openContextModal({
                modal: 'department',
                title: t`Create Department`,
                innerProps: {department: undefined}
            })
        }}
      >
        <Trans>Create a new department</Trans>
      </Button>
      <SimpleGrid cols={{ base: 1, sm: 2, md: 3, lg: 4 }}>
        {departments?.map((dep) => (
          <DepartmentCard key={dep.id} department={dep} />
        ))}
      </SimpleGrid>
    </Box>
  );
}
