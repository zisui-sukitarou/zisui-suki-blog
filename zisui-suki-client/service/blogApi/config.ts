export const blogApiConfig = {
    findBlogByTag: {
        url: "http://localhost:3002/find/blog/by/tag",
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        },
    },
    findBlogByUserId: {
        url: "http://localhost:3002/find/blog/by/userId",
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        },
    },
    findBlogByUserName: {
        url: "http://localhost:3002/find/blog/by/userName",
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        },
    },

    findDraftById: {
        url: "http://localhost:3002/find/draft/by/id",
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        },
    },
    findDraftByUserId: {
        url: "http://localhost:3002/find/draft/by/userId",
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        },
    },
    newDraft: {
        url: "http://localhost:3002/new/draft",
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
    },
    updateDraft: {
        url: "http://localhost:3002/update/draft",
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
    },
}