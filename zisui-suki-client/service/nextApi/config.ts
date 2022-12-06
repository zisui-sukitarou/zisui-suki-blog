export const nextApiConfig = {
    login: {
        url: "/api/login",
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
    },
    findDraftById: {
        url: "/api/draft/find/by/id",
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        },
    },
    findDraftByUserId: {
        url: "/api/draft/find/by/userId",
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        },
    },
    newDraft: {
        url: "/api/draft/new",
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        }
    },
    updateDraft: {
        url: "/api/draft/update",
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        }
    }
}