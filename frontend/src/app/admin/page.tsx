
import {
    Table,
    TableBody,
    TableCaption,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table"
import { UserForAdmin } from "../../types/userSchema"
import {AdminButton} from "@/components/admin-button"
import { Done } from "@/components/ui/done";
import { Incorrect } from "@/components/ui/incorrect";
import { userAdminUpdateWithAtomStorage } from "@/atom/userAtom";
import { useAtom } from "jotai";

export default async function Admin() {
    const API_URL = process.env.BACKEND_API;
    let response: any = await fetch(`${API_URL}/users`, { cache: 'no-store' });
    response = await response.json();
    const users: UserForAdmin[] = response["users"];
    
   
    return (
         
        
        <Table>
            <TableCaption>Users in the portal</TableCaption>
            <TableHeader>
                <TableRow>
                    <TableHead className="w-[100px]">User Id</TableHead>
                    <TableHead>Username</TableHead>
                    <TableHead>Email</TableHead>
                    <TableHead>Admin</TableHead>
                    <TableHead className="text-right">Action</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                {users && users.map((user) => (<TableRow key={user.id}>
                    
                    <TableCell className="font-medium">{user.id}</TableCell>
                    <TableCell>{user.username}</TableCell>
                    <TableCell>{user.email}</TableCell>
                    <TableCell>{user.isAdmin?<Done/>:<Incorrect/>}</TableCell>
                    <TableCell className="text-right">
                        <AdminButton user={user} href={`/admin/users/${user.id}/edit`}/>
                    </TableCell>
                </TableRow>))}

            </TableBody>
        </Table>
    )
}