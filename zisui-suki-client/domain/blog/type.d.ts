type Blog = {
    blogId: string
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

type BlogOverview = {
    blogId: string
    title: string
    abstract: string
    evaluation: number
    status: number
    writer: User
    tags: Array<Tag>
    createdAt: string
    updatedAt: string
}