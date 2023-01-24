import { NextPage } from "next";
import { useEffect, useState } from "react";
import { findDraftsByUserId } from "../../service/nextApi/draft/find/by/userId";
import DraftCard from "../molecules/draftCard";

type Props = {
    drafts: Array<DraftOverview>
}

const Page: NextPage<Props> = () => {

    const [drafts, setDrafts] = useState<Array<DraftOverview>>([])

    useEffect( () => { 
        (async () => {
            const { error, drafts } = await findDraftsByUserId()
            setDrafts(drafts)
        })()
    }, [])

    return (
    <div className="h-screen px-4 pb-24 overflow-auto md:px-6c w-full p-12 bg-inherit">

        <div className="grid grid-cols-1 gap-12 md:grid-cols-2 xl:grid-cols-2">
            {drafts.map(d => <DraftCard draft={d}></DraftCard>)}
        </div>
    </div>
    )

}

export default Page