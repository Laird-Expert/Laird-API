require "net/http"
require "uri"
require "openssl"

uri = URI.parse("https://test2-lairdassessors.swiftcase.co.uk/api/v2/YOURAPIKEY/workflows.json")
http = Net::HTTP.new(uri.host, uri.port)
http.use_ssl = true

request = Net::HTTP::Get.new(uri.request_uri)
request["User-Agent"] = "Laird Assessors API Example"
request["Connection"] = "close"
request["Accept-Encoding"] = "gzip, deflate"
request["Accept"] = "*/*"
response = http.request(request)

puts "Status code: "+response.code
puts "Response body: "+response.body
