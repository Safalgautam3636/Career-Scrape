
"use client";
import {
    Card,
    CardContent,
    CardDescription,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import { Input } from "@/components/ui/input"
import Link from "next/link"
import { useEffect, useState } from "react"
import { Toggle } from "@/components/ui/toggle"
import { Switch } from "./ui/switch";
import { useRouter } from "next/navigation";
import { useAtom } from "jotai";
import { userAdminUpdateWithAtomStorage } from "@/atom/userAtom";
import { SingleUserWithId } from "../../types/userSchema";

export function UpdateUser() {

    const [userToBeUpdated] = useAtom(userAdminUpdateWithAtomStorage)

    const [form_username, setUsername] = useState("")
    const [form_email, setEmail] = useState("")
    const [form_isAdmin, setAdmin] = useState(false)
    const [form_userId, setUserId] = useState("")
    const [form_password, setPassword] = useState("")
    useEffect(() => {
        if (userToBeUpdated) {
            const { username, email, isAdmin, id, password } = userToBeUpdated as SingleUserWithId
            setUsername(username);
            setEmail(email);
            setAdmin(isAdmin);
            setUserId(id)
            setPassword(password)
        }

    }, [userToBeUpdated])

    const handleSubmitForm = async (e: any) => {
        e.preventDefault()
        const myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");
        // TODO: FIX ID
        const id = form_userId
        let user = await fetch(`http://localhost:8000/users/${id}`, {
            method: "PUT",
            headers: myHeaders,
            body: JSON.stringify({
                "username": form_username,
                "email": form_email,
                "isAdmin": form_isAdmin,
                "password": form_password,
                "id": id
            }),
        });
        user = await user.json()
        user = user.newUser
        const username = user.username;
        const email = user.email;
        const isAdmin = user.isAdmin;
        setUsername(username)
        setEmail(email)
        setAdmin(isAdmin)
        //TODO: make sucesspage
        alert("Updated")
    }

    return (
        <Card className="mx-auto max-w-sm">
            <CardHeader>
                <CardTitle className="text-xl">Edit User Information</CardTitle>
                <CardDescription>
                    Modify information to update an account
                </CardDescription>
            </CardHeader>
            <CardContent>
                <form className="grid gap-4" onSubmit={handleSubmitForm}>
                    <div className="grid gap-2">
                        <Label htmlFor="username">Username</Label>
                        <Input
                            id="username"
                            placeholder="johnm1"
                            value={form_username}
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
                            value={form_email}
                            onChange={(e) => setEmail(e.target.value)}
                            required
                        />
                    </div>
                    <div className="grid gap-2">
                        Admin Status<Switch checked={form_isAdmin} onCheckedChange={(e) => setAdmin(!form_isAdmin)} />
                    </div>
                    <Button type="submit" className="w-full">
                        Update
                    </Button>

                </form>
            </CardContent>

        </Card>
    )
}