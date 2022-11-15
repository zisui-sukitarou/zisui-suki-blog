import { userError } from "../../domain/user/errors"
import { nextApiConfig } from "./config"

export const login = async (name: string, password: string): Promise<{
    error: number,
    userData: CurrentUser | null,
}> => {
    const response = await fetch(nextApiConfig.login.url, { 
        method: nextApiConfig.login.method,
        headers: nextApiConfig.login.headers,      
        body: JSON.stringify({
            name: name,
            password: password,
        }),
    })

    /* handle with server error */
    if (!response.ok) {
        return {
            error: userError.ServerError,
            userData: null,
        }
    }

    /* return user data */
    const jsonData = await response.json()
    return {
        error: jsonData.error,
        userData: jsonData.userData,
    }
}