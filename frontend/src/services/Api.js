function token() {
    return window.localStorage.getItem('token') ? window.localStorage.getItem('token') : ''
}

export async function ApiQuery(endpoint, params = {}) {
    const newParams = {
        ...params,
        withCredentials: true,    
        crossorigin: true,    
        credentials: 'include',
        headers: {
            ...params.headers, 
            'Cookie': 'token=' + token()
        }
    }
    const response = await fetch('http://localhost:8080/api/v1' + endpoint, newParams);
    if (!response.ok) {
        return {
            statusCode: response.status,
            status: "failed",
            message: response.statusText
        }
    }

    const responseData = await response.json();

    if (responseData && !responseData.hasOwnProperty('data')) {
        return {
            statusCode: response.status,
            status: "failed",
            message: response.statusText
        };
    }

    if (responseData.data && responseData.data.hasOwnProperty('token')) {
        window.localStorage.setItem('token', responseData.data.token);
        console.log("set token" + responseData.data.token);
    }
    return {
        statusCode: response.status,
        status: responseData.status,
        data: responseData.data,
        message: responseData.message
    };
}

export async function TokenCheck() {
    const endpoint = "/check"
    const params = {
        method: 'post',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            "token": token(),
        })
    }
    const response = await fetch('http://localhost:8080/api/v1' + endpoint, params);

    if (!response.ok || response.status != 200) {
        var total = "" + response.status;
        console.log(total);
        return false;
    }
    
    const responseData = await response.json();

    if (responseData.status == "failed" || !responseData.data) {
        var total = "" + responseData.status + ", " + responseData.message;
        console.log(total);
        return false;
    }

    if (!responseData.data.exists) {
        return false;
    }

    return true;
}

export async function HandleDifferentResponses(response) {

}