function getCookie(name) {
    var arr, reg = new RegExp("(^| )" + name + "=([^;]*)(;|$)");

    if (arr = document.cookie.match(reg)) {

        return (arr[2]);
    } else {
        return null;
    }
}

function AddAnswerComments(qid, aid) {
    uid = getCookie("uid");
    message = $("#acomments_" + aid).html();
    var json = {
        "Question_id": qid,
        "Answer_id": aid,
        "Uid": uid,
        "Answer_content": message
    };
    alert(aid);
    alert(qid);

    alert(json["Answer_content"])

    $.post("/answer/AddAnswer ", json,function(data, status) {
            alert("Data:" + data + "\nStatus: " + status);
        });
}


//鼠标滑动显示评论......

function bgchanged(id) {
    $("#qc_ " + id + "_dd ").css("display ", "");

    $("#qc_ " + id).mouseout(function() {
        $("#qc_ " + id + "_dd ").css("display ", "none ");
    });
}
//感谢

function thank(id) {

}
//评论回复

function AddComments(id) {
    $('#answer_' + id).toggle('fast');
    /*
    $("
button ").click(function() {
        $.post("
", {
                name: "
Donald Duck ",
                city: "
Duckburg "
            },
            function(data, status) {
                alert("
Data: " + data + "\
nStatus: " + status);
            });
    });
*/
}
