import { GetServerSideProps, NextPage } from "next"
import { useRouter } from "next/router"
import { userError } from "../../domain/user/errors"
import { validateUserName } from "../../domain/user/validations"
import { useCurrentUser } from "../../hooks/useCurrentUser"
import { findBlogsByUserName } from "../../service/blogApi/blog/find/by/userName"
import Template from "../../components/templates/user_name/index"

type Props = {
    blogs: Array<BlogOverview>
}

const UserItemPage = ({ blogs }: Props) => {

    const {currentUser} = useCurrentUser()

    return (
        <Template blogs={blogs} currentUser={currentUser}/>
    )
}

export default UserItemPage

export const getServerSideProps: GetServerSideProps = async (context): Promise<{props: Props} | {redirect: Redirect}> => {
    const { user_name, page } = context.query
    /* calc begin & end (= begin + 10) of blogs */
    let begin: number, end: number
    if (!page) {
        begin = 0
        end = 9
    } else if (!Number(page)){
        return {
            redirect: {
                permanent: false,
                destination: `/${user_name}/items`
            }
        }
    } else {
        begin = (Number(page) - 1) * 10
        end = Number(page) * 10 - 1
    }

    if (typeof user_name !== "string") {
        return {
            props: {
                blogs: []
            }
        } 
    }

    /* user name validation */
    const err = validateUserName(user_name)
    if (err !== userError.OK) {
        return {
            props: {
                blogs: []
            }
        } 
    }

    /* fetch data from zisui-suki-blog-api */
    const {error, blogs} = await findBlogsByUserName(user_name, begin, end)
    console.log("blogs", blogs)
    if (error !== userError.OK) {
        return {
            props: {
                blogs: []
            }
        }
    }
    return {
        props: {
            blogs: blogs
        }
    }
}