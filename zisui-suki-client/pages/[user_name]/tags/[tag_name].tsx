import { NextPage } from "next"
import { useRouter } from "next/router"

const BlogPage: NextPage = () => {
    const router = useRouter()
    const { user_name, tag_name } = router.query

    return (
        <div>
            <h1>User Tag Items Page</h1>
            <h2>{user_name}</h2>
            <h2>{tag_name}</h2>
        </div>
    )
}

export default BlogPage