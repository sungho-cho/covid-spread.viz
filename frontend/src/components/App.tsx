import React, { useEffect } from 'react'
import './App.css'

import CovidMap from './CovidMap'
import { get_countries_data } from '../api/grpc_client'

function App() {
  const getData = () => {
    const date = new Date("2022-5-22")
    get_countries_data(date, (countriesData) => { console.log(countriesData) })
  }

  useEffect(() => {
    getData();
  }, [])

  return (
    <CovidMap />
  );
}

export default App;
