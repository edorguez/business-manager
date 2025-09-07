export type SideNavItem = {
  title: string;
  path: string;
  icon?: JSX.Element;
  submenu?: boolean;
  subMenuItems?: SideNavItem[];
  roleIds: number[];
};

export type BreadcrumItem = {
  label: string;
  href: string;
}
