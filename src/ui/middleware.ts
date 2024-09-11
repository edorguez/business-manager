import { NextRequest, NextResponse } from "next/server";
import isValidLogin from "./app/actions/isValidLogin";

export function middleware(request: NextRequest) {
  const isUserLogged: boolean = isValidLogin();

  if (isUserLogged) return NextResponse.next();

  return NextResponse.redirect(new URL("/login", request.url));
}

export const config = {
  matcher: ["/management/:path*"], // Routes that should use this middleware
};
