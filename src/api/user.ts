import service from "../utils/request";

export function patientLogin(data:any){
    return service({
        method: 'post',
        url: '/api/patient/login',
        data
    })
}

export function staffLogin(data:any){
    return service({
        method: 'post',
        url: '/api/staff/login',
        data
    })
}

export function agencyLogin(data:any){
    return service({
        method: 'post',
        url: '/api/agency/login',
        data
    })
}

export function patientRegister(data:any){
    return service({
        method: 'post',
        url: '/api/patient/register',
        data
    })
}

export function staffRegister(data:any){
    return service({
        method: 'post',
        url: '/api/staff/register',
        data
    })
}

export function agencyRegister(license:string){
    return service({
        method: 'post',
        url: '/api/agency/register',
        data: license
    })
}