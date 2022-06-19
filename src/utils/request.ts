import axios from 'axios'

export const getCookie = (cname: string) => {
    var name = cname + "=";
    var ca = document.cookie.split(';');
    for(var i=0; i<ca.length; i++)
    {
        var c = ca[i].trim();
        if (c.indexOf(name)===0) return c.substring(name.length,c.length);
    }
    return "";
}

const service = axios.create({
    baseURL: "http://localhost:8080",
    timeout: 10000
  })

export default service

export const setCookie = (name: string, value: string) => {
    let cookie = name + '=' + value + ';'
    let d = new Date()
    d.setTime(d.getTime() + (30 * 24 * 60 * 60 * 1000))
    let expires = 'expires=' + d.toUTCString()
    document.cookie = cookie + ' ' + expires
}

export const delCookie = (name: string) => {
    document.cookie= name + "=;expires=Thu, 01 Jan 1970 01:00:00 GMT;";
}
