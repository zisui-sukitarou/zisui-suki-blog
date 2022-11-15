import { NextPage } from "next"
import { useRouter } from "next/router";
import { FormEvent, useEffect, useState } from "react";
import { useRecoilState, useResetRecoilState, useSetRecoilState } from "recoil";
import { baseUrl } from "../../consts/config";
import { currentUserState } from "../../stores/currentUser";
import { login } from "../../service/nextApi/login";
import { userError } from "../../domain/user/errors";
import { useAuth } from "../../hooks/useAuth";

const LoginPage: NextPage = () => {

    const {login, loginErrorMessage} = useAuth()
    const [name, setName] = useState("")
    const [password, setPassword] = useState("")
    
    const onSubmit = async (e: FormEvent) => {
        e.preventDefault()
        login({name: name, password: password})
    }

    return (
        <div>
            <h1>Login Page</h1>
            <form onSubmit={onSubmit}>
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
                {loginErrorMessage}
                <br/>
                <button type="submit">login</button>
            </form>
        </div>
    )
}

export default LoginPage