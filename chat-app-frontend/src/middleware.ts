import { NextRequest, NextResponse } from "next/server";
const PRIVATE_ROUTES = ["/", "/home"];

export default async function middleware(req: NextRequest) {
  const token = req.cookies.get("chatapp-session")?.value;
  const url = req.nextUrl.clone();
  const pathname = url.pathname;

  if (!PRIVATE_ROUTES.includes(pathname)) {
    return NextResponse.next();
  }

  if (!token) {
    url.pathname = "/login";
    return NextResponse.redirect(url);
  }

  const response = await fetch("http://localhost:1323/v1/auth/verify", {
    method: "POST",
    headers: req.headers,
  });

  if (response.status != 200) {
    url.pathname = "/login";
    return NextResponse.redirect(url);
  }

  const res = NextResponse.next();

  const json = await response.json()
  res.headers.set("x-user-id", json.userId);

  return res;
}
