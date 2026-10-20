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

function normalizeHostname(value: string | null): string | undefined {
  if (!value) return undefined;

  // Cloudflare can send a comma-separated list and local tests may include :port.
  return value.split(",")[0].trim().split(":")[0].toLowerCase() || undefined;
}

function handleSubdomains(request: NextRequest) {
  const url = request.nextUrl;

  const hostname =
    normalizeHostname(request.headers.get("x-forwarded-host")) ??
    normalizeHostname(request.headers.get("host"));

  if (!hostname) {
    return NextResponse.next(); // No host header, continue normally
  }

  const baseDomain = (
    process.env.BASE_DOMAIN || "edoutput.dev"
  ).toLowerCase();

  const reservedSubdomains = (
    process.env.RESERVED_SUBDOMAINS ||
    "www,tourism-agency,gps-medical-assistance,business-manager"
  )
    .split(",")
    .map((value) => value.trim().toLowerCase())
    .filter(Boolean);

  // The apex domain is the future portfolio, not a tenant.
  if (hostname === baseDomain) {
    return NextResponse.next();
  }

  let currentHost: string | undefined;

  if (hostname.endsWith(`.${baseDomain}`)) {
    currentHost = hostname.slice(0, -`.${baseDomain}`.length);
  }

  console.log("Detected Hostname:", hostname);
  console.log("Extracted Subdomain:", currentHost);

  // Main app and explicit project subdomains must not be rewritten as tenants.
  if (!currentHost || reservedSubdomains.includes(currentHost)) {
    return NextResponse.next();
  }

  console.log("Tenant Found:", currentHost);

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