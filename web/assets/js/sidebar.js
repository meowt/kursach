const checkboxNav = document.getElementById('sidebar')
const checkboxProf = document.getElementById('profbar')
const bodyContainer = document.querySelector('.bodyContainer')
let width = window.innerWidth


checkboxNav.addEventListener('change', (event) => {
    let obj = document.querySelector('.navbar');
    let bodyContainer = document.querySelector('.bodyContainer');
    if (event.currentTarget.checked) {
        obj.style.left = '0';
        bodyContainer.style.margin = '0 0 0 150px';
        width -= 150;
        bodyContainer.style.width = width + 'px';

    } else {
        obj.style.left = '-150px';
        bodyContainer.style.margin = '0 0 0 0';
        width += 150;
        bodyContainer.style.width = width + 'px';
    }
})

checkboxProf.addEventListener('change', (event) => {
    let obj = document.querySelector('.profbar');
    if (event.currentTarget.checked) {
        obj.style.right = '0';
        bodyContainer.style.margin = '0 150px 0 0';
        width -= 150;
        bodyContainer.style.width = width + 'px';
    } else {
        obj.style.right = '-150px';
        bodyContainer.style.margin = '0 0 0 0';
        width += 150;
        bodyContainer.style.width = width + 'px';
    }
})

