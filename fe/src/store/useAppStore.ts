import { create } from 'zustand';
import { Department } from '@/api/model';

type AppStore = {
  activeDepartment: Department | null;
  setActiveDepartment: (department: Department | null) => void;
};

export const useAppStore = create<AppStore>((set) => ({
  activeDepartment: null,
  setActiveDepartment: (department: Department | null) =>
    set(() => ({
      activeDepartment: department,
    })),
}));
