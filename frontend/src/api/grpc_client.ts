import { Date as ProtoDate, GetCountriesDataRequest, CountriesData } from "../protos/covid_data_pb";
import { CovidDataClient } from "../protos/Covid_dataServiceClientPb";

const gRPC_client = new CovidDataClient('http://localhost:8080');

export function get_countries_data(date: Date, callback: (countries_data?: CountriesData) => void) {
  var proto_date = new ProtoDate();
  proto_date.setYear(date.getFullYear());
  proto_date.setMonth(date.getMonth() + 1);
  proto_date.setDay(date.getDate());
  var request = new GetCountriesDataRequest();
  request.setDate(proto_date);
  gRPC_client.getCountriesData(request, {}, (err, response) => {
    if (response == null) {
      console.log(err)
    } else {
      callback(response.getCountriesData());
    }
  });
}