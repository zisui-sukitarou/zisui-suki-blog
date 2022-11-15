import { GetServerSideProps, NextPage } from "next"
import { useRouter } from "next/router"
import { userError } from "../../../domain/user/errors"
import { validateUserName } from "../../../domain/user/validations"
import { findBlogsByUserName } from "../../../service/blogApi/find/by/userName"

type Props = {
    blogs: Array<BlogOverview>
}

const UserItemPage = ({ blogs }: Props) => {
    const Blogs = blogs.map((blog) => {
        return (<div key={blog.blogId}>
            <h2>{blog.title}</h2>
            <h3>{"written by " + blog.writer.name}</h3>
            <h3>{blog.abstract}</h3>
        </div>)
    })

    return (
        <div>
            <h1>User Item Page</h1>
            {Blogs}
        </div>
    )
}

export default UserItemPage

type Redirect =
| {
    statusCode: 301 | 302 | 303 | 307 | 308; // ステータスコード
    destination: string; // リダイレクト先のURL
    basePath?: false; // `basePath`を無効にします
  }
| {
    permanent: boolean; // 永続的なリダイレクト化のフラグ
    destination: string; // 同上
    basePath?: false; // 同上
  };

export const getServerSideProps: GetServerSideProps = async (context): Promise<{props: Props} | {redirect: Redirect}> => {
    const { user_name, page } = context.query
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