import { NextRequest, NextResponse } from "next/server";
import isValidLogin from "./app/actions/isValidLogin";
import { SIDENAV_ITEMS } from "./app/constants";
import { validateUserInRoles } from "./app/utils/Utils";
import getCurrentUserServer from "./app/actions/getCurrentUserServer";

export function middleware(request: NextRequest) {
  const { pathname } = request.nextUrl;

  // ✅ Exclude static files and Next.js API from rewriting
  if (
    pathname.startsWith("/_next/") || // ✅ Exclude Next.js static files & images
    pathname.startsWith("/images/") || // ✅ Exclude images directory
    pathname === "/favicon.ico" || // ✅ Exclude favicon
    pathname.startsWith("/api/") // ✅ Exclude API requests
  ) {
    return NextResponse.next();
  }

  if (pathname.startsWith("/management")) {
    return handleManagementRoute(request);
  }

  return handleSubdomains(request);
}

function handleSubdomains(request: NextRequest) {
  console.log("---- Handling Subdomain ----");

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

  // ✅ If it's the main domain, continue normally
  if (!currentHost || currentHost === baseDomain) {
    return NextResponse.next();
  }

  console.log("✅ Tenant Found:", currentHost);

  return NextResponse.rewrite(
    new URL(`/${currentHost}${url.pathname}`, request.url)
  );
}

function handleManagementRoute(request: NextRequest) {
  const isUserLogged = isValidLogin();
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

// ✅ Ensure Next.js doesn't process static files
export const config = {
  matcher: ["/((?!_next/|images/|favicon.ico|api/).*)"],
};