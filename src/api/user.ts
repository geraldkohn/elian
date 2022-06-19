import service from "../utils/request";

export function login(data:any){
    return service({
        method: 'post',
        url: '/api/patient/login',
        data
    })
}