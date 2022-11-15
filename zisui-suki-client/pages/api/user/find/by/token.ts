import type { NextApiRequest as Req, NextApiResponse as Res } from "next";
import { parseCookies, setCookie, destroyCookie } from "nookies"
import { zisuiSukiIdBaseUrl } from "../../../../../consts/config";

const handler = async (req: Req, res: Res) => {
    const cookies = parseCookies({ req })
    const session = cookies.session || ""
    const userData = await fetchUserData(session)
    res.json(userData)
}

const fetchUserData = async (session: string) => {
    const response = await fetch(zisuiSukiIdBaseUrl + "/find/by/token", {
        method: "GET",
        headers: {
            "Authorization": "Bearer " + session
        }
    })
    return await response.json()
}

export default handler
