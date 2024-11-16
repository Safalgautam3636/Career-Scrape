export const fetchUsers = async () => {
    // const JOBS_URL = "http://localhost:8000/users";
    const API_URL = process.env.BACKEND_API;
    const response = await fetch(`${API_URL}/users`,{cache:"no-cache"});
    const users = await response.json();
    return users["users"];
};