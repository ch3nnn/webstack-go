document.write('<script src="https://cdn.bootcdn.net/ajax/libs/mouse0270-bootstrap-notify/3.1.3/bootstrap-notify.min.js"></script>');

function SuccessNotify(content) {
    $.notify({
        icon: "mdi mdi-alert",
        title: "",
        message: content,
        url: "",
        target: ""
    }, {
        type: "success",
        allow_dismiss: true,
        newest_on_top: false,
        placement: {
            from: "top",
            align: "right",
        },
        offset: {
            x: "20",
            y: "20"
        },
        spacing: "10",
        z_index: "1031",
        delay: "3000",
        animate: {
            enter: "animated fadeInDown",
            exit: "animated fadeOutUp"
        },
        onClosed: null,
        mouse_over: null
    });
}
