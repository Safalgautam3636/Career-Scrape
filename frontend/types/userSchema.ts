import { UUID } from "crypto";

export type User = {
    message: string;
    token: string;
    user: SingleUser;
}
export type SingleUser = {
    username: string;
    email: string;
    isAdmin: boolean;
}
export type SingleUserWithId = {
    username: string,
    email: string,
    isAdmin: boolean,
    id: string;
    password: string;
}
export type UserForAdmin = {
    ID: number,
    createdAt: string,
    updatedAt: string,
    deletedAt: string,
    id: string,
    username: string,
    email: string,
    password: string,
    isAdmin: boolean
}