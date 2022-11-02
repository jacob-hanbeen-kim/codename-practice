
const baseUrl = `${process.env.REACT_APP_API_URL}`;

const getGreet = async () => {
    console.log(baseUrl);
    return fetch(`${baseUrl}/greet`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        },
    }).then((res) => {
        if (res.ok) return res.json()
        return res.json().then(json => Promise.reject(json));
    }).then((res) => {
        return Promise.resolve(res);
    }).catch(e => {
        console.error(e)
    })
}

export const HelloWorldService = {
    getGreet
}