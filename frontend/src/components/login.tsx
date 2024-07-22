"use client"; // This is a client component 👈🏽

import Link from "next/link"

import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useState } from "react"
import { SingleUser, User } from "@/app/types/userSchema";
import { useRouter } from "next/navigation";
import { setCookie,getCookie } from "cookies-next";

import { NavigationMenuDemo } from "@/components/navigation-menu";
import { userAtom } from "@/atom/userAtom";
import { useAtom } from "jotai";

export function LoginForm() {
  const router = useRouter();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [_, setUser] = useAtom(userAtom);


  const handleSubmit = async (e: any) => {
    e.preventDefault();
    try {
      const myHeaders = new Headers();
      myHeaders.append("Content-Type", "application/json");

      let user = await fetch("http://localhost:8000/users/login", {
        method: "POST",
        headers: myHeaders,
        body: JSON.stringify({
          "username": username,
          "password": password,
        }),
      });
      if (!user.ok) {
        throw new Error("Login Failed");
      }
      const userInfo: User = await user.json()
      if (!userInfo.token) {
        throw new Error(userInfo.message);
      }
      else {
        //login
        const token = userInfo.token;

        setCookie("jwt_token", token);
        const user:SingleUser|null = userInfo.user;
        if (user) {
          setUser(user);
        }
        router.push("/");
        
      }
    }
    catch (error: any) {
      console.log(error);
    }

  }

  return (<>
    <Card className="mx-auto max-w-sm">
      <CardHeader>
        <CardTitle className="text-2xl">Login</CardTitle>
        <CardDescription>
          Enter your username below to login to your account
        </CardDescription>
      </CardHeader>
      <CardContent>
        <form onSubmit={handleSubmit} className="grid gap-4" >
          <div className="grid gap-2">
            <Label htmlFor="username">Username</Label>
            <Input
              id="username"
              type="string"
              placeholder="johnmayer1"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
            />
          </div>
          <div className="grid gap-2">
            <div className="flex items-center">
              <Label htmlFor="password">Password</Label>
              <Link href="#" className="ml-auto inline-block text-sm underline">
                Forgot your password?
              </Link>
            </div>
            <Input id="password" type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
          </div>
          <Button type="submit" className="w-full">
            Login
          </Button>
          <Button variant="outline" className="w-full">
            Login with Google
          </Button>
        </form>
        <div className="mt-4 text-center text-sm">
          Don&apos;t have an account?{" "}
          <Link href="/signup" className="underline">
            Sign up
          </Link>
        </div>
      </CardContent>
    </Card>
    </>
  )
}
