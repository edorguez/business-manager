import { Icon } from "@iconify/react";
import { SideNavItem } from "../types";

export class USER_ROLE_ID {
  public static SUPER_ADMIN = 1;
  public static ADMIN = 2;
  public static REGULAR = 3;
}

export class PLAN_ID {
  public static BASIC = 1;
  public static PRO = 2;
}

export class PRODUCT {
  public static MAX_BASIC_PLAN_ITEMS = 10;
}

export class PAYMENT {
  public static MAX_BASIC_PLAN_ITEMS = 5;
}

export class PASSWORD {
  public static MIN_PASSWORD_LEGTH = 6;
}

export const SIDENAV_ITEMS: SideNavItem[] = [
  {
    title: "Home",
    path: "/management/home",
    icon: <Icon icon="lucide:home" width="20" height="20" />,
    roleIds: [
      USER_ROLE_ID.SUPER_ADMIN,
      USER_ROLE_ID.ADMIN,
      USER_ROLE_ID.REGULAR,
    ],
  },
  {
    title: "Clientes",
    path: "/management/customers",
    icon: <Icon icon="material-symbols:person" width="20" height="20" />,
    roleIds: [
      USER_ROLE_ID.SUPER_ADMIN,
      USER_ROLE_ID.ADMIN,
      USER_ROLE_ID.REGULAR,
    ],
  },
  {
    title: "Inventario",
    path: "/management/products",
    icon: <Icon icon="fluent-mdl2:product" width="20" height="20" />,
    submenu: true,
    subMenuItems: [
      {
        title: "Productos",
        path: "/management/products",
        roleIds: [
          USER_ROLE_ID.SUPER_ADMIN,
          USER_ROLE_ID.ADMIN,
          USER_ROLE_ID.REGULAR,
        ],
      },
      {
        title: "Crear Producto",
        path: "/management/products/create",
        roleIds: [
          USER_ROLE_ID.SUPER_ADMIN,
          USER_ROLE_ID.ADMIN,
          USER_ROLE_ID.REGULAR,
        ],
      },
    ],
    roleIds: [
      USER_ROLE_ID.SUPER_ADMIN,
      USER_ROLE_ID.ADMIN,
      USER_ROLE_ID.REGULAR,
    ],
  },
  {
    title: "Empresa",
    path: "/management/company",
    icon: <Icon icon="fluent:building-24-regular" width="20" height="20" />,
    submenu: true,
    subMenuItems: [
      {
        title: "Métodos de Pago",
        path: "/management/payments",
        roleIds: [
          USER_ROLE_ID.SUPER_ADMIN,
          USER_ROLE_ID.ADMIN,
          USER_ROLE_ID.REGULAR,
        ],
      },
    ],
    roleIds: [
      USER_ROLE_ID.SUPER_ADMIN,
      USER_ROLE_ID.ADMIN,
      USER_ROLE_ID.REGULAR,
    ],
  },
  // {
  //   title: "WhatsApp",
  //   path: "/management/whatsapp",
  //   icon: <Icon icon="ic:baseline-whatsapp" width="20" height="20" />,
  //   roleIds: [
  //     USER_ROLE_ID.SUPER_ADMIN,
  //     USER_ROLE_ID.ADMIN,
  //     USER_ROLE_ID.REGULAR,
  //   ],
  // },
  {
    title: "Usuarios",
    path: "/management/users",
    icon: <Icon icon="fluent:people-team-20-filled" />,
    roleIds: [USER_ROLE_ID.SUPER_ADMIN, USER_ROLE_ID.ADMIN],
  },
  // {
  //   title: "Help",
  //   path: "/help",
  //   icon: <Icon icon="lucide:help-circle" width="20" height="20" />,
  //   roleIds: [
  //     USER_ROLE_ID.SUPER_ADMIN,
  //     USER_ROLE_ID.ADMIN,
  //     USER_ROLE_ID.REGULAR,
  //   ],
  // },
];
