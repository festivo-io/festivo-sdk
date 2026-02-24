package io.festivo;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

public class FestivoClientTest {
    @Test
    public void testClientInitialization() {
        FestivoClient client = new FestivoClient("test-key");
        assertNotNull(client);
    }
}

