import type { NextApiRequest as Req, NextApiResponse as Res } from "next"
import { parseCookies } from "nookies"
import { draftError } from "../../../domain/draft/error";
import * as blogApi from "../../../service/blogApi/draft/new"

const handler = async (req: Req, res: Res<{
    error: number, 
    draftId: string | null,
}>) => {
    /* "POST"以外は、"404 Not Found"を返す */
    if (req.method !== "POST") {
        return res.status(404).send({
            error: draftError.ServerError,
            draftId: null,
        });
    }

    /* extract token */
    const cookies = parseCookies({ req })
    if (!cookies.token || !cookies.userId || !cookies.name || !cookies.displayName || !cookies.email || !cookies.icon || !cookies.status) {
        res.status(401).send({
            error: draftError.ServerError,
            draftId: null,
        })
    }

    console.log("cookie:", cookies)

    /* get my drafts */
    const {error, draftId} = await blogApi.newDraft(cookies.token, cookies.userId)

    console.log("draftId:", draftId)

    /* response */
    return res.send({ 
        error: error,
        draftId: draftId,
    })
}

export default handler