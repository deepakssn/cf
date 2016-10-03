console.log("Running Skidta Careers - Deepak Radhakrishnan");
$(document).ready(function() {
    $("#frm-otp").hide();
    $("#btn-email").click(function() {
      //TEST POST CALL
        $.post( "http://127.0.0.1:8081/email", { email: "deepak@skidata.com" } );
        $("#frm-email").hide();
        $("#frm-otp").show();
    });
});
