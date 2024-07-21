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

interface JobSchema {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string;
    id: string;
    job_title: string;
    job_link: string;
    job_location: string;
    job_posted: string;
    company_name: string;
    exact_date: string;
    job_type: string;
    company_domain: string;
    job_level: string;
    company_link: string;
    description: string;
    pulled_date: string;
}



export function Job(prop: JobSchema) {
    return (
        <Card className="w-[350px] mb-2">
            <CardHeader>
                <CardTitle>
                    {prop.job_title}
                </CardTitle>
                <CardDescription className="flex justify-between items-center w-full">
                    <Link href={`${prop.company_link}`} className="text-lg">{prop.company_name}</Link>
                    <div>{prop.job_location}</div>
                    
                </CardDescription>
            </CardHeader>

            <CardContent className="grid gap-4">
                 {/* TODO: hydration bug fix it! */}
                <CardDescription className="flex justify-between items-center w-full">
                    <div className="text-medium">{prop.job_level}</div>
                </CardDescription>
                <CardDescription className="flex justify-between items-center w-full">
                    <div className="text-xs font-extrabold">{prop.job_posted}</div>
                    <div className="text-xs font-extrabold"> {prop.job_type}</div>
                    <div className="text-xs font-extrabold">{prop.company_domain}</div>
                </CardDescription>
                <CardDescription className="flex justify-between items-center w-full">
                    <div className="text-xs">Posted:{prop.exact_date}</div>
                    <div className="text-xs">Pulled:{prop.pulled_date}</div>
                </CardDescription>
            </CardContent>

            <CardFooter>
                <Button className="w-full" asChild>
                    <Link href={`${prop.job_link}`}><CheckIcon className="mr-2 h-4 w-4" />Apply</Link>
                </Button>
            </CardFooter>

        </Card>
    )
}