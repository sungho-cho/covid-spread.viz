import { memo } from 'react'
import { styled } from '@mui/material/styles'
import InputBase from '@mui/material/InputBase'
import MenuItem from '@mui/material/MenuItem'
import Select, { SelectChangeEvent } from '@mui/material/Select'
import './DataSelect.css'

interface DataSelectProps {
  displayData: string,
  setDisplayData: (displayData: string) => void,
}

const DataSelect = (props: DataSelectProps) => {
  const { displayData, setDisplayData } = props

  const StyledInput = styled(InputBase)(({ theme }) => ({
    'label + &': {
      marginTop: theme.spacing(3),
    },
    '& .MuiInputBase-input': {
      borderRadius: 4,
      color: 'white',
      backgroundColor: 'rgba(255, 255, 255, 0.12)',
      fontSize: '14px',
      fontWeight: 700,
      lineHeight: '6px',
      textAlign: 'center',
      verticalAlign: 'middle',
      padding: '16px 10px 2px 12px',
      position: 'relative',
      opacity: '0.8',
    },
  }))

  const handleChange = (event: SelectChangeEvent) => {
    setDisplayData(event.target.value as string)
  }

  return (
    <div className="data-select">
      <Select
        id="data-select"
        value={displayData}
        label="Data to display"
        onChange={handleChange}
        input={<StyledInput />}
      >
        <MenuItem value={"confirmed"}>Confirmed</MenuItem>
        <MenuItem value={"recovered"}>Recovered</MenuItem>
        <MenuItem value={"deaths"}>Deaths</MenuItem>
      </Select>
    </div>
  )
}

export default memo(DataSelect);