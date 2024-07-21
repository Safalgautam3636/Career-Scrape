import { Job } from "@/components/job";
import { JobSchema } from "./types/JobSchema";
import { getCookie } from "cookies-next";
import { redirect } from "next/navigation";
// import { useRouter } from "next/navigation";

const fetchJobs = async () => {
  const JOBS_URL = "http://localhost:8000/jobs/";
  const response = await fetch(JOBS_URL);
  const jobs = await response.json();
  return jobs["jobs"];
};

export default async function Home() {
  // const route = useRouter()
  const jobs: JobSchema[] = await fetchJobs();
  //console.log(jobs.map)
  // return <div>Hi mf{ jobs[0].CreatedAt}</div>
  const token = getCookie("jwt_token");
  
  return (
      <div className="grid grid-cols-4 p-5">
        {jobs && jobs.map((job) => (
          <Job key={job.id} {...job} />
        ))}
      </div>)


}
