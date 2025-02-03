import { NextRequest, NextResponse } from "next/server";
import isValidLogin from "./app/actions/isValidLogin";
import { SIDENAV_ITEMS } from "./app/constants";
import { validateUserInRoles } from "./app/utils/Utils";
import getCurrentUserServer from "./app/actions/getCurrentUserServer";

export function middleware(request: NextRequest) {
  const { pathname } = request.nextUrl;

  if (
    pathname.startsWith("/_next/static") ||
    pathname.startsWith("/_next/image") ||
    pathname.startsWith("/api") ||
    pathname === "/favicon.ico"
  ) {
    return NextResponse.next();
  }

  if (pathname.startsWith("/management")) {
    return handleManagementRoute(request);
  }

  return handleSubdomains(request);
}

function handleSubdomains(request: NextRequest) {
  console.log("----------------------------------");

  const url = request.nextUrl;
  let hostname = request.headers.get("host");

  if (!hostname) {
    return NextResponse.next(); // No host header, continue normally
  }

  // Cloudflare might send 'x-forwarded-host'
  const forwardedHost = request.headers.get("x-forwarded-host");
  if (forwardedHost) {
    hostname = forwardedHost;
  }

  let currentHost: string | undefined;
  const baseDomain = process.env.BASE_DOMAIN || "edezco.com";

  if (hostname.endsWith(baseDomain)) {
    currentHost = hostname.replace(`.${baseDomain}`, "");
  }

  console.log("Detected Hostname:", hostname);
  console.log("Extracted Subdomain:", currentHost);

  // If it's the main domain, continue normally
  if (!currentHost || currentHost === baseDomain) {
    return NextResponse.next();
  }

  // Fetch tenant-specific data based on the hostname
  // const tenantData = getTenantData(currentHost); // Replace with actual data fetching

  // if (!tenantData) {
  //   console.log(`No tenant found for ${currentHost}`);
  //   return NextResponse.rewrite(new URL("/not-found", request.url));
  // }

  console.log("Tenant Found:", currentHost);

  return NextResponse.rewrite(
    new URL(`/${currentHost}${url.pathname}`, request.url)
  );
}

function handleManagementRoute(request: NextRequest) {
  const isUserLogged = isValidLogin(); // Ensure this function is correctly defined
  if (!isUserLogged) {
    return NextResponse.redirect(new URL("/login", request.url));
  }

  const sideNavItem = SIDENAV_ITEMS.find((x) =>
    request.nextUrl.pathname.includes(x.path)
  );
  if (sideNavItem) {
    const user = getCurrentUserServer();
    if (user && !validateUserInRoles(sideNavItem.roleIds, user.roleId)) {
      return NextResponse.redirect(new URL("/management/home", request.url));
    }
  }

  return NextResponse.next();
}

// // Mock function for tenant data (Replace this with database or API call)
// function getTenantData(subdomain: string): any {
//   const mockTenants = {
//     test: { site_id: "test" },
//     demo: { site_id: "demo" },
//   };

//   return mockTenants[subdomain] || null;
// }

export const config = {
  matcher: ["/((?!api|_next/static|_next/image|favicon.ico).*)"],
};
