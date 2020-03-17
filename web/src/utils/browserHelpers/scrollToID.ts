export default (id: string): void => {
  const regexIsID = /^\/?#[a-zA-Z\d-]+/
  const isLinkID = regexIsID.test(id)
  if (!isLinkID) {
    return
  }
  id = id.startsWith('/') ? id.substr(2) : id.substr(1)
  id = id.startsWith('#') ? id.substr(1) : id
  const idElement = document.getElementById(id)
  if (!idElement) {
    return
  }
  const { offsetTop } = idElement
  window.scroll({
    top: offsetTop,
    left: 0,
    behavior: 'smooth'
  })
}
