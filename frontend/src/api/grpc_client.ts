import { Empty, Date as ProtoDate, GetCountriesDataRequest, CountriesData, GetAllDataResponse } from "../protos/covid_data_pb";
import { CovidDataClient } from "../protos/Covid_dataServiceClientPb";

const grpcClient = new CovidDataClient('http://localhost:8080');

export const getAllData = (callback: (data: CountriesData[], firstDate: Date, lastDate: Date) => void) => {
  var request = new Empty()
  grpcClient.getAllData(request, {}, (err, response) => {
    if (response == null) {
      console.log(err)
    } else {
      callback(
        response.getDataList()!,
        dateFromProto(response.getFirstDate()!),
        dateFromProto(response.getLastDate()!)
      )
    }
  })
}

export const getCountriesData = (date: Date, callback: (countriesData?: CountriesData) => void) => {
  var protoDate = new ProtoDate()
  protoDate.setYear(date.getFullYear())
  protoDate.setMonth(date.getMonth() + 1)
  protoDate.setDay(date.getDate())
  var request = new GetCountriesDataRequest()
  request.setDate(protoDate)
  grpcClient.getCountriesData(request, {}, (err, response) => {
    if (response == null) {
      console.log(err)
    } else {
      callback(response.getCountriesData()!)
    }
  })
}

export const getLastDate = (callback: (lastDate: Date) => void) => {
  var request = new Empty()
  grpcClient.getMostRecentDate(request, {}, (err, response) => {
    if (response == null) {
      console.log(err)
    } else {
      callback(dateFromProto(response))
    }
  })
}

const dateFromProto = (protoDate: ProtoDate): Date => {
  return new Date(Date.UTC(protoDate.getYear(), protoDate.getMonth() - 1, protoDate.getDay()))
}