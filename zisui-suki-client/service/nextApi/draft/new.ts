import { baseUrl } from "../../../consts/config"
import { draftError } from "../../../domain/draft/error"
import { nextApiConfig } from "../config"

export const newDraft = async(): Promise<{
    error: number,
    draftId: string | null,
}> => {
    const response = await fetch(baseUrl + nextApiConfig.newDraft.url, {
        method: nextApiConfig.newDraft.method,
        headers: nextApiConfig.newDraft.headers,
    })

    /* handle with server error */
    if (!response.ok) {
        return {
            error: draftError.ServerError,
            draftId: null,
        }
    }

    /* response */
    const jsonData = await response.json()
    return {
        error: jsonData.error,
        draftId: jsonData.draftId
    }
}