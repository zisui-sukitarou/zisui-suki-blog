import { draftError } from "../../../../../domain/draft/error"
import { blogApiConfig } from "../../../config"

export const findDraftById = async (token: string, userId: string, draftId: string): Promise<{
    error: number,
    draft: Draft | null
}> => {

    const query = `?draftId=${draftId}`
    const response = await fetch(blogApiConfig.findDraftById.url + query, {
        method: blogApiConfig.findDraftById.method,
        headers: {
            "Authorization": "Bearer " + token,
            ...blogApiConfig.findDraftById.headers,
        }
    })

    /* handle with server error */
    if (!response.ok) {
        return {
            error: draftError.ServerError,
            draft: null,
        }
    }

    /* return blog data */
    const jsonData = await response.json()
    return {
        error: draftError.OK,
        draft: jsonData,
    }
}