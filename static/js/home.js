var sock = null;
var wsuri = "ws://localhost:8080/ws";

window.onload = function () {
    console.log("onload");
    sock = new WebSocket(wsuri);
    sock.onopen = function () {
        console.log("connected to " + wsuri);
    };
    sock.onclose = function (e) {
        console.log("connection closed (" + e.code + ")");
    };
    sock.onmessage = function (e) {
        // console.log("message received: " + e.data);
        var obj = JSON.parse(e.data);
        // console.log(obj);
        var i;
        for (i = 0; i < obj.EnvData.length; i++) {
            var doc = document.getElementById(obj.Node.NodeID + "-" + obj.EnvData[i].EnvName);
            doc.innerText = obj.EnvData[i].EnvValue;
        }
    };
};

function send() {
    var msg = document.getElementById('message').value;
    sock.send(msg);
};