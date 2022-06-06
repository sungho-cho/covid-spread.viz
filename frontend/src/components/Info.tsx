import { useState, memo } from 'react'
import Box from '@mui/material/Box'
import IconButton from '@mui/material/IconButton'
import InfoIcon from '@mui/icons-material/Info'
import GitHubIcon from '@mui/icons-material/GitHub'
import LinkedInIcon from '@mui/icons-material/LinkedIn'
import EmailIcon from '@mui/icons-material/Email'
import Modal from '@mui/material/Modal'
import Typography from '@mui/material/Typography'
import './Info.css'

const Info = () => {
  const [open, setOpen] = useState(false)
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
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
          </Typography>
          <Typography id="info-modal-name" variant="subtitle1" component="h3">
            Sungho Cho
            <GitHubIcon sx={{ fontSize: '20px' }} />
            <LinkedInIcon sx={{ fontSize: '20px' }} />
            <EmailIcon sx={{ fontSize: '20px' }} />
          </Typography>
          <Typography id="info-modal-career" variant="subtitle2" component="h4" gutterBottom>
            Carnegie Mellon '20 / Uber ATG / Aurora Innovation
          </Typography>
          <Typography id="info-modal-description" variant="body2">
            This is a description of this website
          </Typography>
        </Box>
      </Modal>
    </div>
  )
}

export default memo(Info);