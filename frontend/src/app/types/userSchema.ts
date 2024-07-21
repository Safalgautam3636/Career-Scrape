export type User = {
    message: string;
    token: string;
    user: {
        username: string;
        email: string;
        isAdmin: boolean;
    }
}