import type { NextApiRequest as Req, NextApiResponse as Res } from "next"
import { parseCookies, setCookie, destroyCookie } from "nookies"
import { draftError } from "../../../../../../domain/draft/error";
import { findDraftsByUserId } from "../../../../../../service/blogApi/draft/find/by/userId"

const handler = async (req: Req, res: Res<{
    error: number, 
    drafts: Array<DraftOverview>,
}>) => {
    /* "POST"以外は、"404 Not Found"を返す */
    if (req.method !== "POST") {
        return res.status(404).send({
            error: draftError.ServerError,
            drafts: [],
        });
    }

    /* calc begin & end of drafts from query param */
    const { page } = req.query
    const begin = (Number(page) - 1) * 10 | 0
    const end = Number(page) * 10 - 1 | 9

    /* extract token */
    const cookies = parseCookies({ req })
    if (!cookies.token || !cookies.userId || !cookies.name || !cookies.displayName || !cookies.email || !cookies.icon || !cookies.status) {
        res.status(401).send({
            error: draftError.ServerError,
            drafts: [],
        })
    }

    /* get my drafts */
    const {error, drafts} = await findDraftsByUserId(cookies.token, cookies.userId, begin, end)

    /* response */
    return res.send({ 
        error: error,
        drafts: drafts,
    })
}