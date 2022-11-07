import { NextPage } from "next"
import { useRouter } from "next/router";
import { FormEvent, useEffect, useState } from "react";
import { useRecoilState, useResetRecoilState, useSetRecoilState } from "recoil";
import { baseUrl } from "../../const/config";
import { idResponseStatus } from "../../const/status";
import { currentUserState } from "../../stores/currentUser";
import { CurrentUser } from "../../types/user";
import { login } from "../../utils/login";

const LoginPage: NextPage = () => {

    const router = useRouter()
    const [name, setName] = useState("")
    const [password, setPassword] = useState("")
    const [errorMessage, setErrorMessage] = useState("")
    const setCurrentUser = useSetRecoilState(currentUserState)

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault()
        const { status, userData } = await login(name, password)
        switch (status) {
            case idResponseStatus.OK:
                if (!userData) {
                    setErrorMessage("サーバーエラー")
                } else {
                    setCurrentUser(userData)
                    router.push(`/${userData?.name}`)
                }
                break
            case idResponseStatus.NotExist:
                setErrorMessage("このユーザは存在しません")
                break
            case idResponseStatus.InvalidPassword:
                setErrorMessage("パスワードが間違っています")
                break
            default:
                setErrorMessage("サーバーエラー")
        }
        return
    }

    return (
        <div>
            <h1>Login Page</h1>
            <form onSubmit={handleSubmit}>
                <label htmlFor="name">名前：</label>
                <input 
                    id = "name"
                    value = {name}
                    onInput = {(e) => {setName(e.currentTarget.value)}}
                />
                <br/>
                <label htmlFor="password">パスワード：</label>
                <input 
                    id = "password"
                    type = "password"
                    value = {password}
                    onInput = {(e) => {setPassword(e.currentTarget.value)}}
                />
                <br/>
                {errorMessage}
                <br/>
                <button type="submit">login</button>
            </form>
        </div>
    )
}

export default LoginPage