import { NextPage } from "next"
import { useRouter } from "next/router"
import { useRecoilValue } from "recoil"
import { useCurrentUser } from "../../hooks/useCurrentUser"
import { useRequireLogin } from "../../hooks/useRequireLogin"
import { currentUserState } from "../../stores/currentUser"

const DraftPage: NextPage = () => {

    useRequireLogin()
    const {isAuthChecking, currentUser} = useCurrentUser()
    
    return (
        <div>
            <h1>Draft Page</h1>
            <div>{currentUser?.displayName}</div>
        </div>
    )
}

export default DraftPage