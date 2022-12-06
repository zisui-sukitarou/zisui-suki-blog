import { baseUrl } from "../../../../../consts/config"
import { draftError } from "../../../../../domain/draft/error"
import { nextApiConfig } from "../../../config"

export const findDraftsByUserId = async (): Promise<{
    error: number,
    drafts: Array<DraftOverview>
}> => {
    const response = await fetch(baseUrl + nextApiConfig.findDraftByUserId.url, {
        method: nextApiConfig.findDraftByUserId.method,
        headers: nextApiConfig.findDraftByUserId.headers,
    })

    /* handle with server error */
    if (!response.ok) {
        return {
            error: draftError.ServerError,
            drafts: [],
        }
    }

    /* response */
    const jsonData = await response.json()
    return {
        error: jsonData.error,
        drafts: jsonData.drafts,
    }
}