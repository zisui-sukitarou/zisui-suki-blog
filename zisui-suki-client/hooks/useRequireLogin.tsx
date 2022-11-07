import { useEffect } from 'react';
import { useRouter } from 'next/router';
import { useCurrentUser } from "./useCurrentUser"

export const useRequireLogin = () => {
    const { isAuthChecking, currentUser } = useCurrentUser();
    const router = useRouter();

    console.log("current_user:", currentUser)
    console.log("is_auth_checking:", isAuthChecking)

    useEffect(() => {
        if (isAuthChecking) {
            return
        }
        if (!currentUser) { 
            router.push("/login")
        }
    },[isAuthChecking, currentUser])
}