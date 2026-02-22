require 'net/http'
require 'uri'
require 'json'

module Festivo
  class Client
    attr_accessor :api_key, :base_url

    def initialize(api_key = nil, base_url = 'https://api.getfestivo.com')
      @api_key = api_key
      @base_url = base_url
    end

    def get_holidays(country, year, regions: nil, type: nil, language: nil, timezone: nil)
      params = { country: country, year: year }
      params[:regions] = regions if regions
      params[:type] = type if type
      params[:language] = language if language
      params[:timezone] = timezone if timezone
      request('/v3/public-holidays/list', params)
    end

    def get_city_holidays(country, city_code, year, type: nil, language: nil, timezone: nil)
      get_holidays(country, year, regions: city_code, type: type, language: language, timezone: timezone)
    end

    def get_regional_holidays(country, region_code, year, type: nil, language: nil, timezone: nil)
      get_holidays(country, year, regions: region_code, type: type, language: language, timezone: timezone)
    end

    def check_holiday(country, date, regions: nil)
      params = { country: country, date: date }
      params[:regions] = regions if regions
      request('/v3/public-holidays/check', params)
    end

    private

    def request(path, params = {})
      uri = URI.join(@base_url, path)
      uri.query = URI.encode_www_form(params)
      req = Net::HTTP::Get.new(uri)
      req['Accept'] = 'application/json'
      req['Authorization'] = "Bearer #{@api_key}" if @api_key
      res = Net::HTTP.start(uri.hostname, uri.port, use_ssl: uri.scheme == 'https') { |http| http.request(req) }
      raise "Festivo API error: #{res.code} #{res.message}" unless res.is_a?(Net::HTTPSuccess)
      JSON.parse(res.body)
    end
  end
end
