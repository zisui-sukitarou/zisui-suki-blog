import { NextPage } from "next"
import Router, { useRouter } from "next/router"
import { useEffect, useState } from "react"
import { useCurrentUser } from "../../../hooks/useCurrentUser"
import { useRequireLogin } from "../../../hooks/useRequireLogin"
import { findDraftById } from "../../../service/nextApi/draft/find/by/id"
// import SimpleMde from "react-simplemde-editor";
import dynamic from "next/dynamic";
const SimpleMde = dynamic(() => import("react-simplemde-editor"), { ssr: false });
import "easymde/dist/easymde.min.css";

const DraftEditPage: NextPage = () => {

    useRequireLogin()
    const {currentUser} = useCurrentUser()
    const [draft, setDraft] = useState<Draft | null>()

    const router = useRouter()
    const { draftId } = router.query

    useEffect(() => { 
        (async () => {
            if(typeof draftId !== "string"){ return }
            const { error, draft } = await findDraftById(draftId)
            setDraft(draft)
        })()
    }, [draftId])

    const onChange = (content: string) => {
        if(!draft) { return; }
        const nDraft: Draft = {
            ...draft,
            content: content,
        }
        setDraft(nDraft)
    }

    if (!draft) {
        return <div></div>
    }
    
    return (
        <div>
            <h1>{currentUser?.displayName + "の下書き"}</h1>
            <h2>タイトル</h2>
            <textarea value={draft.title}></textarea>
            <h2>概要</h2>
            <textarea value={draft.abstract}></textarea>
            <h2>本文</h2>
            <SimpleMde value={draft.content} onChange={onChange} />
        </div>
    )
}

export default DraftEditPage