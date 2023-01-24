import { NextPage } from "next";
import Icon from "./sIcon";

type Props = {
    createdAt: string
    updatedAt: string
}

const Page: NextPage<Props> = ({createdAt, updatedAt}: Props) => {
    if (createdAt === updatedAt) {
        return (
            <p className="text-gray-400">
                作成日：{createdAt.substring(0, 10)}
            </p>
        )
    } else {
        return (
            <p className="text-gray-400">
                作成日：{createdAt.substring(0, 10)}, 更新日：{updatedAt.substring(0, 10)}
            </p>
        )
    }
}

export default Page