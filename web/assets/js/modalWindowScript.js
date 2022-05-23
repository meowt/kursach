function openModal(windowName) {
    if (windowName == "#logModalWindow") {
        var vr = $("#regModalWindow");
        vr.removeClass('is-show');
    } else {
        var vl = $("#logModalWindow");
        vl.removeClass('is-show');
    }
    var modal = $(windowName);
    modal.addClass('is-show');
    // document.body.style.overflow = "hidden";
}

function closeModal(windowName) {
    var modal = $(windowName)
    modal.removeClass('is-show');
    // document.body.style.overflow = "unset";
}
