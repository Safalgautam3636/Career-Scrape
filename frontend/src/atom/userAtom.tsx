import { SingleUser } from "../types/userSchema";
import { atomWithStorage } from "jotai/utils";
export const userWithAtomStorage = atomWithStorage<SingleUser | null>("userInfo", null);
export const userAdminUpdateWithAtomStorage=atomWithStorage<SingleUser|null>("userToBeUpdated",null)

