import type { NextApiRequest as Req, NextApiResponse as Res } from "next"
import { parseCookies } from "nookies"
import { draftError } from "../../../../../../domain/draft/error";
import * as blogApi from "../../../../../../service/blogApi/draft/find/by/userId"

const handler = async (req: Req, res: Res<{
    error: number, 
    drafts: Array<DraftOverview>,
}>) => {
    /* "GET"以外は、"404 Not Found"を返す */
    if (req.method !== "GET") {
        return res.status(404).send({
            error: draftError.ServerError,
            drafts: [],
        });
    }

    /* extract token */
    const cookies = parseCookies({ req })
    if (!cookies.token || !cookies.userId || !cookies.name || !cookies.displayName || !cookies.email || !cookies.icon || !cookies.status) {
        res.status(401).send({
            error: draftError.ServerError,
            drafts: [],
        })
    }

    /* get my drafts */
    const {error, drafts} = await blogApi.findDraftsByUserId(cookies.token, cookies.userId, 0, 9)

    /* response */
    return res.send({ 
        error: error,
        drafts: drafts,
    })
}

export default handler