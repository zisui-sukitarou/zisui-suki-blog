import { useRecoilValue } from "recoil"
import { currentUserState } from "../stores/currentUser"

export const useCurrentUser = () => {
    const currentUser = useRecoilValue(currentUserState)
    const isAuthChecking = currentUser === undefined

    return {
        currentUser,
        isAuthChecking
    }
}