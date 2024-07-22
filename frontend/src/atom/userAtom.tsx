import { atom } from "jotai";
import { SingleUser } from "@/app/types/userSchema";

export const userAtom = atom<SingleUser|null>(null);