import { NextPage } from "next";
import { useEffect, useState } from "react";
import { findDraftsByUserId } from "../../service/nextApi/draft/find/by/userId";
import BlogCard from "../molecules/blogCard";

type Props = {
    blogs: Array<BlogOverview>
}

const Page: NextPage<Props> = ({blogs}: Props) => {

    const [blogList, setBlogList] = useState(Array<BlogOverview>)
    useEffect(()=>{
        setBlogList(blogs)
    }, [])

    return (
    <div className="h-screen overflow-auto w-full bg-inherit">
        <div className="grid grid-cols-1 gap-6 md:grid-cols-1 xl:grid-cols-1">
            {blogList.map(b => <BlogCard blog={b}></BlogCard>)}
        </div>
    </div>
    )

}

export default Page