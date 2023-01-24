type Draft = {
    draftId: string
    content: string
    title: string
    abstract: string
    evaluation: number
    status: number
    writer: User
    tags: Array<Tag>
    createdAt: string
    updatedAt: string
}

type DraftOverview = {
    draftId: string
    title: string
    abstract: string
    evaluation: number
    status: number
    writer: User
    tags: Array<Tag>
    createdAt: string
    updatedAt: string
}