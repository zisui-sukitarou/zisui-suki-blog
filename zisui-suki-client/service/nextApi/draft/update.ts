import { baseUrl } from "../../../consts/config"
import { draftError } from "../../../domain/draft/error"
import { nextApiConfig } from "../config"

export const updateDraft = async(userId: string, draft: {
    draftId: string,
    title: string,
    abstract: string,
    content: string,
    evaluation: string,
    tags: Array<string>,
}): Promise<{
    error: number,
}> => {
    const response = await fetch(baseUrl + nextApiConfig.updateDraft.url, {
        method: nextApiConfig.updateDraft.method,
        headers: nextApiConfig.updateDraft.headers,
        body: JSON.stringify(draft)
    })

    /* handle with server error */
    if (!response.ok) {
        return {
            error: draftError.ServerError,
        }
    }

    /* response */
    const jsonData = await response.json()
    return {
        error: jsonData.error,
    }
}