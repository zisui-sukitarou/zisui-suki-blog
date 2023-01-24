import { NextPage } from "next";
import Header from "../../organisms/header"
import DraftList from "../../organisms/draftList"

type Props = {
    drafts: Array<DraftOverview>
    currentUser: CurrentUser | null | undefined
}

const Page: NextPage<Props> = ({drafts, currentUser}: Props) => {
    return (
        <div>
            <Header currentUser={currentUser} usedPage={null}></Header>
            
            <div className="relative h-screen overflow-hidden bg-primary">
                <div className="flex items-start justify-between">
                    <div className="relative hidden h-screen lg:block w-96">
                        <div className="h-full bg-inherit">
                        </div>
                    </div>
                    <DraftList></DraftList>
                    <div className="relative hidden h-screen lg:block w-96">
                        <div className="h-full bg-inherit">
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Page