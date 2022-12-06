import { blogError } from "../../../../domain/blog/error"
import { blogApiConfig } from "../../config"

export const findBlogsByUserId = async (userId: string, begin: number, end: number): Promise<{
    error: number,
    blogs: Array<BlogOverview>
}> => {

    const query = `?userId=${userId}&begin=${begin}&end=${end}`
    const response = await fetch(blogApiConfig.findDraftByUserId.url + query, {
        method: blogApiConfig.findDraftByUserId.method,
        headers: blogApiConfig.findDraftByUserId.headers,
    })

    /* handle with server error */
    if (!response.ok) {
        return {
            error: blogError.ServerError,
            blogs: []
        }
    }

    /* return blog data */
    const jsonData = await response.json()
    return {
        error: blogError.OK,
        blogs: jsonData,
    }
}