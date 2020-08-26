export const apiUrl = "http://localhost:8080/"

//TODO: Fetch token from session
let Token = "" // "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoyMzQzNzI1MzkxMjkxNjE4MzA1LCJjb21wYW55IjoiRHVjayBJbmMuIn0sImV4cCI6MTU5ODU1MzMzMSwiaWF0IjoxNTk4NDY2OTMxLCJpc3MiOiJnb1RlbXAudXNlcnNydiJ9.7V66xm-TF1Sy13UMZUtxcnxy8MuO9by7LPbeS_C_xc8"
export function httpGet(path, myFetch, myToken) {
    Token = myToken
    return req(path, 'GET', null, myFetch)
}

export function httpPost(path, data, myToken) {
    Token = myToken
    return req(path, 'POST', data)
}

export function httpPut(path,data, myToken) {
    Token = myToken
    return req(path, 'PUT', data)
}

export function httpDelete(path,data, myToken) {
    Token = myToken
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
    // console.log(res.status)
    // console.log(json)
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