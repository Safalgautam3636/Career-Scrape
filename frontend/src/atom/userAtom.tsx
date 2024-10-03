import { atom } from "jotai";
import { SingleUser } from "../../types/userSchema";
import { atomWithStorage } from "jotai/utils";
export const userAtom = atom<SingleUser | null>(null);
export const userWithAtomStorage = atomWithStorage<SingleUser | null>("userInfo", null);
export const userAdminUpdateWithAtomStorage=atomWithStorage<SingleUser|null>("userToBeUpdated",null)

