export const apiUrl = "http://localhost:8080/"

//TODO: Fetch token from session
let Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoyMzQzNzI1MzkxMjkxNjE4MzA1LCJjb21wYW55IjoiRHVjayBJbmMuIn0sImV4cCI6MTU5Nzg2OTMzNSwiaWF0IjoxNTk3NzgyOTM1LCJpc3MiOiJnb1RlbXAudXNlcnNydiJ9.-n8W9LTCseSqdxNBYzicbWS__etYObKzLDt66_hvHZk"
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
    console.log(json)
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