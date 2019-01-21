function checkInput() {
    var uname = document.getElementById("uname");
    if (uname.value.length == 0) {
        showInfo("请输入账号");
        return false;
    } else if(uname.value.indexOf(" ") != -1) {
        showInfo("账户中不能包含空格");
        return false;
    }
    var pwd = document.getElementById("pwd");
    if (pwd.value.length == 0) {
        showInfo("请输入密码");
        return false;
    } else if(uname.value.indexOf(" ") != -1) {
        showInfo("密码中不能包含空格");
        return false;
    }
    return true;
}

function showInfo(info) {
    var hint = document.getElementById("hint");
    if (hint.hasAttribute("hidden")) {
        hint.removeAttribute("hidden");
    }
    hint.innerText = info;
}

function login() {
    if (!checkInput()) {
        return;
    }
    var xmlhttp;
    if (window.XMLHttpRequest) {
        // code for IE7+, Firefox, Chrome, Opera, Safari
        xmlhttp = new XMLHttpRequest();
    } else {
        // code for IE6, IE5
        xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
    }
    if (xmlhttp != null) {
        xmlhttp.onreadystatechange = function () {
            if (xmlhttp.readyState === 4) {
                if(xmlhttp.responseText == "true") {
                    window.location.href = "/";
                    return;
                }
                return showInfo(xmlhttp.responseText);
            } else {
                // return fail(xmlhttp.status);
                return "";
            }
        };
    }
    // 发送请求:
    xmlhttp.open('POST', '/login');
    var uname = document.getElementById("uname");
    var pwd = document.getElementById("pwd");
    var autologin = "off";
    if (document.getElementById("autoLogin").checked) {
        autologin = "on";
    }
    xmlhttp.setRequestHeader("Content-type","application/x-www-form-urlencoded");
    xmlhttp.send("uname="+uname.value+"&pwd="+pwd.value+"&autoLogin="+autologin);
}

var submit = document.getElementById("submit");
submit.addEventListener("click", login);