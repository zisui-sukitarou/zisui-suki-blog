import { NextPage } from "next"
import { useRouter } from "next/router"
import { useEffect, useState } from "react"
import { useCurrentUser } from "../../hooks/useCurrentUser"
import { useRequireLogin } from "../../hooks/useRequireLogin"
import { findDraftsByUserId } from "../../service/nextApi/draft/find/by/userId"
import { newDraft } from "../../service/nextApi/draft/new"

const DraftPage: NextPage = () => {

    useRequireLogin()
    const {currentUser} = useCurrentUser()
    const [drafts, setDrafts] = useState<Array<DraftOverview>>([])

    useEffect( () => { 
        (async () => {
            const { error, drafts } = await findDraftsByUserId()
            setDrafts(drafts)
        })()
    }, [])

    const Drafts = drafts.map( draft => {
        return (<div key={draft.draftId}>
            <h2>{draft.title}</h2>
            <h3>{"概要：" + draft.abstract}</h3>
        </div>)
    })
    
    return (
        <div>
            <h1>{currentUser?.displayName + "の下書き"}</h1>
            {Drafts}
            <button onClick={() => {newDraft()}}>新規作成</button>
        </div>
    )
}

export default DraftPage