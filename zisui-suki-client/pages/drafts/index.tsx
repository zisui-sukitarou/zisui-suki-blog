import { NextPage } from "next"
import { useRouter } from "next/router"
import { useEffect, useState } from "react"
import { useCurrentUser } from "../../hooks/useCurrentUser"
import { useRequireLogin } from "../../hooks/useRequireLogin"
import { findDraftsByUserId } from "../../service/nextApi/draft/find/by/userId"
import { newDraft } from "../../service/nextApi/draft/new"
import Template from "../../components/templates/drafts/index"

const Page: NextPage = () => {

    useRequireLogin()
    const {currentUser} = useCurrentUser()
    const [drafts, setDrafts] = useState<Array<DraftOverview>>([])

    useEffect( () => { 
        (async () => {
            const { error, drafts } = await findDraftsByUserId()
            setDrafts(drafts)
        })()
    }, [])
    
    return <Template 
        drafts={drafts} 
        currentUser={currentUser}/>
}

export default Page