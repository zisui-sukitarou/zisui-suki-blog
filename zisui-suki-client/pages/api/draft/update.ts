import type { NextApiRequest as Req, NextApiResponse as Res } from "next"
import { detectContentType } from "next/dist/server/image-optimizer";
import { parseCookies } from "nookies"
import { title } from "process";
import { stringify } from "querystring";
import { draftError } from "../../../domain/draft/error";
import * as blogApi from "../../../service/blogApi/draft/update"

const handler = async (req: Req, res: Res<{
    error: number,
}>) => {
    /* "POST"以外は、"404 Not Found"を返す */
    if (req.method !== "POST") {
        return res.status(404).send({
            error: draftError.ServerError,
        });
    }

    /* extract token & draft info */
    const cookies = parseCookies({ req })
    const reqParams = JSON.parse(req.body)
    if (!cookies.token || !cookies.userId || !cookies.name || !cookies.displayName || !cookies.email || !cookies.icon || !cookies.status) {
        res.status(401).send({
            error: draftError.ServerError,
        })
    }

    console.log("req params:", reqParams)

    /* get my drafts */
    const { error } = await blogApi.updateDraft(cookies.token, cookies.userId, reqParams)

    console.log("error:", error)

    /* response */
    return res.send({ 
        error: error,
    })
}

export default handler