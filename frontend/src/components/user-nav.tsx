import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { useState } from "react";
import { setCookie,getCookie,deleteCookie } from "cookies-next";
import { useAtom } from "jotai";
import { userWithAtomStorage } from "@/atom/userAtom";
import { useRouter } from "next/navigation";

export function UserNav({ email, name, userImage }: {
  email: string;
  name: string;
  userImage: string | undefined;
}) {
  const [userInfo, setUserInfo] = useAtom(userWithAtomStorage);
  const router = useRouter();
  const logout = () => {
    if (getCookie("jwt_token")) {
      deleteCookie("jwt_token")
    }
    if (userInfo) {
      setUserInfo(null)
    }
    router.push("/login")
    router.refresh()
  }
  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="ghost" className="relative h-10 w-10 rounded-full">
          <Avatar className="h-10 w-10">
            <AvatarImage src={userImage} alt="User Image" />
            <AvatarFallback>{name[0]}</AvatarFallback>
          </Avatar>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" className="w-56" forceMount>
        <DropdownMenuLabel className="font-normal">
          <div className="flex flex-col space-y-1">
            <p className="text-sm font-medium leading-none">{name}</p>
            <p className="text-xs leading-none text-muted-foreground">
              {email}
            </p>
          </div>
        </DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuGroup>
          <DropdownMenuItem>
            <button onClick={logout}>Signout</button></DropdownMenuItem>
          <DropdownMenuItem>test</DropdownMenuItem>
          <DropdownMenuItem>test</DropdownMenuItem>
          <DropdownMenuItem>test</DropdownMenuItem>
        </DropdownMenuGroup>
        <DropdownMenuSeparator />
        <DropdownMenuItem asChild>
          Logout
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
