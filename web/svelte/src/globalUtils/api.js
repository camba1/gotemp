export const apiUrl = "http://localhost:8080/"

//TODO: Fetch token from session
let Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoyMzQzNzI1MzkxMjkxNjE4MzA1LCJjb21wYW55IjoiRHVjayBJbmMuIn0sImV4cCI6MTU5NDkyODU5NiwiaWF0IjoxNTk0ODQyMTk2LCJpc3MiOiJnb1RlbXAudXNlcnNydiJ9.VhkC0xXndWxcz9B3VPrcXVOCw9FOb8j7AOMgOfqtnuM"

export function httpGet(path) {
    return req(path, 'GET')
}

export function httpPost(path, data) {
    return req(path, 'POST', data)
}

export function httpPut(path,data) {
    return req(path, 'PUT', data)
}

async function req(path, method, data) {
    let header = getHeader()
    const res = await fetch(apiUrl + path, {
        method,
        headers: header,
        body: data && JSON.stringify(data)
    })
    let json = await res.json()
    console.log(res.status)
    return {ok: res.ok, data: json}
}

function getHeader(){
    if (Token === "") {
        return { 'Content-Type': 'application/json'};
    } else {
        return {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + Token,
        }
    }
}