import { draftError } from "../../../domain/draft/error"
import { blogApiConfig } from "../config"

export const updateDraft = async (token: string, userId: string, draft: {
    draftId: string,
    title: string,
    abstract: string,
    content: string,
    evaluation: string,
    tags: Array<string>,
}): Promise<{
    error: number
}> => {

    const response = await fetch(blogApiConfig.updateDraft.url, {
        method: blogApiConfig.updateDraft.method,
        headers: {
            "Authorization": "Bearer " + token,
            ...blogApiConfig.updateDraft.headers,
        },
        body: JSON.stringify({
            draftId: draft.draftId,
            userId: userId,
            content: draft.content,
            abstract: draft.abstract,
            title: draft.title,
            evaluation: draft.evaluation,
            tags: draft.tags,
        })
    })

    /* handle with server error */
    if (!response.ok) {
        return {
            error: draftError.ServerError,
        }
    }

    /* return OK */
    return {
        error: draftError.OK,
    }
}