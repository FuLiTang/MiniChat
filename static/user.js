function mesMini(data) {
    $("#mes").prepend("<div class=\"layui-row\" id=\"show\" onclick=\"f(" +
        data.name +
        ")\">\n" +
        "            <div style=\"width: 100%;height: 50px;\" class=\"layui-row\">\n" +
        "              <div style=\"height: 98%;width:100%;padding: 5px\" class=\"layui-col-md12\">\n" +
        "                <div class=\"layui-col-md3\" style=\"padding-left: 10px\">\n" +
        "                  <img src=" +
        data.img +
        " style=\"width: 40px;border-radius: 20px;\">\n" +
        "                </div>\n" +
        "                <div class=\"layui-col-md9\">\n" +
        "                  <div class=\"layui-row\">\n" +
        "                    <strong style=\"width: auto;font-weight: bold;\">" +
        data.name +
        "</strong>\n" +
        "                    <small id=\"date\" style=\"color: #aaa;float: right;font-size: smaller;\">" +
        data.time +
        "</small>\n" +
        "                  </div>\n" +
        "                 <div class=\"layui-row\" style=\"margin: 5px 0 0 0;overflow: hidden\" >\n" +
        "                   <div class=\"layui-col-md12\">\n" +
        "                     <small id=\"date\" class=\"layui-col-md10\" style=\"padding-right:10px;color: #aaa;font-size: 12px;overflow: hidden;text-overflow: ellipsis;white-space: nowrap\">" +
        data.context +
        "</small>\n" +
        "                     <span class=\"layui-badge layui-bg-cyan layui-col-md2\">1</span>\n" +
        "                   </div>\n" +
        "\n" +
        "                 </div>\n" +
        "                </div>\n" +
        "              </div>\n" +
        "            </div>\n" +
        "            <fieldset class=\"layui-elem-field layui-field-title layui-row\" style=\"margin: 0 10px 0 60px;\">\n" +
        "            </fieldset>\n" +
        "          </div>");
}
    var ws = null;
    var Data = new Array()
    ws = new WebSocket("ws://localhost:8080/socket/socketall?sender=" + "{{.user.Name}}" + "&identity=" + "{{.user.Identity}}")
    ws.onerror = function () {
    console.log("连接错误")
}
    ws.onopen = function () {
    console.log("连接成功，状态码：" + ws.readyState)
}
    ws.onclose = function () {
    console.log("连接已关闭")
}
    window.onbeforeunload = function () {
    ws.close()
}
    ws.onmessage = function (data) {
    var c = JSON.parse(data.data)
    var arrName = c
    Data.push(arrName)
    var mesData = {
    type:c.type,
    context:c.context,
    time:c.time,
}
    var sessionMes = sessionStorage.getItem(c.name)
    if (sessionMes===null){
    var userMesData = [mesData]
    sessionStorage.setItem(c.name,JSON.stringify(userMesData))
}else {
    var ll = JSON.parse(sessionMes)
    ll.push(mesData)
    sessionStorage.setItem(c.name,JSON.stringify(ll))

}
    var slr = "/web/miniChat?name=" + "{{.user.Name}}" + "&identity=" + "{{.user.Identity}}" + "&client=" + c.name
    if ($('#' + c.name).length > 0) {
    document.getElementById(c.name + 'a').innerText = c.name + c.context
} else {
    mesMini(c)
}
}
    function imageAjax(data) {
    $.ajax({})
}
