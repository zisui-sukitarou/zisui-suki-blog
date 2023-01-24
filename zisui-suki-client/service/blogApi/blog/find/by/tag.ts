import { blogError } from "../../../../../domain/blog/error"
import { blogApiConfig } from "../../../config"

export const findBlogsByTagName = async (userName: string, begin: number, end: number): Promise<{
    error: number,
    blogs: Array<BlogOverview>
}> => {

    const query = `?tagName=${userName}&begin=${begin}&end=${end}`
    const response = await fetch(blogApiConfig.findBlogsByTagName.url + query, {
        method: blogApiConfig.findBlogsByTagName.method,
        headers: blogApiConfig.findBlogsByTagName.headers,
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
    if (!jsonData) {
        return {
            error: blogError.OK,
            blogs: [],
        }
    }
    return {
        error: blogError.OK,
        blogs: jsonData,
    }
}