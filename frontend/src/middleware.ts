import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

export async function middleware(request: NextRequest) {
  const token = request.cookies.get("jwt_token");
  if (!token) {
    return NextResponse.redirect(new URL("/login", request.url));
  }
  const VERIFY_TOKEN = "http://localhost:8000/users/verify-token"
  const myHeaders = new Headers();
  myHeaders.append("Content-Type", "application/json");

  const response = await fetch(VERIFY_TOKEN, {
    method: "POST",
    headers: myHeaders,
    body: JSON.stringify({
      "token":token.value,
    }),
  });

  const tokenResponse = await response.json();
  if (!tokenResponse.valid) {
    return NextResponse.redirect(new URL("/login", request.url));
  }
  return NextResponse.next();
}

export const config = {
  matcher: "/" // Adjust the matcher to apply middleware to the desired routes
};
