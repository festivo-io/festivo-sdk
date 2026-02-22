import com.festivo.FestivoClient;
import java.util.HashMap;

public class Example {
    public static void main(String[] args) throws Exception {
        FestivoClient client = new FestivoClient(System.getenv("FESTIVO_API_KEY"));
        // Get all holidays
        FestivoClient.HolidaysResponse holidays = client.getHolidays("US", 2026, null);
        System.out.println(holidays.holidays);
        // Get city holidays
        FestivoClient.HolidaysResponse cityHolidays = client.getCityHolidays("IT", "IT-MILAN", 2026, null);
        System.out.println(cityHolidays.holidays);
        // Get regional holidays
        FestivoClient.HolidaysResponse regionalHolidays = client.getRegionalHolidays("GB", "GB-SCT", 2026, null);
        System.out.println(regionalHolidays.holidays);
        // Check if a date is a holiday
        System.out.println(client.checkHoliday("US", "2026-12-25", null));
    }
}

