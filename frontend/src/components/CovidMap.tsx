import { useRef, useEffect, useState } from 'react'
// @ts-ignore
import mapboxgl from 'mapbox-gl/dist/mapbox-gl'
// @ts-ignore
import MapboxWorker from 'mapbox-gl/dist/mapbox-gl-csp-worker';
import 'mapbox-gl/dist/mapbox-gl.css'
import CircularProgress from '@mui/material/CircularProgress'
import Typography from '@mui/material/Typography';
import { MAPBOX_TOKEN, stops, setFeatureStates, } from '../api/mapbox_config'
import { CountriesData } from '../protos/covid_data_pb';
import useInterval from '../api/use_interval';
import DataSelect from './DataSelect'
import DateController from './DateController'
import Info from './Info'
import Legend from './Legend'
import './CovidMap.css';

mapboxgl.workerClass = MapboxWorker
mapboxgl.accessToken = MAPBOX_TOKEN

interface CovidMapProps {
  data: { [dateKey: number]: CountriesData },
  firstDate: Date,
  lastDate: Date,
}

interface FeatureState {
  confirmed: number,
  recovered: number,
  deaths: number,
}

interface MapboxMouseEvent {
  lngLat: {
    lng: number,
    lat: number
  },
  features: Array<mapboxgl.feature>,
}

// number of days of data to show per second (e.g. 5 means 5 days per second)
export const speedOptions =
  [-30, -15, -10, -5, -4, -3, -2, -1, -0.5, -0.25, 0,
    0.25, 0.5, 1, 2, 3, 4, 5, 10, 15, 30]
export const defaultSpeed = 5 // default speed of 5 days per second

const CovidMap = (props: CovidMapProps) => {
  /* Props */
  const data = props.data
  const firstDate = props.firstDate
  const lastDate = props.lastDate

  /* State, Ref, Callback Hooks */
  const defaultSpeedIdx = speedOptions.indexOf(defaultSpeed)
  const mapContainer = useRef<HTMLDivElement | null>(null)
  const map = useRef<mapboxgl.Map | null>(null)
  const popup = useRef<mapboxgl.Popup | null>(null)
  const [loading, setLoading] = useState(true)
  const [date, setDate] = useState(firstDate)
  const [playing, setPlaying] = useState(false)
  const [speedIdx, setSpeedIdx] = useState(defaultSpeedIdx)
  const [delay, setDelay] = useState<number | null>(1000 / defaultSpeed)
  const [displayData, setDisplayData] = useState("confirmed")
  const displayDataRef = useRef("confirmed")
  const hoveredFeature = useRef<mapboxgl.MapboxGeoJSONFeature | null>(null)

  /* Effect and Interval Hooks */
  useEffect(() => { // Initialize map when component mounts
    map.current = new mapboxgl.Map({
      container: mapContainer.current == null ? "" : mapContainer.current,
      style: 'mapbox://styles/mapbox/dark-v10',
      center: [5.76, 26.77],
      zoom: 1.59
    })
    popup.current = new mapboxgl.Popup({
      closeButton: false,
      closeOnClick: false
    })
    map.current.on('load', () => {
      map.current!.addSource('country-source', {
        type: 'vector',
        url: 'mapbox://mapbox.country-boundaries-v1'
      })
      map.current!.addLayer({
        id: 'country-layer',
        source: 'country-source',
        'source-layer': 'country_boundaries',
        type: 'fill',
        paint: {
          'fill-opacity': 1.0,
        }
      }, 'admin-0-boundary-bg')
      map.current!.addLayer({
        id: 'country-highlight-layer',
        source: 'country-source',
        'source-layer': 'country_boundaries',
        type: 'line',
        paint: {
          'line-color': '#dcdcdc'
        }
      })
      map.current!.setPaintProperty('country-label', 'text-color', "#ffffff")
      map.current!.setPaintProperty('country-label', 'text-opacity', 0.7)
      map.current!.setPaintProperty('country-layer', 'fill-color', [
        'step',
        ['number', ['feature-state', displayData], 0],
        defaultCountryColor,
        ...stops,
      ])
      map.current!.setPaintProperty('country-highlight-layer', 'line-width', [
        'case',
        ['boolean', ['feature-state', 'hover'], false],
        2.0,
        0.0,
      ])
      map.current!.on('mousemove', 'country-layer', (e: MapboxMouseEvent) => {
        // Highlight on hover
        if (hoveredFeature.current && hoveredFeature.current.id === e.features![0].id) {
          popup.current!.setLngLat([e.lngLat.lng, e.lngLat.lat])
        } else {
          if (hoveredFeature.current) {
            map.current!.setFeatureState(
              { id: hoveredFeature.current!.id, source: 'country-source', sourceLayer: 'country_boundaries' },
              { hover: false }
            )
          }
          hoveredFeature.current = e.features![0]
          map.current!.setFeatureState(
            { id: hoveredFeature.current!.id, source: 'country-source', sourceLayer: 'country_boundaries' },
            { hover: true }
          )
          // Show popup on hover
          map.current!.getCanvas().style.cursor = 'pointer'
          const text = getPopupText(e.features![0].properties!.name_en, e.features![0].state as FeatureState)
          popup.current!.setLngLat([e.lngLat.lng, e.lngLat.lat]).setHTML(text).addTo(map.current!)
        }
      })
      map.current!.on('mouseleave', 'country-layer', () => {
        // Disable highlight on leave
        if (hoveredFeature.current) {
          map.current!.setFeatureState(
            { id: hoveredFeature.current!.id, source: 'country-source', sourceLayer: 'country_boundaries' },
            { hover: false }
          )
        }
        hoveredFeature.current = null
        // Hide popup on leave
        map.current!.getCanvas().style.cursor = ''
        popup.current!.remove()
      })
      map.current!.once('idle', () => { setLoading(false) })
    })
    return () => map.current!.remove(); // clean up on unmount
  }, [])

  useEffect(() => {
    if (!map.current! || !map.current!.isStyleLoaded()) return;
    const countriesData = data[date.getTime()]
    setFeatureStates(countriesData, (feature, state) => map.current!.setFeatureState(feature, state))
    if (hoveredFeature.current) {
      const featureState: FeatureState = map.current!.getFeatureState({
        id: hoveredFeature.current!.id,
        source: 'country-source',
        sourceLayer: 'country_boundaries'
      }) as FeatureState
      const text = getPopupText(hoveredFeature.current!.properties!.name_en, featureState)
      popup.current!.setHTML(text).addTo(map.current!)
    }
  }, [date])

  useEffect(() => {
    if (!map.current! || !map.current!.isStyleLoaded()) return;
    map.current!.setPaintProperty('country-layer', 'fill-color', [
      'step',
      ['number', ['feature-state', displayData], 0],
      defaultCountryColor,
      ...stops,
    ])
    displayDataRef.current = displayData
  }, [displayData])

  useEffect(() => {
    const speed = speedOptions[speedIdx]
    if (speed == 0) setDelay(null)
    else setDelay(Math.abs(1000 / speed))
  }, [speedIdx])

  useInterval(() => {
    if (!playing) return
    const newDate = new Date(Date.UTC(
      date.getUTCFullYear(),
      date.getUTCMonth(),
      speedOptions[speedIdx] > 0 ? date.getUTCDate() + 1 : date.getUTCDate() - 1
    ))
    if (newDate.getTime() > lastDate.getTime() || newDate.getTime() < firstDate.getTime()) {
      setPlaying(false)
      return
    }
    setDate(newDate)
  }, delay!);

  /* Constants */
  const getPopupText = (countryName: string, featureState: FeatureState) => {
    let confirmedTxt =
      `Confirmed: ${featureState.confirmed ? featureState.confirmed.toLocaleString('en-US') : '0'}`
    let recoveredTxt =
      `Recovered: ${featureState.recovered ? featureState.recovered.toLocaleString('en-US') : '0'}`
    let deathsTxt =
      `Deaths: ${featureState.deaths ? featureState.deaths.toLocaleString('en-US') : '0'}`
    confirmedTxt = displayDataRef.current == "confirmed" ? `<b>${confirmedTxt}</b>` : confirmedTxt
    recoveredTxt = displayDataRef.current == "recovered" ? `<b>${recoveredTxt}</b>` : recoveredTxt
    deathsTxt = displayDataRef.current == "deaths" ? `<b>${deathsTxt}</b>` : deathsTxt
    return `
      <b>${countryName}</b><br>
      ${confirmedTxt}<br>
      ${recoveredTxt}<br>
      ${deathsTxt}<br>
    `
  }
  const formatDate = (date: Date) => {
    let year = date.getFullYear();
    let month = (1 + date.getMonth()).toString().padStart(2, '0');
    let day = date.getDate().toString().padStart(2, '0');
    return month + '/' + day + '/' + year;
  }
  const defaultCountryColor = "#343332"
  const selectProps = { displayData, setDisplayData }
  const controllerProps = { date, setDate, speedIdx, setSpeedIdx, playing, setPlaying, firstDate, lastDate }

  return (
    <div>
      <div ref={mapContainer} className="map-container" />
      {loading ? <div className="map-loading"><CircularProgress /></div> : <></>}
      <div className="date-label">
        <Typography variant="overline" sx={{ fontWeight: 700, fontSize: '14px' }}>
          {formatDate(date)}
        </Typography>
      </div>
      <Info />
      <Legend displayData={displayData} stops={stops} />
      <DataSelect {...selectProps} />
      <DateController {...controllerProps} />
    </div>
  )
}

export default CovidMap;