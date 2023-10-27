$(document).ready(function(){
    $(window).on('load', function() {
        setTimeout(function(){
            $(".loader").hide("slide", {direction: "up"}, 400);
        }, 400);
    });
});


$(".goto-login").click(function(){
    setTimeout(function(){
        $(".nmlogin").hide("slide", {direction: "up"}, 400);
        $(".signup").hide("slide", {direction: "up"}, 400);
        $(".login").show("slide", {direction: "down"}, 400);
    })
})

$(".goto-signup").click(function(){
    setTimeout(function(){
        $(".nmlogin").hide("slide", {direction: "up"}, 400);
        $(".login").hide("slide", {direction: "up"}, 400);
        $(".signup").show("slide", {direction: "down"}, 400);
    })
})

$(".goto-nm").click(function(){
    setTimeout(function(){
        $(".login").hide("slide", {direction: "down"}, 400);
        $(".signup").hide("slide", {direction: "down"}, 400);
        $(".nmlogin").show("slide", {direction: "up"}, 400);
    })
})

function checkEmailValidity(){
    var email = document.getElementById("signup-email").value;
    if(email.indexOf("@") == -1 || email.indexOf(".") == -1){
        $(".error-systems").show("slide", {direction: "down"}, 400);
        document.querySelector(".error-systems").classList.add("error-active");
    } else{
        $(".error-systems").hide("slide", {direction: "up"}, 400);
        document.querySelector(".error-systems").classList.remove("error-active");

    }
    setTimeout(function(){
        changeSignupWidth()
    }, 400)
}

function checkPasswordValidity(){
    var password = document.getElementById("signup-password").value;
    var password2 = document.getElementById("check-password").value;
    if(password != password2){
        $(".error-systems2").show("slide", {direction: "down"}, 400);        
        document.querySelector(".error-systems2").classList.add("error-active");
        
    }else{
        document.querySelector(".error-systems2").classList.remove("error-active");
        $(".error-systems2").hide("slide", {direction: "up"}, 400);
    }
    setTimeout(function(){
        changeSignupWidth()
    },400)
}

function changeSignupWidth(){
    if($(".error-systems2").hasClass("error-active") === true || $(".error-systems").hasClass("error-active") === true){
        $(".ls-main").animate({height: "35em"}, 200);
        if ($(".error-systems2").hasClass("error-active") === true && $(".error-systems").hasClass("error-active") === true){
            $(".ls-main").animate({height: "37em"}, 200).finish();
            return
        }
        return
    } else{
        $(".ls-main").animate({height: "34em"}, 200);
    }
    
}


function checkSignupValidity(){
    document.getElementById
}


var today = new Date();
var dd = today.getDate();
var mm = today.getMonth() + 1;
var yyyy = today.getFullYear();

if (dd < 10) {
   dd = '0' + dd;
}

if (mm < 10) {
   mm = '0' + mm;
} 
    
today = yyyy + '-' + mm + '-' + dd;
if (document.getElementById("datepicker")){
    document.getElementById("datepicker").setAttribute("max", today);
}
if (document.getElementById("booking-date")){
    document.getElementById("booking-date").setAttribute("min", today);
}
$("#signup-form").submit(function(){
    var email = document.getElementById("signup-email").value;
    var fname = document.getElementById("signup-fname").value;
    var lname = document.getElementById("signup-lname").value;
    var dob = document.getElementById("datepicker").value;
    var mindate = new Date("1920-01-01")
    var mindate_month = mindate.getMonth() + 1
    var mindate_date = mindate.getDate()
    var mindate_year = mindate.getFullYear()
    if (mindate_month < 10){
        mindate_month = "0" + mindate_month
    }
    if (mindate_date < 10){
        mindate_date = "0" + mindate_date
    }
    mindate = mindate_year + "-" + mindate_month + "-" + mindate_date
    var phone = document.getElementById("signup-phone").value;
    var password = document.getElementById("signup-password").value;
    var password2 = document.getElementById("check-password").value;
    var file = document.getElementById("aadharupload").files.length;
    event.preventDefault();
    if (email === ""){
        $("#signup-email").animate({borderColor: "red", borderWidth: '1px'}, 400).finish();
    } else{
        $("#signup-").animate({borderColor: "#6C6C6C", borderWidth: '0.5px'}, 400).finish()
    }
    if (fname === ""){
       $("#signup-fname").animate({borderColor: "red", borderWidth: '1px'}, 400).finish();
    } else{
        $("#signup-").animate({borderColor: "#6C6C6C", borderWidth: '0.5px'}, 400).finish()
    }
    if (lname === ""){
       $("#signup-lname").animate({borderColor: "red", borderWidth: '1px'}, 400).finish();
    } else{
        $("#signup-").animate({borderColor: "#6C6C6C", borderWidth: '0.5px'}, 400).finish()
    }
    if (dob === ""){
       $("#datepicker").animate({borderColor: "red", borderWidth: '1px'}, 400).finish();
    } else{
        $("#signup-").animate({borderColor: "#6C6C6C", borderWidth: '0.5px'}, 400).finish()
        if (document.getElementById("datepicker").value < today && document.getElementById("datepicker").value > mindate){
            $("#datepicker").animate({borderColor: "#6C6C6C", borderWidth: '0.5px'}, 400).finish();
        } else{
            $("#datepicker").animate({borderColor: "red", borderWidth: '1px'}, 400).finish();
            alert("Please enter a valid date of birth")
        }

    }
    if (phone === ""){
       $("#signup-phone").animate({borderColor: "red", borderWidth: '1px'}, 400).finish();
    } else{
        $("#signup-").animate({borderColor: "#6C6C6C", borderWidth: '0.5px'}, 400).finish()
    }
    if (password === ""){
       $("#signup-password").animate({borderColor: "red", borderWidth: '1px'}, 400).finish();
    } else{
        $("#signup-").animate({borderColor: "#6C6C6C", borderWidth: '0.5px'}, 400).finish()
    }
    if (password2 === ""){
       $("#check-password").animate({borderColor: "red", borderWidth: '1px'}, 400).finish();
    } else{
        $("#signup-").animate({borderColor: "#6C6C6C", borderWidth: '0.5px'}, 400).finish()
    }
    if (file === 0){
       $("#aubtn").animate({borderColor: "red", borderWidth: '1px'}, 400).finish();
    } 
    
    if (email != "" && fname != "" && lname != "" && dob != "" && phone != "" && password != "" && password2 != "" && file != 0 && password === password2 && email.indexOf("@") != -1 && email.indexOf(".") != -1){
        var formData = new FormData(); 
        var fileInput = document.getElementById("aadharupload"); 
        var file = fileInput.files[0];
        formData.append("fname", fname); 
        formData.append("lname", lname);
        formData.append("dob", dob);
        formData.append("phone", phone);
        formData.append("email", email);
        formData.append("password", password);
        formData.append("file", file); 
        console.log(formData)
        $.ajax({
            url: "/ajax/signup",
            type: "POST",
            data: formData, 
            processData: false,
            cache: false,
            contentType: false,
            success: function(response) {
                console.log("response: ", response)
                if (response['status'] === "success"){
                    window.location.reload()
                }
            },
            error: function(error) {
                console.log(error);
                alert(error['responseJSON']['status'])
            }
        });
    } else {
        alert("Invalid form data");
    }
   
})

$("#login-form").submit(function(){
    var email = document.getElementById("login-email").value;
    var password = document.getElementById("login-password").value;
    event.preventDefault();
    $.ajax({
        url: "/ajax/login",
        type: "POST",
        data: {
            email: email,
            password: password
        },
        success: function(response) {
            console.log(response)
            if (response['status'] === "authorised"){
                window.location.reload()
            }
        },
        error: function(error) {
            console.log(error);
            if(error['status'] === 409){
                alert("Invalid credentials")
            }
        }
    });
})
var allnavlinks = document.getElementsByClassName("navlink-active")
$(".goto-home").click(function() {
	$(allnavlinks).removeClass("navlink-active")
	$(this).addClass("navlink-active")
	$(".dashboard-gradient").fadeIn()
                $(".music-container").fadeOut()
                $(".bookings-step2-container").fadeOut()
	$("#uheh").fadeOut()
	$(".booking-gradient").fadeOut()
	$(".bookin-train").hide("slide", {
		direction: "down"
	}, 500)
	$(".checkin-s1").hide("slide", {
		direction: "left"
	}, 600)
	$(".checkin-s3").hide("slide", {
		direction: "right"
	}, 600)
	$(".checkin-s2").hide("slide", {
		direction: "up"
	}, 600)
	setTimeout(function() {
		$(".bookings-tickets").hide("slide", {
			direction: "left"
		}, 600)
		$(".booking-userarea").hide("slide", {
			direction: "right"
		}, 600)
		$(".dashboard-main").fadeIn(100)
		$(".checkin-container").fadeOut(100)

	}, 550)
	setTimeout(function() {
		$(".dashboard-section1").show("slide", {
			direction: "left"
		}, 600)
		setTimeout(function() {
			$(".dashboard-section3").show("slide", {
				direction: "right"
			}, 600)
		}, 300)
		$(".dashboard-section2").show("slide", {
			direction: "up"
		}, 300)
	}, 1100)
})

$(".goto-book").click(function() {
	$(allnavlinks).removeClass("navlink-active")
	$(this).addClass("navlink-active")
                $(".music-container").fadeOut()
                $(".dashboard-gradient").fadeOut()
	$(".bookings-step2-container").fadeOut()
	$("#uheh").fadeOut(200)
	$(".booking-gradient").fadeIn()
	$(".dashboard-section1").hide("slide", {
		direction: "left"
	}, 600)
	$(".dashboard-section3").hide("slide", {
		direction: "right"
	}, 600)
	$(".dashboard-section2").hide("slide", {
		direction: "up"
	}, 600)
	$(".checkin-s1").hide("slide", {
		direction: "left"
	}, 600)
	$(".checkin-s3").hide("slide", {
		direction: "right"
	}, 600)
	$(".checkin-s2").hide("slide", {
		direction: "up"
	}, 600)
	setTimeout(function() {
		$(".dashboard-main").fadeOut(100)
		$(".checkin-container").fadeOut(100)
		setTimeout(function() {
			$(".bookin-train").hide()
			$(".bookings-tickets").hide()
			$(".booking-userarea").hide()
			$(".bookings-container").show()

			setTimeout(function() {
				$(".bookings-tickets").show("slide", {
					direction: "left"
				}, 600)
                setTimeout(function(){
                    $(".booking-userarea").show("slide", {
                        direction: "right"
                    }, 600)
                }, 500)
				setTimeout(function() {
					$(".bookin-train").show("slide", {
						direction: "down"
					}, 1500)
				}, 500)
			}, 100)
		}, 200)
	}, 600)
})

$(".goto-checkin").click(function() {
	$(allnavlinks).removeClass("navlink-active")
	$(this).addClass("navlink-active")
	$(".dashboard-gradient").fadeIn()
	$(".bookings-step2-container").fadeOut()
                $(".music-container").fadeOut()
                $("#uheh").fadeOut(200)
	$(".booking-gradient").fadeOut()
	$(".dashboard-section1").hide("slide", {
		direction: "left"
	}, 600)
	$(".dashboard-section3").hide("slide", {
		direction: "right"
	}, 600)
	$(".dashboard-section2").hide("slide", {
		direction: "up"
	}, 600)
	setTimeout(function() {
		$(".dashboard-main").fadeOut(100)
		setTimeout(function() {
			$(".bookin-train").hide("slide", {
				direction: "down"
			}, 500)
			setTimeout(function() {
				$(".bookings-tickets").hide("slide", {
					direction: "left"
				}, 600)
				$(".booking-userarea").hide("slide", {
					direction: "right"
				}, 600)
				$(".dashboard-main").fadeOut(100)
			}, 550)
			$(".checkin-container").show()

			setTimeout(function() {
				$(".checkin-s1").show("slide", {
					direction: "left"
				}, 600)
				setTimeout(function() {
					$(".checkin-s3").show("slide", {
						direction: "right"
					}, 600)
				}, 300)
				$(".checkin-s2").show("slide", {
					direction: "up"
				}, 300)
			}, 1100)

		}, 200)
	}, 600)
})

$(".goto-music").click(function(){
    $(allnavlinks).removeClass("navlink-active")
	$(this).addClass("navlink-active")
	$(".dashboard-gradient").fadeOut()
	$(".bookings-step2-container").fadeOut()
	$("#uheh").fadeOut(200)
	$(".booking-gradient").fadeOut()
    $(".music-gradient").fadeIn()
    $(".booking-gradient").fadeOut()
	$(".dashboard-section1").hide("slide", {
		direction: "left"
	}, 600)
	$(".dashboard-section3").hide("slide", {
		direction: "right"
	}, 600)
	$(".dashboard-section2").hide("slide", {
		direction: "up"
	}, 600)
	setTimeout(function() {
		$(".dashboard-main").fadeOut(100)
		setTimeout(function() {
			$(".bookin-train").hide("slide", {
				direction: "down"
			}, 500)
			setTimeout(function() {
				$(".bookings-tickets").hide("slide", {
					direction: "left"
				}, 600)
				$(".booking-userarea").hide("slide", {
					direction: "right"
				}, 600)
				$(".dashboard-main").fadeOut(100)
			}, 550)
            setTimeout(function() {
                $(".checkin-s1").hide("slide", {
                    direction: "left"
				}, 600)
				setTimeout(function() {
					$(".checkin-s3").hide("slide", {
						direction: "right"
					}, 600)
				}, 300)
				$(".checkin-s2").hide("slide", {
                    direction: "up"
				}, 300)
                $(".music-container").fadeIn()
			}, 1100)
            
        }, 200)

})
})

// $(".dashboard-main").fadeOut(100)

$("#booking-from").click(function(){
    $('.booking-stations').fadeIn()
})
$("#booking-to").click(function(){
    $('.bs22').fadeIn()
})
$(".suqwrj").click(function(){
    $('.bs22').fadeOut()

    $('.booking-stations').fadeOut()
})


var originStationCode = null;
var destinationStationCode = null;

$(".sysqu3r").click(function () {
    var id = $(this).attr('id');
    id = id.replace("station-", "");

    originStationCode = id;

    $.ajax({
        url: "/ajax/getstation",
        type: "POST",
        data: {
            id: id
        },
        success: function (response) {
            $("#booking-from").val(response['location']);
            $("#booking-from").attr("id", `booking-from selected-${response['code']}`);
            $('.booking-stations').fadeOut();
            handleTripData();
        },
        error: function (error) {
            console.log(error);
        }
    });
});

$(".sysjrewj").click(function () {
    var id = $(this).attr('id');
    id = id.replace("station-", "");

    destinationStationCode = id;

    $.ajax({
        url: "/ajax/getstation",
        type: "POST",
        data: {
            id: id
        },
        success: function (response) {
            $("#booking-to").val(response['location']);
            $("#booking-to").attr("id", `booking-to selectedto-${response['code']}`);
            $('.bs22').fadeOut();
            handleTripData();
        },
        error: function (error) {
            console.log(error);
        }
    });
});


$(".booking-userform").submit(function(){
    event.preventDefault()
    $("#booking-submit").click()
})
$("#booking-submit").click(function(){
    storeForm1()
})
var maxSeatsToSelect = $("#booking-passengers").val(); 
$("#booking-passengers").change(function () {
    if ($("#booking-passengers").val() < maxSeatsToSelect) {
        selectedSeats = [];
        $(".seat-selected").animate({backgroundColor: "#CBCBCB"}, 400).finish();
    }
    maxSeatsToSelect = $("#booking-passengers").val(); 
    console.log("maxSeatsToSelect: ", maxSeatsToSelect)
    handleTripData();
});
var selectedSeats = [];
$(".lvisbhaiii").click(function(){
    
    if ($(this).hasClass("seat-occupied")) {
        console.log("This seat is already occupied");
        return;
    }
    var seatid = $(this).attr('id');
    $(".lvisbhaiii").css("cursor", "pointer")
    seatid = seatid.replace("seat-", "");
    if (selectedSeats.includes(seatid)) {
        selectedSeats.splice(selectedSeats.indexOf(seatid), 1);
        $(this).animate({backgroundColor: "#CBCBCB"}, 200).finish();
        return;
    }
    if (selectedSeats.length >= maxSeatsToSelect) {
        $(".lvisbhaiii").css("cursor", "not-allowed")
        return;
    }
    selectedSeats.push(seatid);
    $(this).animate({backgroundColor: "rgb(50, 154, 251)"}, 200).finish();
    $(this).addClass("seat-selected");

    console.log("Selected seats:", selectedSeats);
})


function handleTripData() {
    if (originStationCode  && destinationStationCode) {
        $.ajax({
            url: "/ajax/tripdata",
            type: "POST",
            data: {
                origin: originStationCode,
                destination: destinationStationCode
            },
            success: function (response) {
                console.log(response);
                const origin = $('.booking-from').val();
                const destination = $('.booking-to').val();
                console.log(origin, destination)
                var occupied = response['occupied'];
                for (var i = 0; i < occupied.length; i++) {
                    var seat = occupied[i];
                    var seatID = `seat-${seat}`;
                    $(`#${seatID}`).addClass("seat-occupied");
                }
                var stationsContainer = document.querySelector(".all-stations")
                stationsContainer.innerHTML = "";
                $(".bst").text(origin)
                $(".bstt").text(destination)
                for (var i = 0; response['stations'].length > i; i++) {
                    $.ajax({
                        url: "/ajax/getstation",
                        type: "POST",
                        data: {
                            id: response['stations'][i]
                        },
                        success: function (response) {
                            var value = response['location'];
                            var bookingStationDiv = document.createElement('div');
                            bookingStationDiv.className = 'booking-station b-sss';
                            var image1 = document.createElement('img');
                            image1.src = '/static/images/icons/bs.svg';
                            image1.className = 'b-stop b-stopimg';
                            image1.alt = '';
                            bookingStationDiv.appendChild(image1);
                            var p1 = document.createElement('p');
                            p1.className = 'booking-station-text';
                            p1.textContent = value;
                            bookingStationDiv.appendChild(p1);
                            var bookingDivvvDiv = document.createElement('div');
                            bookingDivvvDiv.className = 'booking-divvv';
                            var image2 = document.createElement('img');
                            image2.src = '/static/images/icons/bdiv.svg';
                            image2.className = 'bdiv-img';
                            image2.alt = '';
                            bookingDivvvDiv.appendChild(image2);

                            stationsContainer.appendChild(bookingStationDiv);
                            stationsContainer.appendChild(bookingDivvvDiv);
                        },
                        error: function (error) {
                            console.log(error);
                        }
                    });

                }
                const apiKey = 'ieZCkecJUpvKR0UK1c22D6eBP06kv4eh'; 
                const apiUrl = `https://www.mapquestapi.com/directions/v2/route?key=${apiKey}&from=${origin}&to=${destination}`;
                $("#staiotn").text(response['len'])
                $.get(apiUrl, function(data) {
                    if (data.route) {
                        console.log("data: ", data.route);
                        const distance = Math.floor(data.route.distance);
                        const time = data.route.formattedTime;
                        $("#taime").text(time + " Hrs");
                        var passengers = $("#booking-passengers").val();
                        var fair = Math.floor(distance * 1.4 * passengers);
                        $("#fairr").text("Rs. " + fair);
                        console.log(fair);
                        $("#distunce").text(distance + " KM");
                    } else {
                        console.log('system hang');
                    }
                }).fail(function() {
                    console.error('system crash');
                });
            },
            error: function (error) {
                console.log(error);
            }
        });
    }
}
var baseprice = 0
var storeorigin = null
var storedestination = null
var storeseats = null
var storedate = null
var storetime = null
var storepassengers = null
var storefair = null
var storestart = null
var storetrain = null    
function storeForm1(){
    storeorigin = originStationCode
    storedestination = destinationStationCode
    storeseats = selectedSeats;
    storedate = document.getElementById("booking-date").value;
    storedate = storedate.replace(" Hrs.", "")
    storetime = document.getElementById("taime").textContent;
    storetime = storetime.replace(" Hrs", "")
    storestart = document.getElementById("booking-time").value;
    storepassengers = document.getElementById("booking-passengers").value;
    storefair = document.getElementById("fairr").textContent;
    storefair = storefair.replace("Rs. ", "")
    storefair = parseInt(storefair)
    console.log(storefair)
    
    $.ajax({
        url: "/ajax/gettrain",
        type: "POST",
        data: {
            origin: storeorigin,
            destination: storedestination
        },
        success: function (response) {
            storetrain = response['id']
        },
        error: function (error) {
            console.log(error);
        }
    });
    event.preventDefault();
    if (originStationCode === null || destinationStationCode === null || selectedSeats.length === 0 || document.getElementById("booking-date").value === "" || document.getElementById("taime").textContent === "" || document.getElementById("booking-passengers").value === "" || document.getElementById("fairr").textContent === "" || selectedSeats.length < document.getElementById("booking-passengers").value){
        alert("Please fill all the fields")
        return
    }
    goToBookings2(storeorigin, storedestination, storetrain, storeseats, storedate, storetime, storepassengers, storefair)

}
var totalcost = 0
function goToBookings2(storeorigin, storedestination, storetrain, storeseats, storedate, storetime, storepassengers, storefair){
    $(".bookings-container").hide("slide", {direction: "up"}, 200)
    setTimeout(function(){

        $(".bookings-step2-container").show("slide", {direction: "up"}, 200)
    },200)
    baseprice += storefair
    baseprice += 2000
    updatePrice()
    
}

function updatePrice(){
    $("#addcost").text("₹ " + (Math.floor(baseprice * 8/100)))
    $("#baseprice").text("₹ " +  baseprice)
    totalcost = baseprice+Math.floor(baseprice * 18/100)
    $("#totalcost-book").text("₹ " + totalcost)
}

$(document).on("click", ".removehotel", function() {
 
    $(".bookings2-input").attr("disabled", "disabled");
    $(".bookings2-input").animate({
        backgroundColor: "#3A3A3A"
    });
    $(".bookings2-input").css("cursor", "not-allowed");
    $(".bookings2-input").val("");
    $(this).animate({
        color: "#2CE328",
        backgroundColor: "#1E3222"
    });
    $(`#ong-status2`).text("add");
    $(this).addClass("addhotel");
    $(this).removeClass("removehotel");
    baseprice -= 2000
    updatePrice()

});

$(document).on("click", ".addhotel", function() {
    $(".bookings2-input").removeAttr("disabled");
    $(".bookings2-input").animate({
        backgroundColor: "#1c1c1c"
    });
    $(".bookings2-input").css("cursor", "pointer");
    $(this).animate({
        color: "#E33E28",
        backgroundColor: "#32211E"
    });
    $(`#ong-status2`).text("remove");
    $(this).addClass("removehotel");
    $(this).removeClass("addhotel");
    baseprice += 2000
    updatePrice()
});


function bookingSearch(){
    let searchquery = document.getElementById('sije').value 
    searchquery=searchquery.toUpperCase(); 
    let stations = document.getElementsByClassName('sysqu3r'); 
    for (i = 0; i < stations.length; i++) {  
        if (!stations[i].innerHTML.toUpperCase().includes(searchquery.toUpperCase())) { 
            stations[i].style.display="none"; 
        } 
        else { 
            stations[i].style.display="table-row";                  
        } 
    } 
}


function bookingSearch2(){
    let searchquery = document.getElementById('sijee').value 
    searchquery=searchquery.toUpperCase(); 
    let stations = document.getElementsByClassName('sysjrewj'); 

    for (i = 0; i < stations.length; i++) {  
        if (!stations[i].innerHTML.toUpperCase().includes(searchquery.toUpperCase())) { 
            stations[i].style.display="none"; 
        } 
        else { 
            stations[i].style.display="table-row";                  
        } 
    } 
}
$(".train-2").click(function(){
    console.log("434")
    $(".train1-color").attr("fill", "#969697")
    $(".train1-seccolor").attr("fill", "#a7a7a9")
    $(".train2-color").attr("fill", "black")
    $("#sysgh").attr("fill", "#a7a7a9")
    $(".train2-seccolor").attr("fill", "#2c2c2c")
    $(".train3-color").attr("fill", "#969697")
    $(".train3-seccolor").attr("fill", "#a7a7a9")
    $(".tr1").hide("slide", {direction: "up"}, 600)
    $(".tr3").hide("slide", {direction: "up"}, 600)
    $(".tr2").show("slide", {direction: "down"}, 600)
    
})

$

$(".train-3").click(function(){
    console.log("434")
    $(".train1-color").attr("fill", "#969697")
    $(".train1-seccolor").attr("fill", "#a7a7a9")
    $(".train3-color").attr("fill", "black")
    $(".train3-seccolor").attr("fill", "#2c2c2c")
    $(".train2-color").attr("fill", "#969697")
    $("#sysgh").attr("fill", "#a7a7a9")
    $(".train2-seccolor").attr("fill", "#a7a7a9")
    $(".tr1").hide("slide", {direction: "up"}, 600)
    $(".tr2").hide("slide", {direction: "up"}, 600)
    $(".tr3").show("slide", {direction: "down"}, 600)
})
$(".train-1").click(function(){
    console.log("434")
    $(".train3-color").attr("fill", "#969697")
    $(".train3-seccolor").attr("fill", "#a7a7a9")
    $(".train1-color").attr("fill", "black")
    $(".train1-seccolor").attr("fill", "#2c2c2c")
    $(".train2-color").attr("fill", "#969697")
    $("#sysgh").attr("fill", "#2c2c2c")
    $(".train2-seccolor").attr("fill", "#a7a7a9")
    $(".tr2").hide("slide", {direction: "up"}, 600)
    $(".tr3").hide("slide", {direction: "up"}, 600)
    $(".tr1").show("slide", {direction: "down"}, 600)
})

$(".bookings-traintype").click(function(){
    $(".bookings-traintype").removeClass("active")
    $(this).addClass("active")
})
foodlist = []
$(".foodadddiv").click(function(){
    if ($(this).hasClass("foodadded") == true){
        return
    }
    eid = $(this).attr("id")
    eid = eid.replace("foodadd-", "")
    foodlist.push(eid)
    $.ajax({
        url: '/ajax/getprice',
        type: 'POST',
        data: {
            id: eid
        },
        success: function(response){
            console.log(response)
            baseprice += response['price']
            updatePrice()
            $(`#foodadd-${eid}`).addClass("foodadded")
            $(`#foodadddd-${eid}`).text("Added")
        },
        error: function(error){
            console.log(error)
        }
    })
})
$(document).on("click", ".foodadded", function(){
    eid = $(this).attr("id")
    eid = eid.replace("foodadd-", "")
    foodlist.splice(foodlist.indexOf(eid), 1)
    $.ajax({
        url: '/ajax/getprice',
        type: 'POST',
        data: {
            id: eid
        },
        success: function(response){
            console.log(response)
            baseprice -= response['price']
            updatePrice()
            $(`#foodadd-${eid}`).removeClass("foodadded")
            $(`#foodadddd-${eid}`).text("Add")
        },
        error: function(error){
            console.log(error)
        }
    })
})

$("#bookings-form").submit(function(){
    passToServer();
})
function passToServer(){
    totalcost = parseFloat(totalcost)
    var hotel 
    if ($(".removehotel").hasClass("removehotel") == true){
        hotel = 1
    } else{
        hotel =0
    }
    $.ajax({
        url: "/ajax/bookinfo",
        type: "POST",
        data:{
            from: storeorigin,
            to: storedestination,
            train: storetrain,
            seats: JSON.stringify(storeseats),
            passengers: storepassengers,
            start: storestart,
            duration: storetime,
            date: storedate,
            hotel: hotel,
            price: totalcost,
            food: JSON.stringify(foodlist),
        },
        success: function(response){
            console.log(response)
            window.location.href= "/checkout"
        },
        error: function(error){
            console.log(error)
        }
    })
}
// 969697
// A7A7A9
// async function throwError(s, c, m){
//     if (s === "error"){
//         const sm = document.getElementById("sm"); 
//         var eig = new Date()
//         errorID = `${eig.getFullYear()}${eig.getMonth()}${eig.getDate()}${eig.getHours()}${eig.getMinutes()}${eig.getSeconds()}${eig.getMilliseconds()}${Math.floor(Math.random() * 100000)}`;
//         console.log(errorID)
//         sm.innerHTML += `
//             <div class="server-message" id="${errorID}">
//                 <p class="server-message-status">${s, " ",c}</p>
//                 <p class="server-messagee">${m}</p>
//             </div>
//         `;
        
//         $(`#${errorID}`).fadeIn()
//         setTimeout(function(){
//             $(`#${errorID}`).fadeOut()
//         }, 2000)
//     }
    
// }
