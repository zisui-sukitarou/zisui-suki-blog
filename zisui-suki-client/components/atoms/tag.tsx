import { NextPage } from "next"

type Props = {
    tag: Tag
}

const Page: NextPage<Props> = ({tag}: Props) => {

    if (!tag.icon || tag.icon === ""){
        return (
            <div className="text-xs mr-2 py-1.5 px-4 text-gray-600 bg-white border rounded-2xl">
                {"#" + tag.tagName}
            </div>
        )
    } else {
        return (
            <div className="text-xs mr-2 py-1.5 px-4 text-gray-600 bg-white border rounded-2xl">
                {"#" + tag.tagName}
            </div>
        )
    }
}

export default Page