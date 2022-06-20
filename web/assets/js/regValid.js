const form = document.getElementById('reg_form');
const username = document.getElementById('reg_username');
const email = document.getElementById('reg_email');
const password = document.getElementById('reg_pwd');
const password2 = document.getElementById('reg_pwd_conf');
let i;

form.addEventListener('submit', e => {
    if (!checkInputs()) {
        e.preventDefault();
    }
});

function checkInputs() {
    i = 0
    // trim to remove the whitespaces
    const usernameValue = username.value.trim();
    const emailValue = email.value.trim();
    const passwordValue = password.value.trim();
    const password2Value = password2.value.trim();

    if(usernameValue === '') {
        setErrorFor(username, 'Имя пользователя не может быть пустым');
    } else if(usernameValue.length < 5) {
        setErrorFor(username, 'Имя пользователя должно быть не менее 5 символов');
    } else {
        setSuccessFor(username);
    }

    if(emailValue === '') {
        setErrorFor(email, 'Почта не может быть пустой');
    } else if (!isEmail(emailValue)) {
        setErrorFor(email, 'Не подходящая почта');
    } else {
        setSuccessFor(email);
    }

    if(passwordValue.length < 8) {
        setErrorFor(password, 'Пароль должен быть не менее 8 символов, пробелы не считаются за символы');
    } else {
        setSuccessFor(password);
    }

    if(passwordValue.length < 8) {
        setErrorFor(password2, 'Пароль должен быть не менее 8 символов, пробелы не считаются за символы');
    } else if(passwordValue !== password2Value) {
        setErrorFor(password2, 'Пароли не совпадают');
    } else{
        setSuccessFor(password2);
    }
    console.log(i)
    return i === 0;
}

function setErrorFor(input, message) {
    i++;
    const formControl = input.parentElement;
    const button = formControl.querySelector('button');
    formControl.className = 'form-control error';
    button.ariaLabel = message;
}

function setSuccessFor(input) {
    const formControl = input.parentElement;
    formControl.className = 'form-control success';
}

function isEmail(email) {
    return /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/.test(email);
}