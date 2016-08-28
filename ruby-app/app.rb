require "sinatra"
require "sinatra/base"
require "json"
require "thin"

set :bind, "0.0.0.0"

before do
  request.path_info.sub! %r{/$}, ""
end

get "/" do
  content_type :json
  {
    :body  => "Hello World"
  }.to_json
end

get "/ruby" do
  content_type :json
  {
    :message  => "This is from a Ruby app"
  }.to_json
end

get "/status" do
  content_type :json
  {
    :status  => 200
  }.to_json
end
