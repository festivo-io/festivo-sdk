<?php
namespace Festivo;

use GuzzleHttp\Client as Guzzle;

class FestivoClient {
    private $baseUrl;
    private $apiKey;
    private $http;

    public function __construct($apiKey = null, $baseUrl = 'https://api.getfestivo.com') {
        $this->apiKey = $apiKey;
        $this->baseUrl = $baseUrl;
        $this->http = new Guzzle(['base_uri' => $this->baseUrl]);
    }

    public function getInvoice(string $id) {
        $res = $this->http->request('GET', '/invoices/' . $id, [
            'headers' => $this->headers()
        ]);
        return json_decode($res->getBody()->getContents(), true);
    }

    private function headers() {
        $h = ['Accept' => 'application/json'];
        if ($this->apiKey) {
            $h['Authorization'] = 'Bearer ' . $this->apiKey;
        }
        return $h;
    }
}
