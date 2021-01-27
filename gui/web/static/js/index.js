const canvas = document.querySelector('#canvas')
const ctx = canvas.getContext('2d')

const ws = new ReconnectingWebSocket(`ws://${location.host}/ws`)

ws.onmessage = (msg) => {
  const { data } = msg
  const json = JSON.parse(data)
  ctx.clearRect(0, 0, canvas.width, canvas.height)
  for (const entry of json) {
    const { point: { x, y }, radius, red, blue, green } = entry
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