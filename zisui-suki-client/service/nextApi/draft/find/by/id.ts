import { baseUrl } from "../../../../../consts/config"
import { draftError } from "../../../../../domain/draft/error"
import { nextApiConfig } from "../../../config"

export const findDraftById = async (draftId: string): Promise<{
    error: number,
    draft: Draft | null
}> => {
    const query = `/${draftId}`
    const response = await fetch(baseUrl + nextApiConfig.findDraftById.url + query, {
        method: nextApiConfig.findDraftById.method,
        headers: nextApiConfig.findDraftById.headers,
    })

    /* handle with server error */
    if (!response.ok) {
        return {
            error: draftError.ServerError,
            draft: null,
        }
    }

    /* response */
    const jsonData = await response.json()
    return {
        error: jsonData.error,
        draft: jsonData.draft,
    }
}