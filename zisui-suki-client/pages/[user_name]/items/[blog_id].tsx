import { NextPage } from "next"
import { useRouter } from "next/router"

const BlogPage: NextPage = () => {
    const router = useRouter()
    const { user_name, blog_id } = router.query

    return (
        <div>
            <h1>Blog Page</h1>
            <h2>{user_name}</h2>
            <h2>{blog_id}</h2>
        </div>
    )
}

export default BlogPage