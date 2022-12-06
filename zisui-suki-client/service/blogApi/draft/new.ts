import { draftError } from "../../../domain/draft/error"
import { blogApiConfig } from "../config"

export const newDraft = async (token: string, userId: string): Promise<{
    error: number,
    draftId: string | null,
}> => {

    const response = await fetch(blogApiConfig.newDraft.url, {
        method: blogApiConfig.newDraft.method,
        headers: {
            "Authorization": "Bearer " + token,
            ...blogApiConfig.newDraft.headers,
        },
        body: JSON.stringify({
            userId: userId,
        })
    })

    /* handle with server error */
    if (!response.ok) {
        return {
            error: draftError.ServerError,
            draftId: null,
        }
    }

    /* return blog data */
    const jsonData = await response.json()
    return {
        error: draftError.OK,
        draftId: jsonData.draftId,
    }
}