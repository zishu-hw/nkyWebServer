function checkInput() {
    var uname = document.getElementById("uname");
    var uname_h = document.getElementById("uname-h");
    if (uname.value.length == 0) {
        uname_h.innerText = "请输入账户名";
        return false;
    }
    uname_h.innerText = "";
    var pwd = document.getElementById("pwd");
    var pwd_h = document.getElementById("pwd-h");
    if (pwd.value.length == 0) {
        pwd_h.innerText = "请输入密码";
        return false;
    }
    pwd_h.innerText = "";
    return true;
}

function success(text) {
    var textarea = document.getElementById('pwd-h');
    textarea.innerText = "    " + text;
}

function loadDoc() {
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
                return success(xmlhttp.responseText);
            } else {
                // return fail(xmlhttp.status);
                return "";
            }
        };
    }
    // 发送请求:
    xmlhttp.open('Get', '/login/ajax');
    xmlhttp.send();
}