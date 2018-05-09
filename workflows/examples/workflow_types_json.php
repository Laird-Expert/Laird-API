<?php
$YOURAPIKEY = "";
$req = curl_init("https://test2-lairdassessors.swiftcase.co.uk:443/api/v2/".$YOURAPIKEY."/workflows.json");
curl_setopt($req, CURLOPT_RETURNTRANSFER, true);
curl_setopt($req, CURLOPT_HTTPHEADER, array("User-Agent: Laird Assessors API Example","Connection: close","Accept-Encoding: gzip, deflate","Accept: */*"));
$result = curl_exec($req);

echo "Status code: ".curl_getinfo($req, CURLINFO_HTTP_CODE)."\n";
echo "Response body: ".$result."\n";

curl_close($req);

?>
