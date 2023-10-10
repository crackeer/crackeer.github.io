async function doUploadFile(file) {
    var url = window.location.href + file.name
    const formData = new FormData()
    formData.append('file', file)
    return await fetch(url, {
        method: 'PUT',
        body: formData
    })
}