import { draftError } from "../../../../../domain/draft/error"
import { blogApiConfig } from "../../../config"

export const findDraftsByUserId = async (token: string, userId: string, begin: number, end: number): Promise<{
    error: number,
    drafts: Array<DraftOverview>
}> => {

    const query = `?userId=${userId}&begin=${begin}&end=${end}`
    const response = await fetch(blogApiConfig.findDraftByUserId.url + query, {
        method: blogApiConfig.findDraftByUserId.method,
        headers: {
            "Authorization": "Bearer " + token,
            ...blogApiConfig.findDraftByUserId.headers,
        }
    })

    /* handle with server error */
    if (!response.ok) {
        return {
            error: draftError.ServerError,
            drafts: []
        }
    }

    /* return blog data */
    const jsonData = await response.json()
    return {
        error: draftError.OK,
        drafts: jsonData,
    }
}