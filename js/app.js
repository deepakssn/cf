console.log("Running Skidta Careers - Deepak Radhakrishnan");
$(document).ready(function() {
    $("#frm-otp").hide();
    $("#btn-email").click(function() {
        $.post( "http://127.0.0.1:8081/email", { email: "2pm" } );
        // $.get("http://localhost:8081/domains", function(data) {
        //     $(".result").html(data);
        //     alert("Load was performed.");
        // });
        $("#frm-email").hide();
        $("#frm-otp").show();
    });
});
