import { useState, memo } from 'react'
import Box from '@mui/material/Box'
import IconButton from '@mui/material/IconButton'
import EmailIcon from '@mui/icons-material/Email'
import InfoIcon from '@mui/icons-material/Info'
import GitHubIcon from '@mui/icons-material/GitHub'
import LinkedInIcon from '@mui/icons-material/LinkedIn'
import ShareIcon from '@mui/icons-material/Share'
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
  const email = "email" // TODO: Change to real email address
  const githubURL = "https://github.com/sungho-cho"
  const linkedinURL = "https://linkedin.com/in/sungho-cho/"
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
            <Popover open={popoverOpen} anchorOrigin={anchorOrigin} transformOrigin={transformOrigin}>
              <Typography sx={{ p: 2 }}>{popoverText}</Typography>
            </Popover>
          </Typography>
          <div className="info-modal-name">
            <Typography variant="subtitle1" component="h3">
              Sungho Cho
            </Typography>
            <IconButton aria-label="github" onClick={() => openURL(githubURL)}>
              <GitHubIcon sx={{ fontSize: iconSize }} />
            </IconButton>
            <IconButton aria-label="linkedin" onClick={() => openURL(linkedinURL)}>
              <LinkedInIcon sx={{ fontSize: iconSize }} />
            </IconButton>
            <IconButton aria-label="email" onClick={copyEmail}>
              <EmailIcon sx={{ fontSize: iconSize }} />
            </IconButton>
          </div>
          <Typography id="info-modal-career" variant="subtitle2" component="h4" gutterBottom>
            Carnegie Mellon '20 / Uber ATG / Aurora Innovation
          </Typography>
          <Typography id="info-modal-description" variant="body2">
            This is a description of this website
          </Typography>
        </Box>
      </Modal>
    </div >
  )
}

export default memo(Info);