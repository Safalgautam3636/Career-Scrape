import { fetchJobs } from "@/lib/fetchJobs"
import { JobSchema } from "@/types/JobSchema";
import { Job } from "@/components/job";
export default async function Jobs() {
    let jobs: JobSchema[] = await fetchJobs();
    return (
      <div className="grid grid-cols-4 p-5">
            {jobs && jobs.map((job) => (
                
          <Job key={job.id} {...job} />
        ))}
      </div>)
}