export const apiUrl = "http://localhost:8080/"

let Token = ""

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
    const res = await fetch(apiUrl + path, {
        method,
        headers: {
            'Content-Type': 'application/json',
            // 'Authorization': 'Bearer ' + Token,
        },
        body: data && JSON.stringify(data)
    })
    let json = await res.json()
    return {ok: res.ok, data: json}
}