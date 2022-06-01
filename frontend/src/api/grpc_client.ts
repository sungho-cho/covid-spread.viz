import { Empty, Date as ProtoDate, GetCountriesDataRequest, CountriesData } from "../protos/covid_data_pb";
import { CovidDataClient } from "../protos/Covid_dataServiceClientPb";

const grpcClient = new CovidDataClient('http://localhost:8080');

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
      callback(response.getCountriesData())
    }
  })
}

export const getLastDate = (callback: (lastDate: Date) => void) => {
  var request = new Empty()
  grpcClient.getMostRecentDate(request, {}, (err, response) => {
    if (response == null) {
      console.log(err)
    } else {
      callback(new Date(Date.UTC(response.getYear(), response.getMonth() - 1, response.getDay())))
    }
  })
}