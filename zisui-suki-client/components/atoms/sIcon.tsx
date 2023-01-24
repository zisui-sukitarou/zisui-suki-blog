import { NextPage } from "next"

type Props = {
    iconPath: string
}

const Page: NextPage<Props> = ({iconPath}: Props) => {

    if (!iconPath || iconPath === "" || iconPath === "yyy.png") {
        return (
            <a href="#" className="relative block">
                <img alt="profil" src="/kkrn_icon_user_5.png" className="mx-auto object-cover rounded-full h-10 w-10"/>
            </a>
        )
    } else {
        return (
            <a href="#" className="relative block">
                <img alt="profil" src={iconPath} className="mx-auto object-cover rounded-full h-10 w-10"/>
            </a>
        )
    }
}

export default Page