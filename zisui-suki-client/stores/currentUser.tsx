import { atom } from "recoil"

export const currentUserState = atom<undefined | null | CurrentUser>({
    key: "currentUserState",
    default: undefined,
})