import { memo } from 'react'
import { styled } from '@mui/material/styles'
import InfoIcon from '@mui/icons-material/Info'
import PriorityHighIcon from '@mui/icons-material/PriorityHigh'
import Tooltip from '@mui/material/Tooltip'
import Typography from '@mui/material/Typography'
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
  const optionFontSize = 14
  const recoveryMsg1 = "Recovery data is missing for some countries and dates. Click the info button "
  const recoveryMsg2 = " for more info."

  return (
    <div className="data-select">
      <Select
        id="data-select"
        value={displayData}
        label="Data to display"
        onChange={handleChange}
        input={<StyledInput />}
      >
        <MenuItem value={"confirmed"} sx={{ fontSize: optionFontSize }}>Confirmed</MenuItem>
        <Tooltip placement="left" title={
          <Typography fontSize={13}>
            {recoveryMsg1}<InfoIcon sx={{ fontSize: 16 }} />{recoveryMsg2}
          </Typography>}
        >
          <MenuItem value={"recovered"} sx={{ fontSize: optionFontSize }}>
            Recovered
            <PriorityHighIcon sx={{ fontSize: 15 }} />
          </MenuItem>
        </Tooltip>
        <MenuItem value={"deaths"} sx={{ fontSize: optionFontSize }}>Deaths</MenuItem>
      </Select>
    </div >
  )
}

export default memo(DataSelect);