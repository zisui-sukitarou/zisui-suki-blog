export const blogApiConfig = {
    findByTag: {
        url: "http://localhost:3002/find/by/tag",
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        },
    },
    findByUserId: {
        url: "http://localhost:3002/find/by/userId",
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        },
    },
    findByUserName: {
        url: "http://localhost:3002/find/by/userName",
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        },
    },
}