import { createBrowserRouter, redirect, RouterProvider } from 'react-router-dom';
import { Layout } from '@/layout/Layout';
import { ConsultsPage } from '@/pages/ConsultsPage';
import { DayTemplatesPage } from '@/pages/DayTemplatesPage';
import { DepartmentsPage } from '@/pages/DepartmentsPage';
import { PatientsPage } from '@/pages/PatientsPage';

const router = createBrowserRouter([
  {
    path: '/',
    element: <Layout />,
    children: [
      {
        index: true,
        loader: () => redirect('/consults')
      },
      {
        path: 'consults',
        element: <ConsultsPage />
      },
      {
        path: 'day-templates',
        element: <DayTemplatesPage />
      },
      {
        path: 'departments',
        element: <DepartmentsPage />
      },
      {
        path: 'patients',
        element: <PatientsPage />
      }
    ],
  },
]);

export function Router() {
  return <RouterProvider router={router} />;
}
