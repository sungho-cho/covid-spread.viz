import { useState, memo } from 'react'
import Box from '@mui/material/Box'
import IconButton from '@mui/material/IconButton'
import EmailIcon from '@mui/icons-material/Email'
import InfoIcon from '@mui/icons-material/Info'
import GitHubIcon from '@mui/icons-material/GitHub'
import LinkedInIcon from '@mui/icons-material/LinkedIn'
import PublicIcon from '@mui/icons-material/Public'
import ShareIcon from '@mui/icons-material/Share'
import Link from '@mui/material/Link'
import Modal from '@mui/material/Modal'
import Popover from '@mui/material/Popover'
import Tooltip from '@mui/material/Tooltip'
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

  const handleOpen = () => {
    setOpen(true)
    window.dataLayer.push({ event: 'info-click' })
  }
  const handleClose = () => setOpen(false)
  const copyVizURL = () => {
    navigator.clipboard.writeText(vizURL)
    window.dataLayer.push({ event: 'share-click' })
    setPopoverText("Copied URL to clipboard")
    setPopoverOpen(true)
    setTimeout(() => setPopoverOpen(false), popoverDelay)
  }
  const openURL = (url: string) => {
    window.open(url, '_blank')?.focus()
    if (url === emailURL) {
      window.dataLayer.push({ event: 'email-click' })
    }
    else if ([linkedinURL, githubURL, personalURL].includes(url)) {
      window.dataLayer.push({
        event: 'profile-link-click',
        eventProps: { 'profile-link': url }
      })
    }
  }

  const iconSize = '20px'
  const popoverDelay = 1750
  const vizURL = "https://www.covidviz.com"
  const personalURL = "https://sunghocho.com"
  const emailURL = "mailto:sungh5c@gmail.com"
  const githubURL = "https://github.com/sungho-cho"
  const repoURL = "https://github.com/sungho-cho/covid-spread.viz"
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
            COVID-19 Spread Visualizer
            <Tooltip title="Copy URL">
              <IconButton aria-label="share" onClick={copyVizURL}>
                <ShareIcon sx={{ fontSize: iconSize }} />
              </IconButton>
            </Tooltip>
          </Typography>
          <div className="info-modal-name">
            <Typography variant="subtitle1" component="h3" sx={{ fontWeight: 700 }}>
              Sungho Cho
            </Typography>
            <Tooltip title="LinkedIn">
              <IconButton aria-label="linkedin" onClick={() => openURL(linkedinURL)}>
                <LinkedInIcon sx={{ fontSize: iconSize }} />
              </IconButton>
            </Tooltip>
            <Tooltip title="GitHub">
              <IconButton aria-label="github" onClick={() => openURL(githubURL)}>
                <GitHubIcon sx={{ fontSize: iconSize }} />
              </IconButton>
            </Tooltip>
            <Tooltip title="Personal Website">
              <IconButton aria-label="personal" onClick={() => openURL(personalURL)}>
                <PublicIcon sx={{ fontSize: iconSize }} />
              </IconButton>
            </Tooltip>
            <Tooltip title="Email">
              <IconButton aria-label="email" onClick={() => openURL(emailURL)}>
                <EmailIcon sx={{ fontSize: iconSize }} />
              </IconButton>
            </Tooltip>
          </div>
          <Typography id="info-modal-career" variant="subtitle2" component="h4" gutterBottom>
            Carnegie Mellon '20 / Uber ATG / Aurora Innovation
          </Typography>
          <div id="info-modal-description">
            <Typography variant="body2">
              This website's goal is to provide a fast, interactive visualization of COVID-19's historical data since the outbreak to the present day.
              It was built in order for us to better understand the tragic impact of the coronavirus pandemic over the world population.
            </Typography>
            <Typography variant="body2">
              The data is from <Link href={jhuURL} target="_blank" rel="noopener">JHU CSSE COVID-19 Data</Link>.
              It is updated daily when JHU CSSE publishes new data on their data repository.
            </Typography>
            <Typography variant="body2">
              <strong style={{ color: "#F3AF16" }}>Note:</strong> Recovery data for United States has been no longer recorded since December 14, 2020,
              as well as recovery data for all countries since August 5th, 2021,
              as Johns Hopkins CSSE has stopped tracking recovery data
              (<Link href={jhuRecoveryURL} target="_blank" rel="noopener">view details</Link>).
            </Typography>
            <Typography variant="body2">
              This website was developed using React.js, GoLang, Mapbox, and gRPC.
              The web application and the back-end server are not fully open-sourced,
              but the <Link href={repoURL} target="_blank" rel="noopener">GitHub repository</Link> is
              semi-public to introduce the general project structure and the tech stack.
            </Typography>
            <Typography variant="body2">
              Please feel free to <Link href={emailURL} target="_blank" rel="noopener"> email me</Link> with any feedback or questions :)
            </Typography>
          </div>
        </Box>
      </Modal>
    </div >
  )
}

export default memo(Info);