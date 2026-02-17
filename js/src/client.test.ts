import { describe, it, expect, beforeEach, vi } from 'vitest';
import { FestivoClient } from './client';

describe('FestivoClient', () => {
  let client: FestivoClient;

  beforeEach(() => {
    client = new FestivoClient({ apiKey: 'test-key' });
  });

  describe('initialization', () => {
    it('should initialize with default base URL', () => {
      expect(client.baseUrl).toBe('https://api.getfestivo.com');
      expect(client.apiKey).toBe('test-key');
    });

    it('should accept custom base URL', () => {
      const customClient = new FestivoClient({
        apiKey: 'test-key',
        baseUrl: 'https://custom.api.com',
      });
      expect(customClient.baseUrl).toBe('https://custom.api.com');
    });

    it('should work without API key', () => {
      const noKeyClient = new FestivoClient();
      expect(noKeyClient.apiKey).toBeUndefined();
    });
  });

  describe('getHolidays', () => {
    it('should call correct endpoint with parameters', async () => {
      global.fetch = vi.fn().mockResolvedValue({
        ok: true,
        json: async () => ({ holidays: [], total: 0 }),
      });

      await client.getHolidays('US', 2026);

      expect(fetch).toHaveBeenCalledWith(
        expect.stringContaining('/v3/public-holidays/list'),
        expect.objectContaining({
          headers: expect.objectContaining({
            Accept: 'application/json',
            Authorization: 'Bearer test-key',
          }),
        })
      );

      const callUrl = (fetch as any).mock.calls[0][0];
      expect(callUrl).toContain('country=US');
      expect(callUrl).toContain('year=2026');
    });

    it('should include optional parameters', async () => {
      global.fetch = vi.fn().mockResolvedValue({
        ok: true,
        json: async () => ({ holidays: [], total: 0 }),
      });

      await client.getHolidays('GB', 2026, {
        regions: 'GB-SCT',
        type: 'public',
        language: 'en',
      });

      const callUrl = (fetch as any).mock.calls[0][0];
      expect(callUrl).toContain('regions=GB-SCT');
      expect(callUrl).toContain('type=public');
      expect(callUrl).toContain('language=en');
    });

    it('should throw on API error', async () => {
      global.fetch = vi.fn().mockResolvedValue({
        ok: false,
        status: 401,
        statusText: 'Unauthorized',
      });

      await expect(client.getHolidays('US', 2026)).rejects.toThrow(
        'API error: 401 Unauthorized'
      );
    });
  });

  describe('getCityHolidays', () => {
    it('should call getHolidays with city code in regions', async () => {
      global.fetch = vi.fn().mockResolvedValue({
        ok: true,
        json: async () => ({ holidays: [], total: 0 }),
      });

      await client.getCityHolidays('IT', 'IT-MILAN', 2026);

      const callUrl = (fetch as any).mock.calls[0][0];
      expect(callUrl).toContain('country=IT');
      expect(callUrl).toContain('regions=IT-MILAN');
      expect(callUrl).toContain('year=2026');
    });

    it('should support additional options', async () => {
      global.fetch = vi.fn().mockResolvedValue({
        ok: true,
        json: async () => ({ holidays: [], total: 0 }),
      });

      await client.getCityHolidays('IT', 'IT-ROME', 2026, {
        type: 'public',
        language: 'it',
      });

      const callUrl = (fetch as any).mock.calls[0][0];
      expect(callUrl).toContain('type=public');
      expect(callUrl).toContain('language=it');
    });
  });

  describe('getRegionalHolidays', () => {
    it('should call getHolidays with region code', async () => {
      global.fetch = vi.fn().mockResolvedValue({
        ok: true,
        json: async () => ({ holidays: [], total: 0 }),
      });

      await client.getRegionalHolidays('GB', 'GB-SCT', 2026);

      const callUrl = (fetch as any).mock.calls[0][0];
      expect(callUrl).toContain('country=GB');
      expect(callUrl).toContain('regions=GB-SCT');
      expect(callUrl).toContain('year=2026');
    });
  });

  describe('checkHoliday', () => {
    it('should call check endpoint with correct parameters', async () => {
      global.fetch = vi.fn().mockResolvedValue({
        ok: true,
        json: async () => ({ is_holiday: true }),
      });

      await client.checkHoliday('US', '2026-12-25');

      const callUrl = (fetch as any).mock.calls[0][0];
      expect(callUrl).toContain('/v3/public-holidays/list/check');
      expect(callUrl).toContain('country=US');
      expect(callUrl).toContain('date=2026-12-25');
    });

    it('should support optional regions parameter', async () => {
      global.fetch = vi.fn().mockResolvedValue({
        ok: true,
        json: async () => ({ is_holiday: true }),
      });

      await client.checkHoliday('IT', '2026-12-07', { regions: 'IT-MILAN' });

      const callUrl = (fetch as any).mock.calls[0][0];
      expect(callUrl).toContain('regions=IT-MILAN');
    });
  });

  describe('type safety', () => {
    it('should return properly typed HolidaysResponse', async () => {
      const mockResponse = {
        holidays: [
          {
            date: '2026-01-01',
            name: 'New Year',
            type: 'public',
            observed: '2026-01-01',
            public: true,
            country: 'US',
            subdivisions: [],
          },
        ],
        total: 1,
      };

      global.fetch = vi.fn().mockResolvedValue({
        ok: true,
        json: async () => mockResponse,
      });

      const result = await client.getHolidays('US', 2026);
      expect(result).toEqual(mockResponse);
      expect(Array.isArray(result.holidays)).toBe(true);
      expect(typeof result.total).toBe('number');
    });

    it('should return properly typed CheckHolidayResponse', async () => {
      const mockResponse = {
        is_holiday: true,
        holiday: {
          date: '2026-12-25',
          name: 'Christmas',
          type: 'public',
          observed: '2026-12-25',
          public: true,
          country: 'US',
          subdivisions: [],
        },
      };

      global.fetch = vi.fn().mockResolvedValue({
        ok: true,
        json: async () => mockResponse,
      });

      const result = await client.checkHoliday('US', '2026-12-25');
      expect(result).toEqual(mockResponse);
      expect(typeof result.is_holiday).toBe('boolean');
    });
  });
});
