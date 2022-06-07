import React, { useEffect, useState } from 'react'
import "./App.css"
import CircularProgress from '@mui/material/CircularProgress'
import { ThemeProvider, createTheme } from '@mui/material/styles';

import CovidMap from './CovidMap'
import { getAllData } from '../api/grpc_client'
import { CountriesData } from '../protos/covid_data_pb'

const App = () => {
  const [loading, setLoading] = useState(true)
  const [data, setData] = useState<{ [dateKey: number]: CountriesData } | null>(null)
  const [firstDate, setFirstDate] = useState<Date | null>(null)
  const [lastDate, setLastDate] = useState<Date | null>(null)

  useEffect(() => {
    tryFetchingData()
  }, [])

  const tryFetchingData = () => {
    getAllData(
      (data: { [dateKey: number]: CountriesData }, firstDate: Date, lastDate: Date) => {
        setData(data)
        setFirstDate(firstDate)
        setLastDate(lastDate)
        setLoading(false)
      },
      () => {
        setTimeout(() => tryFetchingData(), 5000)
      }
    )
  }
  const darkTheme = createTheme({
    palette: {
      mode: 'dark',
    },
  });
  const props = { data: data!, firstDate: firstDate!, lastDate: lastDate! }

  return (
    <ThemeProvider theme={darkTheme}>
      {loading ? <div className="loading"><CircularProgress /></div> : <CovidMap {...props} />}
    </ThemeProvider>
  )
}

export default App;
