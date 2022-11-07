import { CurrentUser } from "../types/user"

export const login = async (name: string, password: string): Promise<{
    status: number | null,
    userData: CurrentUser | null,
}> => { 
    const response = await fetch("/api/login", { 
        method: "POST",
        headers: {
            'Content-Type': 'application/json' // JSON形式のデータのヘッダー
        },      
        body: JSON.stringify({
            name: name,
            password: password,
        }),
    })
    const jsonData = await response.json()
    return {
        status: jsonData.status,
        userData: jsonData.userData,
    }
}