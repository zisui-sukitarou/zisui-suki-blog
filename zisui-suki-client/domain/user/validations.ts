import { userError } from "./errors"


export const validateUserName = (name: string): number => {
    if (name.length < 6 || name.length > 15) {
        return userError.InvalidUserName
    }
    const regex = new RegExp(/^[a-zA-Z0-9_-]+$/)
    if (!regex.test(name)){
        return userError.InvalidUserName
    }
    return userError.OK
}

export const validateUserPassword = (password: string): number => {
    if (password.length < 6 || password.length > 15) {
        return userError.InvalidPassword
    }
    const regex = new RegExp(/^[a-zA-Z0-9!?_@-]+$/)
    if (!regex.test(password)){
        return userError.InvalidUserName
    }
    return userError.OK
}