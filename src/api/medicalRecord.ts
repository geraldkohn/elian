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

export function staffSetRW(data:any){
    return service({
        method: 'post',
        url: '',
    })
}