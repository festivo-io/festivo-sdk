require_relative 'lib/festivo'

client = Festivo::Client.new(ENV['FESTIVO_API_KEY'])

# Get all holidays for a country and year
puts client.get_holidays('US', 2026)

# Get city-level holidays
puts client.get_city_holidays('IT', 'IT-MILAN', 2026)

# Get regional holidays
puts client.get_regional_holidays('GB', 'GB-SCT', 2026)

# Check if a date is a holiday
puts client.check_holiday('US', '2026-12-25')

