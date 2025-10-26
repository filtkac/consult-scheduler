import { UseMutationOptions, UseQueryOptions } from '@tanstack/react-query';
import { AxiosError } from 'axios';
import { Department } from '@/api/model';
import { api, queryClient } from './client';

export const getDepartmentsQueryOptions: UseQueryOptions<Department[]> = {
  queryKey: ['departments'],
  queryFn: async () => {
    const response = await api.get<Department[]>('/departments');
    return response.data;
  },
  staleTime: 5 * 60 * 1000, // 5 minutes
};

export const createDepartmentMutationOptions: UseMutationOptions<
  Department,
  AxiosError,
  Department
> = {
  mutationFn: async (department: Department) => {
    const response = await api.post<Department>('/departments', department);
    return response.data;
  },
  onSuccess: async () => {
    await queryClient.invalidateQueries({ queryKey: ['departments'] });
  },
};

export const updateDepartmentMutationOptions: UseMutationOptions<
  Department,
  AxiosError,
  Department
> = {
  mutationFn: async (department: Department) => {
    const response = await api.put<Department>(`/departments/${department.id}`, department);
    return response.data;
  },
  onSuccess: async () => {
    await queryClient.invalidateQueries({ queryKey: ['departments'] });
  },
};

export const deleteDepartmentMutationOptions: UseMutationOptions<void, AxiosError, number> = {
  mutationFn: async (departmentId: number) => {
    await api.delete(`/departments/${departmentId}`);
  },
  onSuccess: async () => {
    await queryClient.invalidateQueries({ queryKey: ['departments'] });
  },
};
