import { blogError } from "../../../../../domain/blog/error"
import { blogApiConfig } from "../../../config"

export const findBlogById = async (blogId: string): Promise<{
    error: number,
    blog: Blog | null
}> => {

    const query = `?blogId=${blogId}`
    const response = await fetch(blogApiConfig.findBlogById.url + query, {
        method: blogApiConfig.findBlogById.method,
        headers: blogApiConfig.findBlogById.headers,
    })

    /* handle with server error */
    if (!response.ok) {
        return {
            error: blogError.ServerError,
            blog: null
        }
    }

    /* return blog data */
    const jsonData = await response.json()
    if (!jsonData) {
        return {
            error: blogError.OK,
            blog: null,
        }
    }
    return {
        error: blogError.OK,
        blog: jsonData,
    }
}