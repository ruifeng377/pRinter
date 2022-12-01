function uploadFile() {
    document.uploadOrPrint.action = "/uploadFile"
    document.uploadOrPrint.submit()
}

function printFile() {
    document.uploadOrPrint.action = "/printFile"
    document.uploadOrPrint.submit()
}