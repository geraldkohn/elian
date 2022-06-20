import service from "../utils/request";

export function staffCreateRecord(data:any){
    return service({
        method: 'post',
        url: '/api/record/staff',
        data
    })
}

export function staffQueryRecord(patientIdcard:number){
    return service({
        method: 'get',
        url: '/api/record/staff',
        params: patientIdcard
    })
}

export function staffSetRWPermission(data:any){
    return service({
        method: 'post',
        url: '/api/record/staff-set-rw-permission',
        data
    })
}

export function staffSetRPermission(data:any){
    return service({
        method: 'post',
        url: '/api/record/staff-set-r-permission',
        data
    })
}

export function staffUpdatePhoto(data:any){
    return service({
        method: 'post',
        url: '/api/record/staff-update-photo',
        data
    })
}

export function staffUpdateDocument(data:any){
    return service({
        method: 'post',
        url: '/api/record/staff-update-document',
        data
    })
}

export function staffUpdateDescription(data:any){
    return service({
        method: 'post',
        url: '/api/record/staff-update-descrtiption',
        data
    })
}

export function patientQuery(patientIdcard:string){
    return service({
        method: 'get',
        url: '/api/record/patient',
        params: patientIdcard
    })
}

export function agencyQuery(patientIdcard:string){
    return service({
        method: 'get',
        url: '/api/record/agency',
        params: patientIdcard
    })
}

export function agencySetRWPermission(data:any){
    return service({
        method: 'post',
        url: '/api/record/agency-set-rw-permission',
        data
    })
}

export function agencySetRPermission(data:any){
    return service({
        method: 'post',
        url: '/api/record/agency-set-r-permission',
        data
    })
}