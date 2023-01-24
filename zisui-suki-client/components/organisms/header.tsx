import "tailwindcss/tailwind.css"
import Header from "next/headers"
import { NextPage } from "next"

type Props = {
    currentUser: CurrentUser | null | undefined
    usedPage: string | null
}

const Page: NextPage<Props> = ({currentUser, usedPage}: Props) => {

    const home = usedPage === "home" ? (
        <a className="font-bold text-gray-300 hover:text-orange-300 hover:text-gray-800px-3 py-2 rounded-md text-md" href="/#">
            ホーム
        </a>
    ) : (
        <a className="font-bold text-gray-300 hover:text-orange-300 hover:text-gray-800px-3 py-2 rounded-md text-md" href="/#">
            ホーム
        </a>
    )

    const content = usedPage === "content" ? (
        <a className="font-bold text-gray-300  hover:text-orange-300 px-3 py-2 rounded-md text-md" href="/#">
            記事
        </a>
    ) : (
        <a className="font-bold text-gray-300  hover:text-orange-300 px-3 py-2 rounded-md text-md" href="/#">
            記事
        </a>
    )

    const itemsHorizontal = (
        <div className="hidden md:block">
            <div className="flex items-baseline ml-10 space-x-4">
                {home}
                {content}
            </div>
        </div>
    )

    const itemsVertical = (
        <div className="md:hidden">
            <div className="px-2 pt-2 pb-3 space-y-1 sm:px-3">
                {home}
                {content}
            </div>
        </div>
    )

    return ( 
        <nav className="bg-white py-4 bg-clip-padding shadow-md z-20 relative">
            <div className="px-8 mx-auto max-w-7xl">
                <div className="flex items-center justify-between h-16">
                    <div className=" flex items-center">
                        <a className="flex-shrink-0" href="/">
                            <img className="w-auto h-16" src="/zisui-suki-blog-logo.png" alt="Workflow"/>
                        </a>
                        {itemsHorizontal}
                    </div>
                </div>
            </div>
            {itemsVertical}
        </nav>
    )
}

export default Page