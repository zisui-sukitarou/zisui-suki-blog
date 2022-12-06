type Draft = {
    draftId: string
    content: string
    title: string
    abstract: string
    evaluation: number
    status: number
    writer: User
    tags: Array<Tag>
}

type DraftOverview = {
    draftId: string
    title: string
    abstract: string
    evaluation: number
    status: number
    writer: User
    tags: Array<Tag>
}