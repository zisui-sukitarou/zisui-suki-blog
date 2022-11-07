import type { NextApiRequest as Req, NextApiResponse as Res } from "next";
import { parseCookies, setCookie, destroyCookie } from "nookies"
import { json } from "stream/consumers";
import { zisuiSukiIdBaseUrl } from "../../../const/config";
import { idResponseStatus } from "../../../const/status";
import { CurrentUser } from "../../../types/user";

const handler = async (req: Req, res: Res<{
    status: number | null, 
    userData: CurrentUser | null
}>) => {
    /* "POST"以外は、"404 Not Found"を返す */
    if (req.method !== "POST") {
        return res.status(404).send({
            status: null,
            userData: null,
        });
    }
    
    const {status, token, userData} = await requestLogin(req.body.name, req.body.password)
    if (status !== idResponseStatus.OK) {
        return res.status(200).send({
            status: status,
            userData: null,
        })
    }

    /* set in cookie */
    const expiresIn = 60 * 1; // 1 minutes
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
        status: status,
        userData: userData,
    })
}

const requestLogin = async (name: string, password: string): Promise<{
    status: number | null,
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
        return {status: null, token: null, userData: null}
    }

    const jsonData = await response.json()
    return {
        status: jsonData.status,
        token: jsonData.token,
        userData: jsonData.userInfo,
    }
}

export default handler