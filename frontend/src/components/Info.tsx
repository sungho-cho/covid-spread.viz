import { useState, memo } from 'react'
import Box from '@mui/material/Box'
import IconButton from '@mui/material/IconButton'
import EmailIcon from '@mui/icons-material/Email'
import InfoIcon from '@mui/icons-material/Info'
import GitHubIcon from '@mui/icons-material/GitHub'
import LinkedInIcon from '@mui/icons-material/LinkedIn'
import ShareIcon from '@mui/icons-material/Share'
import Link from '@mui/material/Link'
import Modal from '@mui/material/Modal'
import Popover from '@mui/material/Popover'
import Typography from '@mui/material/Typography'
import './Info.css'

interface PopoverOrigin {
  vertical: "bottom" | "center" | "top" | number,
  horizontal: "left" | "center" | "right" | number,
}

const Info = () => {
  const [open, setOpen] = useState(false)
  const [popoverOpen, setPopoverOpen] = useState(false)
  const [popoverText, setPopoverText] = useState("")

  const handleOpen = () => setOpen(true)
  const handleClose = () => setOpen(false)
  const copyVizURL = () => {
    navigator.clipboard.writeText(vizURL)
    setPopoverText("Copied URL to clipboard")
    setPopoverOpen(true)
    setTimeout(() => setPopoverOpen(false), popoverDelay)
  }
  const copyEmail = () => {
    navigator.clipboard.writeText(email)
    setPopoverText("Copied email address to clipboard")
    setPopoverOpen(true)
    setTimeout(() => setPopoverOpen(false), popoverDelay)
  }
  const openURL = (url: string) => {
    window.open(url, '_blank')?.focus()
  }

  const iconSize = '20px'
  const popoverDelay = 1750
  const vizURL = "example.com" // TODO: Change to website address
  const email = "sungh5c@gmail.com"
  const githubURL = "https://github.com/sungho-cho"
  const linkedinURL = "https://linkedin.com/in/sungho-cho/"
  const jhuURL = "https://github.com/CSSEGISandData/COVID-19"
  const jhuRecoveryURL = "https://github.com/CSSEGISandData/COVID-19/issues/4465"
  const modalStyle = {
    position: 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    width: 400,
    bgcolor: 'background.paper',
    boxShadow: 24,
    p: 4,
    outline: 0,
    color: 'white',
  }
  const anchorOrigin: PopoverOrigin = {
    vertical: "bottom",
    horizontal: "center",
  }
  const transformOrigin: PopoverOrigin = {
    vertical: "top",
    horizontal: "center",
  }

  return (
    <div className="info-button">
      <Popover open={popoverOpen} anchorOrigin={anchorOrigin} transformOrigin={transformOrigin}>
        <Typography sx={{ p: 1.5, fontSize: "14px" }}>{popoverText}</Typography>
      </Popover>
      <IconButton aria-label="info" onClick={handleOpen}>
        <InfoIcon sx={{ fontSize: 30 }} />
      </IconButton>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="info-modal-title"
        aria-describedby="info-modal-description"
      >
        <Box sx={modalStyle}>
          <Typography id="info-modal-title" variant="h6" gutterBottom component="h2">
            Covid-19 Spread Visualizer
            <IconButton aria-label="share" onClick={copyVizURL}>
              <ShareIcon sx={{ fontSize: iconSize }} />
            </IconButton>
          </Typography>
          <div className="info-modal-name">
            <Typography variant="subtitle1" component="h3" sx={{ fontWeight: 700 }}>
              Sungho Cho
            </Typography>
            <IconButton aria-label="linkedin" onClick={() => openURL(linkedinURL)}>
              <LinkedInIcon sx={{ fontSize: iconSize }} />
            </IconButton>
            <IconButton aria-label="github" onClick={() => openURL(githubURL)}>
              <GitHubIcon sx={{ fontSize: iconSize }} />
            </IconButton>
            <IconButton aria-label="email" onClick={copyEmail}>
              <EmailIcon sx={{ fontSize: iconSize }} />
            </IconButton>
          </div>
          <Typography id="info-modal-career" variant="subtitle2" component="h4" gutterBottom>
            Carnegie Mellon '20 / Uber ATG / Aurora Innovation
          </Typography>
          <div id="info-modal-description">
            <Typography variant="body2">
              This website's goal is to deliver a fast visualizer for the historical data of COVID-19 since the outbreak to the present day.
              It was built in order for us to better understand the tragic impact of the coronavirus pandemic over the world population.
            </Typography>
            <Typography variant="body2">
              The data is from
              <Link href={jhuURL} target="_blank" rel="noopener"> JHU CSSE COVID-19 Data</Link>.
              It is updated daily when JHU CSSE publishes new data.
              Note that recovery data for United States is absent since December 14, 2020,
              as well as recovery data for all countries since August 5th, 2021,
              as JHU CSSE stopped tracking recovery data
              (<Link href={jhuRecoveryURL} target="_blank" rel="noopener">view details</Link>).
            </Typography>
            <Typography variant="body2">
              This website was developed using React.js, GoLang, Mapbox, gRPC, and Envoy.
              The web application as well as the back-end server are not open-sourced,
              but the GitHub repository is semi-public to give a sense of the general project structure.
            </Typography>
            <Typography variant="body2">
              Please feel free to email me with any feedback or questions :)
            </Typography>
          </div>
        </Box>
      </Modal>
    </div >
  )
}

export default memo(Info);