import type { NextApiRequest as Req, NextApiResponse as Res } from "next"
import { parseCookies, setCookie, destroyCookie } from "nookies"
import { zisuiSukiIdBaseUrl } from "../../../consts/config"
import { userError } from "../../../domain/user/errors"

const handler = async (req: Req, res: Res<{
    error: number, 
    userData: CurrentUser | null
}>) => {
    /* "POST"以外は、"404 Not Found"を返す */
    if (req.method !== "POST") {
        return res.status(404).send({
            error: userError.ServerError,
            userData: null,
        });
    }
    
    const {error, token, userData} = await requestLogin(req.body.name, req.body.password)
    if (error !== userError.OK) {
        return res.status(200).send({
            error: error,
            userData: null,
        })
    }

    /* set in cookie */
    const expiresIn = 60 * 10; // 10 minutes
    const options = {
        maxAge: expiresIn,
        httpOnly: true,
        secure: true,
        path: "/",
    };
    if (token) {
        setCookie({ res }, "token", token, options)
    }
    if (userData) {
        setCookie({ res }, "userId", userData.userId, options)
        setCookie({ res }, "name", userData.name, options)
        setCookie({ res }, "email", userData.email, options)
        setCookie({ res }, "displayName", userData.displayName, options)
        setCookie({ res }, "icon", userData.icon, options)
        setCookie({ res }, "status", userData.status.toString(), options)
    }

    /* response */
    return res.send({ 
        error: error,
        userData: userData,
    })
}

const requestLogin = async (name: string, password: string): Promise<{
    error: number,
    token: string | null,
    userData: CurrentUser | null,
}> => {
    const response = await fetch(zisuiSukiIdBaseUrl + "/login", { 
        method: "POST", 
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            name: name,
            password: password
        }),
    })
    if (response.status != 200) {
        return {error: userError.ServerError, token: null, userData: null}
    }

    const jsonData = await response.json()
    return {
        error: jsonData.status,
        token: jsonData.token,
        userData: jsonData.userInfo,
    }
}

export default handler