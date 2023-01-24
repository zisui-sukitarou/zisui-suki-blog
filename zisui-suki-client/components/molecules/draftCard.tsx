import { NextPage } from "next";

type Props = {
    draft: DraftOverview
}

const Page: NextPage<Props> = ({draft}: Props) => {

    const abstract = draft.abstract.length > 100 ? (
        draft.abstract.substring(0, 100) + "..."
    ) : (
        draft.abstract
    )

    return (
        <div className="m-auto overflow-hidden rounded-lg shadow-lg cursor-pointer h-90 w-full">
            <a href="#" className="block w-full h-full">
                <div className="w-full p-4 bg-white ">
                    <p className="mb-2 text-xl font-medium text-gray-800">
                        {draft.title}
                    </p>
                    <p className="font-light text-gray-400 text-md">
                        {abstract}
                    </p>
                    <div className="flex flex-wrap items-center mt-4 justify-starts">
                        <div className="text-xs mr-2 py-1.5 px-4 text-gray-600 bg-orange-100 rounded-2xl">
                            #Car
                        </div>
                        <div className="text-xs mr-2 py-1.5 px-4 text-gray-600 bg-orange-100 rounded-2xl">
                            #Money
                        </div>
                    </div>
                </div>
                <div className="flex items-center mt-4">
                    <a href="#" className="relative block">
                        <img alt="profil" src="/images/person/6.jpg" className="mx-auto object-cover rounded-full h-10 w-10 "/>
                    </a>
                    <div className="flex flex-col justify-between ml-4 text-sm">
                        <p className="text-gray-800">
                            @{draft.writer.name}
                        </p>
                        <p className="text-gray-400">
                            20 mars 2029 - 6 min read
                        </p>
                    </div>
                </div>
            </a>
        </div>
    )
}

export default Page