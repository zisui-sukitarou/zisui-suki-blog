import { NextPage } from "next"
import Header from "../../organisms/header"
import BlogList from "../../organisms/blogList"
import Icon from "../../atoms/lIcon"

type Props = {
    blogs: Array<BlogOverview>
    currentUser: CurrentUser | null | undefined
}

const Page: NextPage<Props> = ({blogs, currentUser}: Props) => {

    // TODO:  icon path 指定

    return <div>
        <Header currentUser={currentUser} usedPage={null}></Header>
        
        <div className="relative h-screen overflow-auto bg-primary pt-10">
            <div className="lg:flex lg:items-start justify-between items-center mx-32 gap-10">
                <div className="relative bg-white lg:block lg:w-[400px] lg:min-w-[400px] rounded-lg pt-4">
                    <Icon iconPath="" /> 
                    <div className="mx-auto justify-center text-center font-semibold text-lg pb-4">
                        @test
                    </div>
                    <div className="text-center text-sm mx-4 pb-4">
                        ああああああああああああああああああああああああああああああああ
                    </div>
                </div>
                <div className="relative w-full">
                    <BlogList blogs={blogs}></BlogList>
                </div>
            </div>
        </div>
    </div>
}

export default Page