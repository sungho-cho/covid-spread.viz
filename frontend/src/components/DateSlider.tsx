import React, { useEffect, useState } from 'react'
import { ThemeProvider } from '@mui/material/styles'
import createTheme from '@mui/material/styles/createTheme'
import IconButton from '@mui/material/IconButton'
import Slider, { SliderThumb } from '@mui/material/Slider'
import CoronavirusIcon from '@mui/icons-material/Coronavirus'
import PlayArrowIcon from '@mui/icons-material/PlayArrow'
import FastForwardIcon from '@mui/icons-material/FastForward'
import FastRewindIcon from '@mui/icons-material/FastRewind'
import './DateSlider.css'

interface DateSliderProps {
  date: Date,
  setDate: (date: Date) => void,
  firstDate: Date,
  lastDate: Date,
}

interface Mark {
  value: number,
  label: string,
}

interface VirusThumbComponentProps extends React.HTMLAttributes<unknown> { }

const DateSlider = (props: DateSliderProps) => {
  const { date, setDate, firstDate, lastDate } = props

  const [marks, setMarks] = useState<Array<Mark>>([])

  const ThemedButton = createTheme({
    components: {
      MuiIconButton: {
        styleOverrides: {
          root: {
            ':hover': {
              backgroundColor: 'unset',
            }
          }
        }
      }
    }
  })

  const ThemedSlider = createTheme({
    components: {
      MuiSlider: {
        styleOverrides: {
          root: {
            color: 'rgba(255, 255, 255, 0.12)',
            height: 8,
          },
          track: {
            border: 'none',
          },
          thumb: {
            backgroundColor: 'unset',
            '&.Mui-focusVisible': {
              boxShadow: 'unset',
            },
            '&.Mui-active': {
              boxShadow: 'unset',
            },
            ':hover': {
              boxShadow: 'unset',
            },
          },
          mark: {
            width: '4px',
            height: '4px',
            borderRadius: '9999px',
            backgroundColor: 'rgba(255, 255, 255, 0.40)',
          },
        }
      }
    },
    palette: {
      mode: 'dark',
    },
  })

  useEffect(() => {
    let markMonth = new Date(Date.UTC(firstDate.getUTCFullYear(), firstDate.getUTCMonth()))
    let lastMonth = new Date(Date.UTC(lastDate.getUTCFullYear(), lastDate.getUTCMonth()))
    var newMarks = []
    for (; markMonth.getTime() <= lastMonth.getTime();) {
      const month = markMonth.getUTCMonth() + 1
      var label = ""
      if (month == 1) {
        label = markMonth.getUTCFullYear().toString()
      } else {
        label = ""
      }
      newMarks.push({ value: markMonth.getTime(), label: label })
      markMonth = new Date(Date.UTC(markMonth.getUTCFullYear(), markMonth.getUTCMonth() + 1))
    }
    setMarks(newMarks)
  }, [])

  const iconColor = 'rgba(255, 255, 255, 0.8)'
  const valueLabelFormat = (dateNumber: number) => {
    const newDate = new Date(dateNumber)
    return newDate.toLocaleDateString("en-US");
  }
  const updateDate = (event: Event, newValue: number | number[]) => {
    if (typeof newValue === 'number') {
      setDate(new Date(newValue))
    }
  }
  const VirusThumbComponent = (props: VirusThumbComponentProps) => {
    const { children, ...other } = props
    return (
      <SliderThumb {...other}>
        {children}
        <CoronavirusIcon sx={{ fontSize: 35, color: iconColor }} />
      </SliderThumb>
    )
  }

  return (
    <div className="date-control">
      <div className="date-buttons">
        <ThemeProvider theme={ThemedButton}>
          <IconButton aria-label="rewind">
            <FastRewindIcon sx={{ color: iconColor }} />
          </IconButton>
          <IconButton aria-label="play">
            <PlayArrowIcon sx={{ color: iconColor }} />
          </IconButton>
          <IconButton aria-label="forward">
            <FastForwardIcon sx={{ color: iconColor }} />
          </IconButton>
        </ThemeProvider>
      </div>
      <div className="date-slider">
        <ThemeProvider theme={ThemedSlider}>
          <Slider
            value={date.getTime()}
            aria-label="Small steps"
            defaultValue={firstDate.getTime()}
            getAriaValueText={valueLabelFormat}
            valueLabelFormat={valueLabelFormat}
            step={86400000}
            min={firstDate.getTime()}
            max={lastDate.getTime()}
            marks={marks}
            valueLabelDisplay="auto"
            onChange={updateDate}
            components={{ Thumb: VirusThumbComponent }}
          />
        </ThemeProvider>
      </div>
    </div>
  )
}

export default DateSlider;