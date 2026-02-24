# Festivo Java SDK

Official Java SDK for the Festivo Public Holidays API - Access holiday data for 250+ countries.

## Installation

Add to your Maven or Gradle project:

**Maven:**
```xml
<dependency>
  <groupId>io.github.festivo-io</groupId>
  <artifactId>festivo-sdk</artifactId>
  <version>0.2.1</version>
</dependency>
```

**Gradle:**
```groovy
implementation 'io.github.festivo-io:festivo-sdk:0.2.1'
```

## Usage

```java
import io.festivo.FestivoClient;
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
mvn clean install
```

## Testing

```bash
mvn test
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
