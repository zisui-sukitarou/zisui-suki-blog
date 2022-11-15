import { useEffect, useState } from "react"
import { userError } from "../domain/user/errors"

export const useAuthErrorMessage = (errorCode: number) => {
    const [loginErrorMessage, setLoginErrorMessage] = useState("")
    const [signupErrorMessage, setSignupErrorMessage] = useState("")

    useEffect( () => {
        switch (errorCode) {
            case userError.OK:
                setLoginErrorMessage("")
                setSignupErrorMessage("")
                break
            case userError.InvalidUserName:
                setLoginErrorMessage("ユーザー名が違います。")
                setSignupErrorMessage("ユーザー名は、英字、数字、'_'、'-'からなる6文字以上15文字以下である必要があります。")
                break
            case userError.InvalidPassword:
                setLoginErrorMessage("パスワードが違います。")
                setSignupErrorMessage("パスワードは、英字、数字、'_'、'-'、'@'、'!'、'?'、からなる6文字以上15文字以下である必要があります。")
                break
            case userError.UserAlreadyExists:
                setSignupErrorMessage("このユーザー名のユーザーはすでに存在しています。")
                break
            case userError.IncorrectPassword:
                setLoginErrorMessage("パスワードが違います。")
                break
            case userError.UserNotExists:
                setLoginErrorMessage("ユーザー名が違います。")
                break
            case userError.ServerError:
                setLoginErrorMessage("サーバーエラー")
                setSignupErrorMessage("サーバーエラー")
                break
            default:
                setLoginErrorMessage("想定外エラー")
                setSignupErrorMessage("想定外エラー")
        }
    }, [errorCode])

    return {loginErrorMessage, signupErrorMessage}
}