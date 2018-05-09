var http = new XMLHttpRequest();
http.open("GET", "/api/v2/YOURAPIKEY/workflows.json", true);

http.setRequestHeader("User-Agent","Laird Assessors API Example");
http.setRequestHeader("Connection","close");
http.setRequestHeader("Accept-Encoding","gzip, deflate");
http.setRequestHeader("Accept","*/*");

http.onreadystatechange = function() {
    if(http.readyState == 4 && http.status == 200) {
        console.info(http.status);
        console.info(http.responseText);
    }
}
http.send(params);

