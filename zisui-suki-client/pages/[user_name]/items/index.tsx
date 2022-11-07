import { NextPage } from "next"
import { useRouter } from "next/router"

const UserItemPage: NextPage = () => {
    const router = useRouter()
    const { user_name } = router.query

    return (
        <div>
            <h1>User Item Page</h1>
            <h2>{user_name}</h2>
        </div>
    )
}

export default UserItemPage