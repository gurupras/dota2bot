const canvas = document.querySelector('#canvas')
const ctx = canvas.getContext('2d')

let imageLoaded = false
let heroData
// Fetch hero json file
; (async () => {
  const response = await fetch('/static/resources/by-name.json')
  heroData = await response.json()
})()

// Load the image
const img = new Image()
img.onload = () => { imageLoaded = true }
img.src = '/static/resources/minimap_hero_sheet.png'

const ws = new ReconnectingWebSocket(`ws://${location.host}/ws`)

ws.onmessage = (msg) => {
  const { data } = msg
  const json = JSON.parse(data)
  ctx.clearRect(0, 0, canvas.width, canvas.height)
  for (const entry of json) {
    const { name, isHero, point: { x, y }, radius, red, blue, green } = entry
    if (isHero && imageLoaded && heroData) {
      const { [name]: data } = heroData
      if (data) {
        const { x: sx, y: sy } = data
        ctx.drawImage(img, sx, sy, 32, 32, x, y, 32, 32)
        continue
      }
    }
    ctx.beginPath()
    ctx.arc(x, y, radius, 0, 2 * Math.PI, false)
    ctx.fillStyle = rgbToHex(red, green, blue)
    ctx.fill()
  }
}

function componentToHex(c) {
  const hex = c.toString(16)
  return hex.length == 1 ? "0" + hex : hex
}

function rgbToHex(r, g, b) {
  return `#${componentToHex(r)}${componentToHex(g)}${componentToHex(b)}`
}