import { NextPage } from "next";
import { useEffect, useState } from "react";
import WriterInfo from "../molecules/writerInfo";
import TagList from "../molecules/tagList";
import TimeInfo from "../atoms/timeInfo"

type Props = {
    blog: BlogOverview

}

const Page: NextPage<Props> = ({blog}: Props) => {

    const [tags, setTags] = useState(Array<Tag>)
    useEffect(() => {
        setTags(blog.tags)
    })

    return (
        <div className="m-auto overflow-hidden rounded-lg shadow-lg cursor-pointer h-90 w-full">
            <a href={"./items/"+blog.blogId} className="block w-full h-full">
                <div className="w-full px-8 pb-2 pt-4 bg-white ">
                    <div className="mb-2 text-xl text-gray-800 font-bold">
                        {blog.title}
                    </div>
                    <div className="font-light text-gray-400 text-md">
                        {blog.abstract}
                    </div>
                    <TimeInfo createdAt={blog.createdAt} updatedAt={blog.updatedAt} />
                    <TagList tags={blog.tags} />
                    <WriterInfo writer={blog.writer}/>
                </div>
                
            </a>
        </div>
    )
}

export default Page