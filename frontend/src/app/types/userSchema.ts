export type User = {
    message: string;
    token: string;
    user: SingleUser;
}
export type SingleUser={
    username: string;
    email: string;
    isAdmin: boolean;
}