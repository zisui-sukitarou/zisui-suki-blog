import { GetServerSideProps, NextPage } from "next"
import { useRouter } from "next/router"
import { userError } from "../../../domain/user/errors"
import { useCurrentUser } from "../../../hooks/useCurrentUser"
import { findBlogById } from "../../../service/blogApi/blog/find/by/id"
import Template from "../../../components/templates/user_name/items/blogId"

type Props = {
    blog: Blog | null
}

const UserItemPage = ({ blog }: Props) => {

    const {currentUser} = useCurrentUser()
    if(!blog) {
        return <h1>404 NotFound</h1>
    }

    return (
        <Template blog={blog} currentUser={currentUser}/>
    )
}

export default UserItemPage

export const getServerSideProps: GetServerSideProps = async (context): Promise<{props: Props} | {redirect: Redirect}> => {
    const { blog_id } = context.query

    if (typeof blog_id !== "string") {
        return {
            props: {
                blog: null
            }
        } 
    }

    /* fetch data from zisui-suki-blog-api */
    const {error, blog} = await findBlogById(blog_id)
    console.log("blog", blog)
    if (error !== userError.OK) {
        return {
            props: {
                blog: null
            }
        }
    }
    return {
        props: {
            blog: blog
        }
    }
}