import type { NextApiRequest as Req, NextApiResponse as Res } from "next";
import { parseCookies, setCookie, destroyCookie } from "nookies"
import { currentUserState } from "../../../stores/currentUser";
import { CurrentUser } from "../../../types/user";

const handler = async (req: Req, res: Res<CurrentUser | null>) => {
    const cookies = parseCookies({ req })
    if (!cookies.token || !cookies.userId || !cookies.name || !cookies.displayName || !cookies.email || !cookies.icon || !cookies.status) {
        res.status(401).send(null)
    }

    res.status(200).json({
        userId: cookies.userId,
        name: cookies.name,
        displayName: cookies.displayName,
        email: cookies.email,
        icon: cookies.icon,
        status: parseInt(cookies.status),
    })
}

export default handler