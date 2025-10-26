import { DepartmentModal } from '@/components/DepartmentModal/DepartmentModal';

export const modals = {
  department: DepartmentModal,
};

declare module '@mantine/modals' {
  export interface MantineModalsOverride {
    modals: typeof modals;
  }
}
