import { NextRequest, NextResponse } from "next/server";
import isValidLogin from "./app/actions/isValidLogin";
import { SideNavItem } from "./app/types";
import { SIDENAV_ITEMS } from "./app/constants";
import { CurrentUser } from "./app/types/auth";
import { validateUserInRoles } from "./app/utils/Utils";
import getCurrentUserServer from "./app/actions/getCurrentUserServer";

export function middleware(request: NextRequest) {
  if (request.nextUrl.pathname.startsWith("/management")) {
    return handleManagementRoute(request);
  }

  return handleSubdomains(request);
}

function handleSubdomains(request: NextRequest) {
  const url = request.nextUrl;
  const pathname = url.pathname;

  // Get hostname (e.g., 'mike.com', 'test.mike.com')
  const hostname = request.headers.get("host");

  let currentHost;
  if (process.env.ENVIRONMENT === "production") {
    // Production logic remains the same
    const baseDomain = process.env.BASE_DOMAIN;
    currentHost = hostname?.replace(`.${baseDomain}`, "");
  } else {
    // Updated development logic
    currentHost = hostname?.split(":")[0].replace(".localhost", "");
  }

  // If there's no currentHost, likely accessing the root domain, handle accordingly
  if (!currentHost) {
    // Continue to the next middleware or serve the root content
    return NextResponse.next();
  }

  // Fetch tenant-specific data based on the hostname
  const response: any = [{ hola: true, chao: 1 }]; //await readSiteDomain(currentHost);

  // Handle the case where no domain data is found
  if (!response || !response.length) {
    // Continue to the next middleware or serve the root content
    return NextResponse.next();
  }

  const site_id = response[0]?.site_id;
  const tenantSubdomain = response[0]?.site_subdomain;
  const mainDomain = response[0]?.site_custom_domain;


  console.log('values')
  console.log(site_id)
  console.log(tenantSubdomain)
  console.log(mainDomain)

  // Determine which domain to use for rewriting
  const rewriteDomain = tenantSubdomain || mainDomain;

  console.log("Hostname:", hostname);
  console.log("Current Host:", currentHost);
  console.log("Rewrite Domain:", rewriteDomain);

  if (rewriteDomain) {
    // Rewrite the URL to the tenant-specific path, using the site_id
    return NextResponse.rewrite(new URL(`/${site_id}${pathname}`, request.url));
  }

  // If no rewrite domain is found, continue to the next middleware
  return NextResponse.next();
}

function handleManagementRoute(request: NextRequest) {
  // Validate if user is logged
  const isUserLogged: boolean = isValidLogin();
  if (!isUserLogged)
    return NextResponse.redirect(new URL("/login", request.url));

  // Validate if user role has permission to path
  const sideNavItem: SideNavItem | undefined = SIDENAV_ITEMS.find((x) =>
    request.nextUrl.pathname.includes(x.path)
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

// export const config = {
//   matcher: ["/management/:path*"], // Routes that should use this middleware
// };
