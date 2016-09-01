var api_url_base = '/api'

function poll_status() {
    $.ajax({
        type: "GET",
        url: api_url_base + "/tasks",
        data: {},
        dataType: "json",
        success: function (response) {
            if(response.code == 0) {
                update_ui(response.data)
            }
        }
    });
}

function reset() {
    $.ajax({
        type: "GET",
        url: api_url_base + "/status/reset",
        data: {},
        dataType: "json",
        success: function (response) {
            console.log(response)
            if(response.code == 0) {
                update_ui(response.data)
            }
        }
    });
}

function update_ui(data) {
    $('#num_clients').text(data.connected_clients)
    $('#num_recv_msgs').text(data.received_msgs)
}

$(function () {
    setInterval(poll_status, 500)
    $('#reset-button').click(function () { 
        reset();
    });   
});

