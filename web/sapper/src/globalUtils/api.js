export const apiUrl = "http://localhost:8080/"

//TODO: Fetch token from session
let Token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoyMzQzNzI1MzkxMjkxNjE4MzA1LCJjb21wYW55IjoiRHVjayBJbmMuIn0sImV4cCI6MTU5ODAzMTEyMywiaWF0IjoxNTk3OTQ0NzIzLCJpc3MiOiJnb1RlbXAudXNlcnNydiJ9.krz6ME_pHkcBIiaB_oFSla8ypnr-LsnmV-WeDmUlyyk"
export function httpGet(path, myFetch) {
    return req(path, 'GET', null, myFetch)
}

export function httpPost(path, data) {
    return req(path, 'POST', data)
}

export function httpPut(path,data) {
    return req(path, 'PUT', data)
}

export function httpDelete(path,data) {
    return req(path, 'DELETE', data)
}

async function req(path, method, data, myFetch) {
    let header = getHeader()
    let fetchIt
    if (myFetch){
        fetchIt = myFetch
    } else {
         fetchIt = fetch
    }
    const res = await fetchIt(apiUrl + path, {
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