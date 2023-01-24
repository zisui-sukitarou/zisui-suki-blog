import { NextPage } from "next";
import Icon from "../atoms/sIcon";

type Props = {
    writer: User
}

const Page: NextPage<Props> = ({writer}: Props) => {
    return (
        <div className="flex items-center bg-inherit">
            <Icon iconPath={writer.icon}/>
            <div className="flex flex-col justify-between ml-1 text-gray-800">
                {"@" + writer.name}                
            </div>
        </div>
    )
}

export default Page