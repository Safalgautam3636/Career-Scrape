import * as React from "react"

import { Button } from "@/components/ui/button"
import {
    Card,
    CardContent,
    CardDescription,
    CardFooter,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { BellIcon, CheckIcon } from "@radix-ui/react-icons"

import Link from "next/link"

interface Job {
    id: string,
    role: string,
    link: string,
    location: string,
    company: string,
    dateposted: string,
    datepulled: string,
    experience: string,
    skills: string[],
    compensation: string,
    jobtype: string
}
//experience, skills, compansation, jobtype-fulltime/internship/remote/hybrid/onsite
const notifications = [
    {
        title: "Your call has been confirmed.",
        description: "1 hour ago",
    },
    {
        title: "You have a new message!",
        description: "1 hour ago",
    },
    {
        title: "Your subscription is expiring soon!",
        description: "2 hours ago",
    },
]




export function Job(prop: Job) {
    return (
        <Card className="w-[350px] mb-2">
            <CardHeader>
                <CardTitle>
                    {prop.role}
                </CardTitle>
                <CardDescription className="flex justify-between items-center w-full">
                    <div className="text-lg">{prop.company}</div>
                    <div>{prop.location}</div>
                </CardDescription>
            </CardHeader>

            <CardContent className="grid gap-4">
                <CardDescription className="flex justify-between items-center w-full">
                    <div className="text-medium">{prop.skills.join(", ")} work experience </div>
                </CardDescription>
                <CardDescription className="flex justify-between items-center w-full">
                    <div className="text-xs font-extrabold">{prop.jobtype}</div>
                    <div className="text-xs font-extrabold">Atleast {prop.experience} {"yr" ? prop.experience == "1" : "yrs"}</div>
                    <div className="text-xs font-extrabold">{prop.compensation}</div>
                </CardDescription>
                <CardDescription className="flex justify-between items-center w-full">
                    <div className="text-xs">Posted:{prop.dateposted}</div>
                    <div className="text-xs">Pulled:{prop.datepulled}</div>
                </CardDescription>
            </CardContent>

            <CardFooter>
                <Button className="w-full" asChild>
                    <Link href={"/jobs"}><CheckIcon className="mr-2 h-4 w-4" />Apply</Link>
                </Button>
            </CardFooter>

        </Card>
    )
}