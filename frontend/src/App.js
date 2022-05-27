import React, { useEffect } from 'react';
// import React from 'react';
import './App.css';
import Map, { Marker } from 'react-map-gl';

import 'mapbox-gl/dist/mapbox-gl.css';

import { Date, GetCountriesDataRequest, GetCountriesDataResponse } from "./protos/covid_data_pb";
import { CovidDataClient } from "./protos/covid_data_grpc_web_pb";

const MAPBOX_TOKEN = 'pk.eyJ1IjoiZGlkb2c5NSIsImEiOiJja3d1Ymd5N2swaGRjMm9xbjU1OGRkMDl0In0.jkmsYL4EYs3VmsU6zUxbPQ';
var gRPC_client = new CovidDataClient('http://localhost:8080');

function App() {
  const getData = () => {
    console.log("called");
    var date = new Date();
    date.setYear(2022);
    date.setMonth(5);
    date.setDay(22);
    var request = new GetCountriesDataRequest();
    request.setDate(date);
    console.log(request);
    gRPC_client.getCountriesData(request, {}, (err, response) => {
      if (response == null) {
        console.log(err)
      } else {
        console.log("Success!")
        console.log(response.getCountriesData())
      }
    });
    return "";
  }

  useEffect(() => {
    getData()
  }, []);

  return (
    <Map
      initialViewState={{
        latitude: 37.8,
        longitude: -122.4,
        zoom: 14
      }}
      style={{ width: 800, height: 600 }}
      mapStyle="mapbox://styles/mapbox/streets-v9"
      mapboxAccessToken={MAPBOX_TOKEN}
    >
      <Marker longitude={-122.4} latitude={37.8} color="red" />
    </Map>
  );
}

export default App;
