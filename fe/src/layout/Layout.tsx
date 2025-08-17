import { Outlet } from 'react-router-dom';
import { AppShell } from '@mantine/core';
import { Header } from '@/layout/Header';
import { Navbar } from '@/layout/Navbar';

export function Layout() {
  return (
    <AppShell
      padding="md"
      navbar={{ width: 190, breakpoint: 'sm' }}
      header={{ height: 60 }}
      layout="alt"
    >
      <AppShell.Header>
        <Header />
      </AppShell.Header>
      <AppShell.Navbar>
        <Navbar />
      </AppShell.Navbar>
      <AppShell.Main>
        <Outlet />
      </AppShell.Main>
    </AppShell>
  );
}
