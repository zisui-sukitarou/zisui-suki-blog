import { useCurrentUser } from "../../hooks/useCurrentUser"
import { useRequireLogin } from "../../hooks/useRequireLogin"

const DraftNewPage = () => {

    useRequireLogin()
    const {currentUser} = useCurrentUser()
    
    return (
        <div>
            <h1>新規作成</h1>
            <h2>{currentUser?.name}</h2>
        </div>
    )
}

export default DraftNewPage