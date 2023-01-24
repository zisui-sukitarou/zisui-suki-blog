import { NextPage } from "next";
import Tag from "../atoms/tag"

type Props = {
    tags: Array<Tag>
}

const Page: NextPage<Props> = ({tags}: Props) => {

    return (
        <div className="flex flex-wrap items-center mt-4 justify-starts">
            {tags.map(t => (
                <Tag tag={t} />
            ))}
        </div>
    )
}

export default Page