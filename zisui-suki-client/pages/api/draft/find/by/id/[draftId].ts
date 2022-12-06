import type { NextApiRequest as Req, NextApiResponse as Res } from "next"
import { parseCookies } from "nookies"
import { draftError } from "../../../../../../domain/draft/error";
import * as blogApi from "../../../../../../service/blogApi/draft/find/by/id"

const handler = async (req: Req, res: Res<{
    error: number, 
    draft: Draft | null,
}>) => {
    /* "GET"以外は、"404 Not Found"を返す */
    if (req.method !== "GET") {
        return res.status(404).send({
            error: draftError.ServerError,
            draft: null,
        });
    }

    /* get draftId from req */
    const { draftId } = req.query
    if (typeof draftId !== "string") {
        return res.status(404).send({
            error: draftError.ServerError,
            draft: null,
        });
    }

    /* extract token */
    const cookies = parseCookies({ req })
    if (!cookies.token || !cookies.userId || !cookies.name || !cookies.displayName || !cookies.email || !cookies.icon || !cookies.status) {
        res.status(401).send({
            error: draftError.ServerError,
            draft: null,
        })
    }

    /* get my drafts */
    const {error, draft} = await blogApi.findDraftById(cookies.token, cookies.userId, draftId)

    /* response */
    return res.send({ 
        error: error,
        draft: draft,
    })
}

export default handler