type Blog = {
    blogId: string
    content: string
    title: string
    abstract: string
    evaluation: number
    status: number
    writer: User
    tags: Array<Tag>
}

type BlogOverview = {
    blogId: string
    title: string
    abstract: string
    evaluation: number
    status: number
    writer: User
    tags: Array<Tag>
}