import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

// require('dotenv').config();

export async function middleware(request: NextRequest) {
  const token = request.cookies.get("jwt_token");
  console.log(token)
  if (!token) {
    return NextResponse.redirect(new URL("/login", request.url));
  }
  let tokenResponse;
  try {
    const API_URL = process.env.BACKEND_API; 
    console.log(API_URL)
    const myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    const response = await fetch(`${API_URL}/users/verify-token`, {
      method: "POST",
      headers: myHeaders,
      body: JSON.stringify({
        "token": token.value,
      }),
      cache:"no-cache"
    });

    tokenResponse = await response.json();
  }
  catch (error) {
    console.log(error);
  }

  if (!tokenResponse.valid) {
    return NextResponse.redirect(new URL("/login", request.url));
  }
  return NextResponse.next();
}

export const config = {
  matcher: "/" // Adjust the matcher to apply middleware to the desired routes
};
