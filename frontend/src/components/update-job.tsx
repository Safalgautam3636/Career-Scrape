
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
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "./ui/select";
import { useEffect, useState } from "react"
import { useAtom } from "jotai";
import { jobWithAtomStorage } from "@/atom/jobAtom";
import { JobSchema } from "../types/JobSchema";

export function UpdateJob() {
    const [jobTitle, setJobTitle] = useState("")
    const [jobLink, setJobLink] = useState("");
    const [jobLocation, setJobLocation] = useState("");
    const [jobPosted, setJobPosted] = useState("") // how many hours ago?
    const [companyName, setCompanyName] = useState("")
    const [exactDate, setExactDate] = useState("") // when (date) was it posted?
    const [jobType, setJobType] = useState("")
    const [comapanyDomain, setCompanyDomain] = useState("")
    const [jobLevel, setJobLevel] = useState("")
    const [companyLink, setCompanyLink] = useState("")
    const [des, setDescription] = useState("")
    const [pulledDate, setPulledDate] = useState("")  //when was it scraped?
    const [jobId, setId] = useState("")
    const [job] = useAtom(jobWithAtomStorage)
    function standardDate(date: string) {
        const [year, month, day] = date.split("-");
        return `${month}-${day}-${year}`
    }
    function formatDate(exact_date: string) {
        const date = new Date(exact_date);
        const year = String(date.getFullYear())
        const month = String(date.getMonth());
        const day = String(date.getDay());
        const htmlFormatExactDate = `${year}-${month.length !== 2 ? '0' + month : month}-${day.length !== 2 ? '0' + day : day}`
        return htmlFormatExactDate;
    }
    function formatTimeAndDate(pulled_date: string) {
        const date_time = new Date(pulled_date);
        const hours = String(date_time.getUTCHours()).padStart(2, '0');
        const minutes = String(date_time.getUTCMinutes()).padStart(2, '0');
        const date = formatDate(date_time.toISOString().split("T")[0])
        const formattedDate = `${date}T${hours}:${minutes}`;
        return formattedDate
    }
    function getJobLevel(job_level: String | null) {
        if (job_level === null) {
            return "";
        }
        else {
            let level = "";
            if (job_level.includes("mid") || job_level.includes("senior")) {
                level = "mid-level"
            }
            else if (job_level.includes("manager")) {
                level = "manager"
            }
            else if (job_level.includes("junior")) {
                level = "entry-level"
            }
            return level
        }
    }
    function parseJobType(job_type: String | null) {
        if (job_type === null) return "";
        let type = ""
        if (job_type.includes("full")) {
            type = "full-time"
        }
        else if (job_type.includes("part")) {
            type = "part-time";
        }
        else if (job_type.includes("contract")) {
            type = "contract";
        }
        return type
    }

    useEffect(() => {
        if (job) {
            const { id, exact_date, pulled_date, job_title, job_link, job_location, job_posted, company_name, company_domain, company_link, job_level, job_type, description } = job;
            setId(id);
            setExactDate(formatDate(exact_date));
            setPulledDate(formatTimeAndDate(pulled_date));
            setJobTitle(job_title);
            setJobLink(job_link);
            setJobLocation(job_location);
            setJobPosted(job_posted);
            setCompanyName(company_name);
            setCompanyDomain(company_domain);
            setCompanyLink(company_link);
            setJobLevel(getJobLevel(job_level));
            setJobType(parseJobType(job_type));
            setDescription(description); 
        }

    }, [job])

    const handleSubmitForm = async (e: any) => {
        e.preventDefault()
        const myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");
        // TODO: FIX ID
        const id = jobId;
        const updateJobObject = {
            "id": jobId,
            "job_title": jobTitle,
            "job_link": jobLink,
            "job_location": jobLocation,
            "job_posted": jobPosted,
            "company_name": companyName,
            "exact_date": new Date(standardDate(exactDate) + "T15:04:05Z07:00"),
            "job_type": jobType,
            "company_domain": comapanyDomain,
            "job_level": jobLevel,
            "company_link": companyLink,
            "description": des,
            "pulled_date": new Date(pulledDate).toISOString()
        }
        const API_URL = process.env.BACKEND_API;
        const response = await fetch(`${API_URL}/jobs/${id}`, {
            method: "PUT",
            headers: myHeaders,
            body: JSON.stringify(updateJobObject),
            cache:"no-cache"
        });
        const job = await response.json();
        const jobObj:JobSchema = job.job as JobSchema;
        const { job_title, job_link, job_location, job_posted, company_name, exact_date, job_type, company_domain, job_level, company_link, description, pulled_date } = jobObj;
        console.log(exact_date);
        console.log(pulled_date);
        // setJobTitle(job_title);
        // setJobLink(job_link);
        // setJobLocation(job_location);
        // setJobPosted(job_posted);
        // setCompanyName(company_name);
        // setExactDate(exact_date);
        // setJobType(job_type);
        // setCompanyDomain(company_domain);
        // setJobLevel(job_level);
        // setCompanyLink(company_link);
        // setDescription(description);
        // setPulledDate(pulled_date);
    }
    return (
        <Card className="mx-auto max-w-xl">
            <CardHeader>
                <CardTitle className="text-xl">Update the Job information</CardTitle>
                <CardDescription>
                    Modify information to update a job
                </CardDescription>
            </CardHeader>
            <CardContent>
                <form className="grid gap-4" onSubmit={handleSubmitForm}>
                    <div className="grid gap-2">
                        <Label htmlFor="title">Title</Label>
                        <Input
                            id="title"
                            placeholder="Software Engineer"
                            value={jobTitle}
                            required
                            onChange={(e) => setJobTitle(e.target.value)}
                        />
                    </div>
                    <div className="grid gap-2">
                        <Label htmlFor="email">Link</Label>
                        <Input
                            id="email"
                            type="text"
                            placeholder="https://google.com/careers/apply/3213-312-42e3-a29c"
                            value={jobLink}
                            onChange={(e) => setJobLink(e.target.value)}
                            required
                        />
                    </div>
                    <div className="flex gap-2">
                        <div>
                            <Label htmlFor="location">Location</Label>
                            <Input
                                id="location"
                                type="text"
                                placeholder="New York, New York"
                                value={jobLocation}
                                onChange={(e) => setJobLocation(e.target.value)}
                                required
                            />
                        </div>
                        <div>
                            <Label htmlFor="company">Company Name</Label>
                            <Input
                                id="company"
                                type="text"
                                placeholder="Google"
                                value={companyName}
                                onChange={(e) => setCompanyName(e.target.value)}
                                required
                            />
                        </div>

                    </div>
                    <div className="flex gap-2">
                        <div>
                            <Label htmlFor="domain">Company Domain</Label>
                            <Input
                                id="domain"
                                type="text"
                                placeholder="Information Tech."
                                value={comapanyDomain}
                                onChange={(e) => setCompanyDomain(e.target.value)}
                                required
                            />
                        </div>
                        <div>
                            <Label htmlFor="pulled">Pulled Date</Label>
                            <Input
                                id="pulled"
                                type="datetime-local"
                                value={pulledDate}
                                onChange={(e) => {
                                    setPulledDate(e.target.value)
                                }}
                                required
                            />
                        </div>

                    </div>
                    <div className="grid gap-2">
                        <Label htmlFor="link">Company Link</Label>
                        <Input
                            id="link"
                            type="text"
                            placeholder="http://google.com"
                            value={companyLink}
                            onChange={(e) => setCompanyLink(e.target.value)}
                            required
                        />
                    </div>
                    <div className="flex gap-2">
                        <div>
                            <Label htmlFor="posted">Posted Hours (if new)</Label>
                            <Input
                                id="posted"
                                type="text"
                                placeholder="5 hrs ago.."
                                value={jobPosted}
                                onChange={(e) => setJobPosted(e.target.value)}
                                required
                            />
                        </div>
                        <div>
                            <Label htmlFor="date">Posted Date</Label>
                            <Input
                                id="date"
                                type="date"
                                value={exactDate}
                                onChange={(e) => setExactDate(e.target.value)}
                                required
                            />
                        </div>

                    </div>
                    <div className="grid gap-2">
                        <Label htmlFor="link">Job Description</Label>
                        <Input
                            id="description"
                            type="text"
                            placeholder="Proficient in Java, C++, Python and Go.."
                            value={des}
                            onChange={(e) => setDescription(e.target.value)}
                            required
                        />
                    </div>
                    <div className="flex gap-2">
                        <div>
                            <Label htmlFor="level">Job Level</Label>
                            <Select value={jobLevel} onValueChange={(value) => setJobLevel(value)}>
                                <SelectTrigger className="w-[180px]">
                                    <SelectValue placeholder="Select" />
                                </SelectTrigger>
                                <SelectContent >
                                    <SelectItem value="entry-level">Entry Level</SelectItem>
                                    <SelectItem value="mid-level">Mid Level</SelectItem>
                                    <SelectItem value="manager">Manager</SelectItem>
                                </SelectContent>
                            </Select>

                        </div>
                        <div>
                            <Label htmlFor="type">Job Type</Label>
                            <Select value={jobType} onValueChange={(value) => setJobType(value)}>
                                <SelectTrigger className="w-[180px]">
                                    <SelectValue placeholder="Select" />
                                </SelectTrigger>
                                <SelectContent>
                                    <SelectItem value="full-time">Full Time</SelectItem>
                                    <SelectItem value="part-time">Part Time</SelectItem>
                                    <SelectItem value="contract">Contract</SelectItem>
                                </SelectContent>
                            </Select>
                        </div>

                    </div>

                    <Button type="submit" className="w-full">
                        Update
                    </Button>

                </form>
            </CardContent>

        </Card>
    )
}