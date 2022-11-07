import { NextPage } from "next"
import { useRouter } from "next/router"
import { useRecoilValue } from "recoil"
import { currentUserState } from "../../stores/currentUser"

const UserPage: NextPage = () => {
    const router = useRouter()
    const { user_name } = router.query
    const currentUser = useRecoilValue(currentUserState)

    const PrivateInfo = () => {
        if (user_name === currentUser?.name) {
            return <div>
                <h3>メール：{currentUser?.email}</h3>
            </div>
        }
        return null
    }

    return (
        <div>
            <h1>User Page</h1>
            <h2>{user_name}</h2>
            <PrivateInfo />
        </div>
    )
}

export default UserPage