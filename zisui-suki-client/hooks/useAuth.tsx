import { useRouter } from "next/router"
import { useState } from "react"
import { userError } from "../domain/user/errors"
import { validateUserName, validateUserPassword } from "../domain/user/validations"
import { useCurrentUser } from "./useCurrentUser"
import * as nextApi from "../service/nextApi/login"
import { useSetRecoilState } from "recoil"
import { currentUserState } from "../stores/currentUser"
import { useAuthErrorMessage } from "./useAuthErrorMessage"

export const useAuth = () => {
    const { isAuthChecking, currentUser } = useCurrentUser()
    const [errorCode, setErrorCode] = useState(userError.OK)
    const { loginErrorMessage, signupErrorMessage } = useAuthErrorMessage(errorCode)
    const setCurrentUser = useSetRecoilState(currentUserState)
    const router = useRouter()
    
    const login = async (arg: {name: string; password: string;}) => {
        /* get args */
        const {name, password} = arg

        /* input validation */
        let err
        err = validateUserName(name)
        if (err !== userError.OK) { 
            return setErrorCode(err)
        }
        err = validateUserPassword(password)
        if (err !== userError.OK) { 
            return setErrorCode(err)
        }

        /* fetch api */
        const { error, userData } = await nextApi.login(name, password)
        if (error !== userError.OK || !userData) {
            return setErrorCode(error)
        } else {
            setCurrentUser(userData)
            return router.push(`/${userData.name}`)
        }
    }

    return {isAuthChecking, currentUser, loginErrorMessage, signupErrorMessage, login}
}