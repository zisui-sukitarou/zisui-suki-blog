import { parseCookies, setCookie, destroyCookie } from 'nookies'
import { baseUrl } from '../consts/config'

export const fetchCurrentUser = async (): Promise<CurrentUser | null> => {
    const response = await fetch(baseUrl + "/api/user")
    if (response.status !== 200) {
        return null
    }

    const jsonData = await response.json()
    return {
        userId: jsonData.userId,
        name: jsonData.name,
        displayName: jsonData.displayName,
        email: jsonData.email,
        icon: jsonData.icon,
        status: jsonData.status,
    }
}