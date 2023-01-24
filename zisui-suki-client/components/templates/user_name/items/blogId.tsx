import { NextPage } from "next";
import Header from "../../../organisms/header";
import ReactMarkdown from "react-markdown";
import { useEffect, useState } from "react";
import TagList from "../../../molecules/tagList"
import WriterInfo from "../../../molecules/writerInfo"
import TimeInfo from "../../../atoms/timeInfo"


type Props = {
    blog: Blog
    currentUser: CurrentUser | null | undefined
}

const text = "# 初めに\n ## 前提知識 \n ### 前提の前提 \nああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああああ \n * ABC \n * DEF \n * ABC \n * DEF \n * ABC \n * DEF \n * ABC \n * DEF \n * ABC \n * DEF \n * ABC \n * DEF \n * ABC \n * DEF \n * ABC \n * DEF \n * ABC \n * DEF \n * ABC \n * DEF \n * ABC \n * DEF \n * ABC \n * DEF "

const Page: NextPage<Props> = ({blog, currentUser}: Props) => {

    return <div>
        <div className="relative h-screen overflow-auto bg-primary">
            <Header currentUser={currentUser} usedPage={null}></Header>
            <div className="flex items-start justify-between mt-10">
                <div className="relative hidden h-screen lg:block w-[1000px]">
                    <div className="h-full bg-inherit fixed">
                        
                    </div>
                </div>
                <div className="md:px-6c w-full min-w-[800px] p-10 mb-10 bg-white rounded-xl">
                    <div className="pb-16 mx-auto max-w-xl text-md">
                        <div className="text-center text-3xl font-bold pb-4">
                            {blog.title}
                        </div>
                        <TimeInfo createdAt={blog.createdAt} updatedAt={blog.updatedAt} />
                        <WriterInfo writer={blog.writer} />
                        <TagList tags={blog.tags}/>
                    </div>
                    <div className="markdown">
                        <ReactMarkdown>
                            {text}
                        </ReactMarkdown>
                    </div>
                </div>
                <div className="relative hidden h-screen lg:block w-[1400px]">
                    <div className="h-full bg-inherit fixed">
                        かきく
                    </div>
                </div>
            </div>
        </div>
    </div>
}

export default Page