import React, { useEffect, useState } from 'react'
import Slider from '@mui/material/Slider'
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

const DateSlider = (props: DateSliderProps) => {
  const { date, setDate, firstDate, lastDate } = props

  const [marks, setMarks] = useState<Array<Mark>>([])

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
        label = month.toString()
      }
      newMarks.push({ value: markMonth.getTime(), label: label })
      markMonth = new Date(Date.UTC(markMonth.getUTCFullYear(), markMonth.getUTCMonth() + 1))
    }
    setMarks(newMarks)
  }, [])

  const valueLabelFormat = (dateNumber: number) => {
    const newDate = new Date(dateNumber)
    return newDate.toLocaleDateString("en-US");
  }

  const updateDate = (event: Event, newValue: number | number[]) => {
    if (typeof newValue === 'number') {
      setDate(new Date(newValue))
    }
  }

  return (
    <div className="dateslider">
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
      />
    </div>
  )
}

export default DateSlider;