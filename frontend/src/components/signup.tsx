"use client";
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
import { getCookie, setCookie } from "cookies-next";
import { userAtom } from "@/atom/userAtom";
import { useAtom } from "jotai";

export function SignupForm() {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("")
  const [password, SetPassword] = useState("")
  const [_, setUser] = useAtom(userAtom);
  const router=useRouter()


  const handleSubmit = async (e: any) => {
    e.preventDefault();
    try {
      const myHeaders = new Headers();
      myHeaders.append("Content-Type", "application/json");

      let user = await fetch("http://localhost:8000/users/signup", {
        method: "POST",
        headers: myHeaders,
        body: JSON.stringify({
          "username": username,
          "email": email,
          "password": password,
        }),
      });
      if (!user.ok) {
        throw Error("Signup Failed!")
      }
      const userInfo: User = await user.json();
      if (!userInfo.token) {
        throw Error(userInfo.message);
      }
      else {
        //signup
        const token = userInfo.token;
        setCookie("jwt_token", token);
        const user:SingleUser|null = userInfo.user;
        if (user) {
          setUser(user);
        }
        router.push("/")
        console.log(userInfo);
      }
    }
    catch (error: any) {
      console.log(error)
    }

  }



  return (
    <Card className="mx-auto max-w-sm">
      <CardHeader>
        <CardTitle className="text-xl">Sign Up</CardTitle>
        <CardDescription>
          Enter your information to create an account
        </CardDescription>
      </CardHeader>
      <CardContent>
        <form onSubmit={handleSubmit} className="grid gap-4">
          <div className="grid gap-2">
            <Label htmlFor="username">Username</Label>
            <Input
              id="username"
              placeholder="johnm1"
              required
              onChange={(e) => setUsername(e.target.value)}
            />
          </div>
          <div className="grid gap-2">
            <Label htmlFor="email">Email</Label>
            <Input
              id="email"
              type="email"
              placeholder="john.max@example.com"
              onChange={(e) => setEmail(e.target.value)}
              required
            />
          </div>
          <div className="grid gap-2">
            <Label htmlFor="password">Password</Label>
            <Input
              id="password"
              type="password"
              onChange={(e) => SetPassword(e.target.value)}
              required
            />
          </div>
          <Button type="submit" className="w-full">
            Create an account
          </Button>
          <Button variant="outline" className="w-full">
            Sign up with GitHub
          </Button>
        </form>
        <div className="mt-4 text-center text-sm">
          Already have an account?{" "}
          <Link href="/login" className="underline">
            Sign in
          </Link>
        </div>
      </CardContent>
    </Card>
  )
}