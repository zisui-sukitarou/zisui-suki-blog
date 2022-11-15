import { CurrentUser } from "../../types/user"

type Props = {
    userData: CurrentUser
}

const Header = ({ userData }: Props) => {

    return (
    <div>
        <h1>ヘッダー</h1>
        <h2>これは自炊好きブログです</h2>
        <h3>ログイン中：{userData.displayName}</h3>
    </div>)
}

export default Header