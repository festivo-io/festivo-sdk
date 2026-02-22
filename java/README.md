# Festivo Java SDK

Official Java SDK for the Festivo Public Holidays API.

## Installation

Add to your Maven/Gradle project (coming soon):

```xml
<dependency>
  <groupId>com.festivo</groupId>
  <artifactId>festivo-sdk</artifactId>
  <version>0.1.0</version>
</dependency>
```

## Usage

```java
import com.festivo.FestivoClient;
FestivoClient client = new FestivoClient("YOUR_API_KEY");
FestivoClient.HolidaysResponse holidays = client.getHolidays("US", 2026, null);
System.out.println(holidays.holidays);
```

## API Methods

- `getHolidays(country, year, options)`
- `getCityHolidays(country, cityCode, year, options)`
- `getRegionalHolidays(country, regionCode, year, options)`
- `checkHoliday(country, date, regions)`

## Development

```bash
cd festivo-sdk/java
javac -cp .:org.json.jar src/FestivoClient.java example.java
java -cp .:org.json.jar Example
```

## Links

- [Official Website](https://getfestivo.com)
- [API Documentation](https://docs.getfestivo.com)
- [GitHub Repository](https://github.com/festivo-io/festivo-sdk)

## Support

- Email: support@getfestivo.com
- [Report Issues](https://github.com/festivo-io/festivo-sdk/issues)

## License

MIT License - see [LICENSE](../LICENSE) for details.
