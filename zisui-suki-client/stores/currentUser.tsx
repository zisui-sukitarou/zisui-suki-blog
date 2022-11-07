import { atom } from "recoil"
import { CurrentUser } from "../types/user"

export const currentUserState = atom<undefined | null | CurrentUser>({
    key: "currentUserState",
    default: undefined,
})