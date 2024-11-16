
export const fetchJobs = async () => {
    // const JOBS_URL = "http://localhost:8000/jobs/";
    const API_URL = process.env.BACKEND_API


    const response = await fetch(`${API_URL}/jobs/`,{cache:"no-cache"});
    const jobs = await response.json();
    return jobs["jobs"];
};