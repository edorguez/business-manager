import { Icon } from '@iconify/react';
import { SideNavItem } from '../types';

export const SIDENAV_ITEMS: SideNavItem[] = [
  {
    title: 'Home',
    path: '/management/home',
    icon: <Icon icon="lucide:home" width="20" height="20" />,
  },
  {
    title: 'Clientes',
    path: '/management/customers',
    icon: <Icon icon="material-symbols:person" width="20" height="20" />,
  },
  {
    title: 'Inventario',
    path: '/management/products',
    icon: <Icon icon="fluent-mdl2:product" width="20" height="20" />,
    submenu: true,
    subMenuItems: [
      { title: 'Productos', path: '/management/products' },
      { title: 'Crear Producto', path: '/management/products/create' }
    ]
  },
  {
    title: 'Empresa',
    path: '/management/company',
    icon: <Icon icon="fluent:building-24-regular" width="20" height="20" />,
    submenu: true,
    subMenuItems: [
      { title: 'Métodos de Pago', path: '/management/payments' }
    ]
  },
  {
    title: 'WhatsApp',
    path: '/management/whatsapp',
    icon: <Icon icon="ic:baseline-whatsapp" width="20" height="20" />,
  },
  {
    title: 'Projects',
    path: '/management/projects',
    icon: <Icon icon="lucide:folder" width="20" height="20" />,
    submenu: true,
    subMenuItems: [
      { title: 'All', path: '/projects' },
      { title: 'Web Design', path: '/projects/web-design' },
      { title: 'Graphic Design', path: '/projects/graphic-design' },
    ],
  },
  {
    title: 'Settings',
    path: '/management/settings',
    icon: <Icon icon="lucide:settings" width="20" height="20" />,
    submenu: true,
    subMenuItems: [
      { title: 'Account', path: '/settings/account' },
      { title: 'Privacy', path: '/settings/privacy' },
    ],
  },
  {
    title: 'Help',
    path: '/help',
    icon: <Icon icon="lucide:help-circle" width="20" height="20" />,
  },
];
