import { NextRequest, NextResponse } from "next/server";
import isValidLogin from "./app/actions/isValidLogin";
import { SideNavItem } from "./app/types";
import { SIDENAV_ITEMS } from "./app/constants";
import { CurrentUser } from "./app/types/auth";
import { validateUserInRoles } from "./app/utils/Utils";
import getCurrentUserServer from "./app/actions/getCurrentUserServer";

export function middleware(request: NextRequest) {
  
  // Validate if user is logged
  const isUserLogged: boolean = isValidLogin();
  if (!isUserLogged)
    return NextResponse.redirect(new URL("/login", request.url));

  // Validate if user role has permission to path
  const sideNavItem: SideNavItem | undefined = SIDENAV_ITEMS.find(
    (x) => request.nextUrl.pathname.includes(x.path)
  );
  if (sideNavItem) {
    const user: CurrentUser | null = getCurrentUserServer();
    if (user && !validateUserInRoles(sideNavItem.roleIds, user.roleId)) {
      return NextResponse.redirect(new URL("/management/home", request.url));
    }
  }

  // Allow access to path
  return NextResponse.next();
}

export const config = {
  matcher: ["/management/:path*"], // Routes that should use this middleware
};
