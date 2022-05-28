function passwordCheck() {
    let password = document.getElementById("reg_pwd").value
    let passwordConf = document.getElementById("reg_pwd_conf").value

    if (password == passwordConf) {
        if (password.length != 0) {
            if (password.length < 8) {
                alert("Password minimal length is 8 symbols.");
                document.getElementById("reg_form").reset();
            } else {
                document.getElementById("reg_form").submit();
            }
        } else{
            alert("Input password and confirm it.")
            document.getElementById("reg_form").reset();
        }
    } else {
        alert("Passwords don't match.")
        document.getElementById("reg_form").reset();
    }
}
