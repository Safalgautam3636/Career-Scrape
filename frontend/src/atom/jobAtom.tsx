import { atomWithStorage } from "jotai/utils";
import { JobSchema } from "../types/JobSchema";

export const jobWithAtomStorage = atomWithStorage<JobSchema | null>("jobInfo", null);
