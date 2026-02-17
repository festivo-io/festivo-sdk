<?php

namespace Festivo\Tests;

use PHPUnit\Framework\TestCase;
use Festivo\FestivoClient;
use GuzzleHttp\Client as GuzzleClient;
use GuzzleHttp\Handler\MockHandler;
use GuzzleHttp\HandlerStack;
use GuzzleHttp\Psr7\Response;
use GuzzleHttp\Exception\RequestException;
use GuzzleHttp\Psr7\Request;

class FestivoClientTest extends TestCase
{
    private $client;

    protected function setUp(): void
    {
        $this->client = new FestivoClient('test-key');
    }

    public function testInitializationWithDefaultBaseUrl()
    {
        $reflection = new \ReflectionClass($this->client);
        $baseUrlProperty = $reflection->getProperty('baseUrl');
        $baseUrlProperty->setAccessible(true);
        
        $this->assertEquals('https://api.getfestivo.com', $baseUrlProperty->getValue($this->client));
    }

    public function testInitializationWithCustomBaseUrl()
    {
        $client = new FestivoClient('test-key', 'https://custom.api.com');
        $reflection = new \ReflectionClass($client);
        $baseUrlProperty = $reflection->getProperty('baseUrl');
        $baseUrlProperty->setAccessible(true);
        
        $this->assertEquals('https://custom.api.com', $baseUrlProperty->getValue($client));
    }

    public function testInitializationFromEnv()
    {
        putenv('FESTIVO_API_KEY=env-key');
        $client = new FestivoClient();
        $reflection = new \ReflectionClass($client);
        $apiKeyProperty = $reflection->getProperty('apiKey');
        $apiKeyProperty->setAccessible(true);
        
        $this->assertEquals('env-key', $apiKeyProperty->getValue($client));
        putenv('FESTIVO_API_KEY'); // Clear env var
    }

    public function testGetHolidaysBasic()
    {
        $mockHandler = new MockHandler([
            new Response(200, [], json_encode(['holidays' => [], 'total' => 0]))
        ]);
        $handlerStack = HandlerStack::create($mockHandler);
        $guzzle = new GuzzleClient(['handler' => $handlerStack]);
        
        $reflection = new \ReflectionClass($this->client);
        $httpProperty = $reflection->getProperty('http');
        $httpProperty->setAccessible(true);
        $httpProperty->setValue($this->client, $guzzle);

        $result = $this->client->getHolidays('US', 2026);

        $this->assertIsArray($result);
        $this->assertArrayHasKey('holidays', $result);
        $this->assertArrayHasKey('total', $result);
        $this->assertEquals(0, $result['total']);
    }

    public function testGetHolidaysWithOptions()
    {
        $mockHandler = new MockHandler([
            new Response(200, [], json_encode(['holidays' => [], 'total' => 0]))
        ]);
        $handlerStack = HandlerStack::create($mockHandler);
        $guzzle = new GuzzleClient(['handler' => $handlerStack]);
        
        $reflection = new \ReflectionClass($this->client);
        $httpProperty = $reflection->getProperty('http');
        $httpProperty->setAccessible(true);
        $httpProperty->setValue($this->client, $guzzle);

        $result = $this->client->getHolidays('GB', 2026, [
            'regions' => 'GB-SCT',
            'type' => 'public',
            'language' => 'en'
        ]);

        $this->assertIsArray($result);
        $this->assertArrayHasKey('holidays', $result);
    }

    public function testGetCityHolidays()
    {
        $mockHandler = new MockHandler([
            new Response(200, [], json_encode(['holidays' => [], 'total' => 0]))
        ]);
        $handlerStack = HandlerStack::create($mockHandler);
        $guzzle = new GuzzleClient(['handler' => $handlerStack]);
        
        $reflection = new \ReflectionClass($this->client);
        $httpProperty = $reflection->getProperty('http');
        $httpProperty->setAccessible(true);
        $httpProperty->setValue($this->client, $guzzle);

        $result = $this->client->getCityHolidays('IT', 'IT-MILAN', 2026);

        $this->assertIsArray($result);
        $this->assertArrayHasKey('holidays', $result);
    }

    public function testGetCityHolidaysWithOptions()
    {
        $mockHandler = new MockHandler([
            new Response(200, [], json_encode(['holidays' => [], 'total' => 0]))
        ]);
        $handlerStack = HandlerStack::create($mockHandler);
        $guzzle = new GuzzleClient(['handler' => $handlerStack]);
        
        $reflection = new \ReflectionClass($this->client);
        $httpProperty = $reflection->getProperty('http');
        $httpProperty->setAccessible(true);
        $httpProperty->setValue($this->client, $guzzle);

        $result = $this->client->getCityHolidays('IT', 'IT-ROME', 2026, [
            'type' => 'public',
            'language' => 'it'
        ]);

        $this->assertIsArray($result);
    }

    public function testGetRegionalHolidays()
    {
        $mockHandler = new MockHandler([
            new Response(200, [], json_encode(['holidays' => [], 'total' => 0]))
        ]);
        $handlerStack = HandlerStack::create($mockHandler);
        $guzzle = new GuzzleClient(['handler' => $handlerStack]);
        
        $reflection = new \ReflectionClass($this->client);
        $httpProperty = $reflection->getProperty('http');
        $httpProperty->setAccessible(true);
        $httpProperty->setValue($this->client, $guzzle);

        $result = $this->client->getRegionalHolidays('GB', 'GB-SCT', 2026);

        $this->assertIsArray($result);
        $this->assertArrayHasKey('holidays', $result);
    }

    public function testCheckHoliday()
    {
        $mockHandler = new MockHandler([
            new Response(200, [], json_encode(['is_holiday' => true]))
        ]);
        $handlerStack = HandlerStack::create($mockHandler);
        $guzzle = new GuzzleClient(['handler' => $handlerStack]);
        
        $reflection = new \ReflectionClass($this->client);
        $httpProperty = $reflection->getProperty('http');
        $httpProperty->setAccessible(true);
        $httpProperty->setValue($this->client, $guzzle);

        $result = $this->client->checkHoliday('US', '2026-12-25');

        $this->assertIsArray($result);
        $this->assertArrayHasKey('is_holiday', $result);
        $this->assertTrue($result['is_holiday']);
    }

    public function testCheckHolidayWithRegions()
    {
        $mockHandler = new MockHandler([
            new Response(200, [], json_encode(['is_holiday' => true]))
        ]);
        $handlerStack = HandlerStack::create($mockHandler);
        $guzzle = new GuzzleClient(['handler' => $handlerStack]);
        
        $reflection = new \ReflectionClass($this->client);
        $httpProperty = $reflection->getProperty('http');
        $httpProperty->setAccessible(true);
        $httpProperty->setValue($this->client, $guzzle);

        $result = $this->client->checkHoliday('IT', '2026-12-07', 'IT-MILAN');

        $this->assertIsArray($result);
        $this->assertTrue($result['is_holiday']);
    }

    public function testResponseStructure()
    {
        $mockData = [
            'holidays' => [
                [
                    'date' => '2026-01-01',
                    'name' => 'New Year\'s Day',
                    'type' => 'public',
                    'observed' => '2026-01-01',
                    'public' => true,
                    'country' => 'US',
                    'subdivisions' => []
                ]
            ],
            'total' => 1
        ];

        $mockHandler = new MockHandler([
            new Response(200, [], json_encode($mockData))
        ]);
        $handlerStack = HandlerStack::create($mockHandler);
        $guzzle = new GuzzleClient(['handler' => $handlerStack]);
        
        $reflection = new \ReflectionClass($this->client);
        $httpProperty = $reflection->getProperty('http');
        $httpProperty->setAccessible(true);
        $httpProperty->setValue($this->client, $guzzle);

        $result = $this->client->getHolidays('US', 2026);

        $this->assertIsArray($result['holidays']);
        $this->assertCount(1, $result['holidays']);
        $this->assertEquals('2026-01-01', $result['holidays'][0]['date']);
        $this->assertEquals('New Year\'s Day', $result['holidays'][0]['name']);
    }
}
